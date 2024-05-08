package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommitData struct {
	ID      int    `json:"id"`
	AUTHOR  string `json:"author"`
	COMMENT string `json:"comment"`
	SHA     string `json:"sha"`
}

var CLIENT *mongo.Client
var MONGO_DB *mongo.Database
var COLLECTION *mongo.Collection
var MONGO_LOGIN string
var MONGO_PASS string
var MONGO_URL string
var MONGO_DB_NAME string
var MONGO_TYPE_CONNECT string

func ConnectMongoDB() {
	options := options.Client().ApplyURI(fmt.Sprintf("%s://%s:%s@%s/", MONGO_TYPE_CONNECT, MONGO_LOGIN, MONGO_PASS, MONGO_URL))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных: ", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Ошибка проверки доступности доступа к БД: ", err)
	} else {
		fmt.Println("Connected to MongoDB!")
		CLIENT = client
		MONGO_DB = client.Database(MONGO_DB_NAME)
	}
}

func DisconnectMongoDB() {
	if CLIENT != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := CLIENT.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}

		log.Println("Disconnected from MongoDB")
	} else {
		log.Println("MongoDB client is not initialized")
	}
}

func getCollection(collectionName string) *mongo.Collection {
	return MONGO_DB.Collection(collectionName)
}

func GetCommitData(id int, collectionName string, commit *CommitData) {

	var data *bson.M

	if err := getCollection(collectionName).FindOne(context.TODO(), bson.M{"id": id}).Decode(&data); err != nil {
		log.Println("ошибка получения данных", err)
		*commit = CommitData{
			COMMENT: "Запрошенный коммит не найден",
		}
	} else {
		jD, err := json.Marshal(data)
		if err != nil {
			log.Println("ошибка преобразования в JSON:", err)
			return
		}

		if err := json.Unmarshal(jD, &commit); err != nil {
			log.Println("ошибка распарсивания JSON:", err)
			return
		}
		log.Println(commit)
	}

}
