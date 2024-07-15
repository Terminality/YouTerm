package models

import (
	"encoding/json"
)

type List struct {
	ID     string
	Bucket string
}

func (l *List) GetID() string         { return l.ID }
func (l *List) GetBucketName() string { return l.Bucket }

func (l *List) MarshalData() ([]byte, error) {
	return json.Marshal(l)
}

func (l *List) Save() {

}
