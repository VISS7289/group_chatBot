package Chat

import (
	"bytes"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

func GetChat(ChatContent string) ([]byte,error) {
	url := "http://localhost:8087/chat"
	post := "{\"chat_content\":\"" + ChatContent +
		"\"}"
	var jsonStr = []byte(post)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		zap.L().Error("Python Server Error", zap.Error(err))
		return nil,err
	}
	return ioutil.ReadAll(resp.Body)
}
