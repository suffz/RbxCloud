package RbxCloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var (
	message_url = "https://apis.roblox.com/messaging-service/v1/universes/%v/topics/%v" // universeID, Topic
)

type Message struct {
	Message string `json:"message"`
}

func (App *Application) SendMessage(topic string, data Message) string {
	req, err := http.NewRequest("POST", fmt.Sprintf(message_url, App.UniversalID, topic), bytes.NewBuffer(StructToJson(data)))
	if err == nil {
		req.Header.Add("x-api-key", App.APIKey)
		req.Header.Add("content-type", "application/json")
		if resp, err := http.DefaultClient.Do(req); err == nil && resp.StatusCode == 200 {
			return `Message sent succesfully.`
		} else {
			if err == nil {
				r, _ := io.ReadAll(resp.Body)
				return fmt.Sprintf(`{"status":"%v","error":"%v"}`, resp.StatusCode, string(r))
			}
			return fmt.Sprintf(`{"status":"%v","error":"%v"}`, "SYSTEM_ERRROR", err)
		}
	}
	return fmt.Sprintf(`{"status":"%v","error":"%v"}`, "SYSTEM_ERRROR", err)
}

func StructToJson(J any) []byte {
	if js, err := json.Marshal(J); err == nil {
		return js
	}
	return []byte{}
}
