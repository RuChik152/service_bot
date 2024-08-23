package bot

import (
	"fmt"
	"io"
	"net/http"
)

type UploadeRequestData struct {
	status bool
	code   int
	msg    string
}

func uploadRequest() (UploadeRequestData, error) {
	var answer UploadeRequestData

	client := &http.Client{}

	res, err := client.Get(fmt.Sprintf("%s/upload", FTP_LOADER_URL))
	if err != nil {
		return UploadeRequestData{}, fmt.Errorf("uploadRequest: %w", err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return UploadeRequestData{}, fmt.Errorf("read body response: %w", err)
	}

	answer.code = res.StatusCode
	answer.msg = string(body)
	answer.status = true

	return answer, nil
}
