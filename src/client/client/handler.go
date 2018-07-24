package client

import (
	"fmt"
)

func HandlerGetGameClass_Send(c *Client) {
	m := make(map[string]interface{}, 1)

	u := new(GameClass)

	m["GameClass"] = u

	if err := c.SendCmd(m); err != nil {
		fmt.Println("GetGameClass Send Error:", err.Error())
	}
	return
}
func HandlerGetGameClass_Recv(data interface{}) {
	fmt.Println("GetGameClass:", data)
}

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

func HandlerChatLogin_Send(c *Client) {
	m := make(map[string]interface{}, 1)
	t := new(ChatLogin)
	fmt.Println("Token:")
	fmt.Scanln(&(t.Token))

	m["ChatLogin"] = t

	if err := c.SendCmd(m); err != nil {
		fmt.Println("Login Send Error:", err.Error())
	}
	return
}

func HandlerChat_Send(c *Client) {

	m := make(map[string]interface{}, 1)

	t := new(ChatTo)
	fmt.Println("Msg:")
	fmt.Scanln(&(t.Msg))
	t.To_Ids = []int64{1}
	t.Name = "fhy"

	m["ChatTo"] = t

	if err := c.SendCmd(m); err != nil {
		fmt.Println("Login Send Error:", err.Error())
	}
	return
}
func HandlerChat_Recv(data interface{}) {
	fmt.Println("Chat Success :", data.(string))
}
