package resources

import (
	"encoding/json"
	"log"

	"dalton.dog/YouTerm/modules/Storage"
)

type User struct {
	ID         string
	Bucket     string
	UserLists  map[string]map[string]bool
	ActiveList string
	apiKey     string
}

func (u *User) GetID() string                { return u.ID }
func (u *User) GetBucketName() string        { return u.Bucket }
func (u *User) MarshalData() ([]byte, error) { return json.Marshal(u) }

func (u *User) AddToList(listName string, ID string) bool {
	log.Printf("Adding ID (%v) to list %v\n", ID, listName)
	list := u.UserLists[listName]
	if list != nil {
		list[ID] = true
		u.UserLists[listName] = list
		return true
	}
	return false
}

func (u *User) RemoveFromList(listName string, ID string) {
	log.Printf("Removing ID (%v) from list %v\n", ID, listName)
	list := u.UserLists[listName]
	if list != nil {
		delete(list, ID)
	}
}

func (u *User) GetList(listName string) map[string]bool {
	log.Printf("Loading list %v\n", listName)
	return u.UserLists[listName]
}

func LoadOrCreateUser(ID string) *User {
	bytes := Storage.LoadResource(Storage.USERS, ID)

	if bytes == nil {
		return NewUser(ID)
	}

	var user *User
	json.Unmarshal(bytes, &user)
	return user
}

func NewUser(ID string) *User {
	log.Printf("Creating User -- %v\n", ID)
	userLists := make(map[string]map[string]bool)

	userLists[SUBBED_TO] = map[string]bool{}
	userLists[WATCH_LATER] = map[string]bool{}

	newUser := &User{
		ID:         ID,
		Bucket:     Storage.USERS,
		UserLists:  userLists,
		ActiveList: SUBBED_TO,
	}

	Storage.SaveResource(newUser)
	return newUser
}
