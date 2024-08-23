package bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type StandartMsgBot struct {
	Event   string `json:"event"`
	Message string `json:"message"`
}

type DeviceInfo struct {
	Log string `json:"log"`
}

type DeviceMessage struct {
	BuildInfo string `json:"build"`
	SendInfo  string `json:"send"`
}

type BuildResultMessage struct {
	Event  string        `json:"event"`
	Device DeviceMessage `json:"device"`
	Info   DeviceInfo    `json:"info"`
}

type CommitData struct {
	Event   string `json:"event"`
	ID      string `json:"ID"`
	SHA     string `json:"sha"`
	AUTHOR  string `json:"author"`
	MESSAGE string `json:"message"`
}

func checkEvent(msg []byte) string {
	var data map[string]interface{}

	if err := json.Unmarshal(msg, &data); err != nil {
		log.Panic("ошибка полчения данных")
		return ""
	}

	if event, ok := data["event"].(string); !ok {
		log.Panic("ошибка получения значения event")
		return ""
	} else {
		return event
	}
}

func formaterBuildMsgBot(msg []byte) string {
	var data BuildResultMessage

	if err := json.Unmarshal(msg, &data); err != nil {
		log.Panic("ошибка преобразования")
	}

	return fmt.Sprintf("------------------------------\n1) %s\n2) %s\n%s\n------------------------------,\nАльтернативный способ загрузки OCULUS ⤵️\n/upload_oculus ", data.Device.BuildInfo, data.Device.SendInfo, data.Info.Log)

}

func formaterAllowMsgBot(msg []byte) string {
	var data StandartMsgBot
	if err := json.Unmarshal(msg, &data); err != nil {
		log.Panic("ошибка преобразования")
	}

	return fmt.Sprintf("------------------------------\n%s\n------------------------------", data.Message)
}

func formaterCommitMsg(msg []byte) string {
	var data CommitData
	if err := json.Unmarshal(msg, &data); err != nil {
		log.Panic("ошибка преобразования")
	}
	return fmt.Sprintf("------------------------------\n%s\n\n%s", data.AUTHOR, data.MESSAGE)
}

func defineQuery(msg []byte) (string, error) {
	data := msg

	switch event := checkEvent(data); event {
	case "build":
		return formaterBuildMsgBot(msg), nil
	case "allow":
		return formaterAllowMsgBot(msg), nil
	case "commit":
		return formaterCommitMsg(msg), nil
	default:
		return "", errors.New("не удалось распознать событие полученное от сервера")
	}

}
