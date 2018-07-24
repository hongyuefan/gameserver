package main

import (
	ct "client"
	"fmt"
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

	c.RegistFunc(ct.Buss_RegistAndLogin_Code, ct.HandlerLogin_Recv)
	c.RegistFunc(ct.Buss_GetGameClass_Code, ct.HandlerGetGameClass_Recv)
	c.RegistFunc(ct.Buss_Chat_Code, ct.HandlerChat_Recv)

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
