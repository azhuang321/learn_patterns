package example1

type User struct {
	name string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}
func NewUser(name string) *User {
	user := &User{}
	user.name = name
	return user
}

func (u *User) SendMessage(message string) {
	chatroom := &ChatRoom{}
	chatroom.ShowMessage(u, message)
}
