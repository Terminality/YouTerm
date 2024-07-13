package models

type Resource struct {
	ID     string
	bucket string
}

func (r *Resource) Save() {}

func (r *Resource) ToString() {}
