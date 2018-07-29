package handler

import (
	"encoding/json"
	"fmt"
	"server/msg"
)

func HandlerGetGameClass_Send(c *Client) {

	m := make(map[string]interface{}, 1)

	u := new(msg.GameClassGetReq)

	m["GameClassGetReq"] = u

	if err := c.SendCmd(m); err != nil {
		fmt.Println("GetGameClass Send Error:", err.Error())
	}
	return
}
func HandlerGetGameClass_Recv(data interface{}) {

	var gameClass []*msg.GameClassGetRsp

	if err := json.Unmarshal(data.([]byte), &gameClass); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range gameClass {
		fmt.Println("GetGameClass:", v)
	}
}
