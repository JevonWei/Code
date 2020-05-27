package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Reciver interface {
	Reciver() (string, error)
}

type Client interface {
	Sender
	Reciver
	Open() error
	Close() error
}

type MSNClient struct{}

func (c MSNClient) Open() error {
	fmt.Println("Open")
	return nil
}

func (c MSNClient) Close() error {
	fmt.Println("Close")
	return nil
}

func (c MSNClient) Send(msg string) error {
	fmt.Println("send:", msg)
	return nil
}

func (c MSNClient) Reciver() (string, error) {
	fmt.Println("recive")
	return "", nil
}

func main() {
	msn := MSNClient{}

	var s Sender = msn
	var r Reciver = msn
	var c Client = msn

	s.Send("1")
	r.Reciver()

	c.Open()
	defer c.Close()

	c.Send("2")
	c.Reciver()

	var closer interface {
		Close() error
	}

	closer = msn
	closer.Close()
}
