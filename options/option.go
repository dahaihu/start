package main

import "fmt"

type Client struct {
	Host string
}

type ClientSetting func(*Client)

func ClientWithHoust(host string) ClientSetting {
	return func(c *Client) {
		c.Host = host
	}
}

func NewClient(ops ...ClientSetting) *Client {
	client := &Client{}
	for _, op := range ops {
		op(client)
	}
	return client
}

func main() {
	client := NewClient(ClientWithHoust("www.baidu.com"))
	fmt.Println("client is ", *client)
}
