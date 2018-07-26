package handler

import (
	"fmt"
)

func HandlerLogin_Send(c *Client) {

	m := make(map[string]interface{}, 1)

	u := new(UserLogin)

	fmt.Println("UserName:")
	fmt.Scanln(&(u.MobileOrEmail))
	fmt.Println("Password:")
	fmt.Scanln(&(u.Password))

	m["UserLogin"] = u

	if err := c.SendCmd(m); err != nil {
		fmt.Println("Login Send Error:", err.Error())
	}
	return
}
func HandlerLogin_Recv(data interface{}) {
	fmt.Println("Login Success Token:", data.(string))
}
