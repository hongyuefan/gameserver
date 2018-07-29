package handler

import (
	"fmt"
	"server/msg"
)

func HandlerLogin_Send(c *Client) {

	m := make(map[string]interface{}, 1)

	u := new(msg.PlayerLogin)

	fmt.Println("UserName:")
	fmt.Scanln(&(u.MobileOrEmail))
	fmt.Println("Password:")
	fmt.Scanln(&(u.Password))

	m["PlayerLogin"] = u

	if err := c.SendCmd(m); err != nil {
		fmt.Println("Login Send Error:", err.Error())
	}
	return
}
func HandlerLogin_Recv(data interface{}) {
	fmt.Println("Login Success Token:", string(data.([]byte)))
}
