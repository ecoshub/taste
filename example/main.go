package example

import (
	"encoding/json"
	"math/rand"
)

const (
	_IDChars      string = "01234567890abcdef"
	_DefaultIDlen int    = 8
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	Users []*User = []*User{
		{ID: generateID(), Name: "eco", Age: 30},
		{ID: generateID(), Name: "any", Age: 29},
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

func generateID() string {
	idArr := make([]byte, _DefaultIDlen)
	for i := range idArr {
		index := rand.Intn(len(_IDChars))
		idArr[i] = _IDChars[index]
	}
	return string(idArr)
}
