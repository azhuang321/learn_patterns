package example1

import "fmt"

type ChatRoom struct{}

func (c *ChatRoom) ShowMessage(user *User, message string) {
	fmt.Printf("chat room %s send %s", user.name, message)
}
