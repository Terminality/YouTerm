package API

import "time"

type Video struct {
	ID              string
	watchStatus     bool
	uploader        string
	publishDateTime time.Time
}

func NewVideo() {

}
