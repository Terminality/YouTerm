package models

import (
	"encoding/json"

	"dalton.dog/YouTerm/modules/Storage"
)

type User struct {
	ID         string
	lists      map[string]List
	SubbedList []string
}

func (u *User) GetID() string         { return u.ID }
func (u *User) GetBucketName() string { return "Users" }

func (u *User) MarshalData() ([]byte, error) {
	return json.Marshal(u)
}

func CreateUser(ID string) *User {
	bytes := Storage.LoadResource("Lists", "Subscribed")

	var list []string

	json.Unmarshal(bytes, &list)

	return &User{
		ID:         ID,
		SubbedList: list,
	}
}

func (u *User) Save() {

}
