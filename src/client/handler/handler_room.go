package handler

import (
	"fmt"
	"server/msg"
)

func HandlerGameRoomJoin_Send(c *Client) {

	m := make(map[string]interface{}, 1)

	u := new(msg.GameRoomJoinReq)

	fmt.Println("classid")
	fmt.Scanln(&u.ClassId)
	fmt.Println("roomid")
	fmt.Scanln(&u.RoomId)
	fmt.Println("playerid")
	fmt.Scanln(&u.PlayerId)

	m["GameRoomJoinReq"] = u

	if err := c.SendCmd(m); err != nil {
		fmt.Println("GameRoomJoinReq Send Error:", err.Error())
	}
	return
}
func HandlerGameRoomJoin_Recv(data interface{}) {
	fmt.Println("GameRoomJoin Recv :", string(data.([]byte)))
}

func HandlerGameRoomExit_Send(c *Client) {
	m := make(map[string]interface{}, 1)

	u := new(msg.GameRoomExitReq)

	fmt.Println("classid")
	fmt.Scanln(&u.ClassId)
	fmt.Println("roomid")
	fmt.Scanln(&u.RoomId)
	fmt.Println("playerid")
	fmt.Scanln(&u.PlayerId)

	m["GameRoomExitReq"] = u

	if err := c.SendCmd(m); err != nil {
		fmt.Println("GameRoomExitReq Send Error:", err.Error())
	}
	return
}
func HandlerGameRoomExit_Recv(data interface{}) {
	fmt.Println("GameRoomExit Recv :", string(data.([]byte)))
}

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
	fmt.Println("GameRoomAdd Recv :", string(data.([]byte)))
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

func HandlerGetGameRoom_Recv(data interface{}) {
	fmt.Println("GetGameRoom Recv:", string(data.([]byte)))
}
