package main

import (
	ct "client/handler"
	"fmt"
	"server/msg"
	"sync"
)

func main() {
	var (
		op string
		wg sync.WaitGroup
	)

	c := ct.NewClient("tcp", "127.0.0.1:3563")

	if err := c.OnInit(); err != nil {
		fmt.Println(err.Error())
		return
	}

	c.RegistFunc(msg.Buss_RegistAndLogin_Code, ct.HandlerLogin_Recv)
	c.RegistFunc(msg.Buss_GameClassGet_Code, ct.HandlerGetGameClass_Recv)
	c.RegistFunc(msg.Buss_Chat_Code, ct.HandlerChat_Recv)
	c.RegistFunc(msg.Buss_GameRoomAdd_Code, ct.HandlerGameRoomAdd_Recv)
	c.RegistFunc(msg.Buss_GameRoomGet_Code, ct.HandlerGetGameRoom_Recv)
	c.RegistFunc(msg.Buss_GameRoomJoin_Code, ct.HandlerGameRoomJoin_Recv)
	c.RegistFunc(msg.Buss_GameRoomExit_Code, ct.HandlerGameRoomExit_Recv)

	go func() {
		wg.Add(1)
		c.OnStart()
		wg.Done()
	}()

	for {
		if _, err := fmt.Scanln(&op); err != nil {
			fmt.Println("input error ", err.Error())
		}
		switch op {
		case "rj":
			ct.HandlerGameRoomJoin_Send(c)
		case "re":
			ct.HandlerGameRoomExit_Send(c)
		case "rg":
			ct.HandlerGetGameRoom_Send(c)
		case "ra":
			ct.HandlerGameRoomAdd_Send(c)
		case "chat":
			ct.HandlerChat_Send(c)
		case "cl":
			ct.HandlerChatLogin_Send(c)
		case "l":
			ct.HandlerLogin_Send(c)
		case "gcl":
			ct.HandlerGetGameClass_Send(c)
		case "q":
			c.OnClose()
			goto Deal
		default:
			fmt.Println("Op ", op, "Not Find")
		}

	}
Deal:
	wg.Wait()
	return
}
