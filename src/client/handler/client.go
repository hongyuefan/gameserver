package handler

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"server/msg"

	"github.com/name5566/leaf/util"
)

type Client struct {
	conn     net.Conn
	protol   string
	addr     string
	chanSend chan []byte
	chanExit chan bool
	funcMap  *util.Map
}

func NewClient(protol, addr string) *Client {
	return &Client{
		protol:   protol,
		addr:     addr,
		chanSend: make(chan []byte, 100),
		chanExit: make(chan bool, 0),
		funcMap:  new(util.Map),
	}
}

func (c *Client) OnInit() (err error) {
	c.conn, err = net.Dial(c.protol, c.addr)
	if err != nil {
		panic(err)
	}
	return
}

func (c *Client) SendCmd(s interface{}) (err error) {
	var data []byte
	if data, err = json.Marshal(s); err != nil {
		return
	}
	select {
	case c.chanSend <- data:
	default:
		return errors.New("send chan has full")
	}
	return
}

func (c *Client) RegistFunc(key msg.BussTypeId, f CallFunc) {
	c.funcMap.Set(key, f)
}

func (c *Client) OnStart() {

	go func() {
		for {
			select {
			case b := <-c.chanSend:
				if err := c.sendMsg(b); err != nil {
					fmt.Println("Send Msg Error :", err.Error())
				}
			case <-c.chanExit:
				return
			}
		}

	}()

	for {
		var (
			b   []byte
			err error
		)
		select {
		case <-c.chanExit:
			return
		default:
			if b, err = c.recvMsg(); err != nil {
				if err == io.EOF {
					c.OnClose()
					fmt.Println("Server Closed Please Enter 'q' to Exit")
					return
				}
				fmt.Println("Recv Msg Error:", err.Error())
				continue
			}
			c.handlerMsg(b)
		}

	}
}

func (c *Client) OnClose() {
	select {
	case <-c.chanExit:
	default:
		close(c.chanExit)
		c.conn.Close()
	}
}

func (c *Client) sendMsg(b []byte) (err error) {

	m := make([]byte, 2+len(b))

	binary.BigEndian.PutUint16(m, uint16(len(b)))

	copy(m[2:], b)

	if _, err = c.conn.Write(m); err != nil {
		fmt.Println("Send Msg Error:", err.Error())
	}
	return
}

func (c *Client) recvMsg() (b []byte, err error) {

	var rByte []byte

	rByte = make([]byte, 2)

	if _, err = io.ReadFull(c.conn, rByte); err != nil {
		return
	}

	len := binary.BigEndian.Uint16(rByte)

	b = make([]byte, len)

	if _, err = io.ReadFull(c.conn, b); err != nil {
		return
	}
	return
}

func (c *Client) handlerMsg(b []byte) {

	var err error

	m := make(map[string]*Response, 0)

	if err = json.Unmarshal(b, &m); err != nil {
		fmt.Println("Json Unmarshal Error:", err.Error())
		return
	}

	rsp := m["Response"]

	if !rsp.Success {
		fmt.Println("some errors:", rsp.Message)
		return
	}

	f := c.funcMap.Get(rsp.BussId)
	if f == nil {
		fmt.Println("Function ", rsp.BussId, "not Regist")
	}

	f.(CallFunc)(rsp.Data)

	return
}
