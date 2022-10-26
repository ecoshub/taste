package example

import (
	"encoding/json"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	Users []*User = []*User{
		{ID: "a4fb4201", Name: "eco", Age: 30},
		{ID: "43bd1a0d", Name: "any", Age: 29},
	}
)

func GetUser(name string) (*User, bool) {
	for _, u := range Users {
		if u.Name == name {
			return u, true
		}
	}
	return nil, false
}

func AddUser(user *User) {
	Users = append(Users, user)
}

func (u *User) Marshal() []byte {
	return MarshalDiscardError(u)
}

func MarshalDiscardError(i interface{}) []byte {
	enc, _ := json.Marshal(i)
	return enc
}
