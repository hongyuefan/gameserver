package handler

import (
	"fmt"
	"server/msg"
)

func HandlerGameRoomAdd_Send(c *Client) {
	m := make(map[string]interface{}, 1)

	u := new(msg.GameRoomAdd)

	fmt.Println("creatorId")
	fmt.Scanln(&u.CreatorId)
	fmt.Println("classid")
	fmt.Scanln(&u.GameClassId)
	fmt.Println("password")
	fmt.Scanln(&u.RoomPass)

	m["GameRoomAdd"] = u

	if err := c.SendCmd(m); err != nil {
		fmt.Println("GameRoomAdd Send Error:", err.Error())
	}
	return
}
func HandlerGameRoomAdd_Recv(data interface{}) {
	fmt.Println("GameRoomAdd Recv :", data)
}

func HandlerGetGameRoom_Send(c *Client) {

	m := make(map[string]interface{}, 1)

	u := new(msg.GameRoomGetReq)

	fmt.Println("ClassId:")

	fmt.Scanln(&(u.ClassId))

	m["GameRoomGetReq"] = u

	if err := c.SendCmd(m); err != nil {
		fmt.Println("GetGameRoom Send Error:", err.Error())
	}
	return
}

func handlerGetGameRoom_Recv(data interface{}) {
	fmt.Println("GetGameRoom Recv:", data.(*msg.GameRoomGetRsp))
}
