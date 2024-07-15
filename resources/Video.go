package resources

import "github.com/evertras/bubble-table/table"

type Video struct {
	ID          string
	bucket      string
	title       string
	description string

	viewCount    uint64
	likeCount    uint64
	dislikeCount uint64
	commentCount uint64
}

func (v *Video) MarshalData() []byte {
	return nil
}

func (v *Video) ToRow() table.Row {
	return table.NewRow(table.RowData{})
}

func (v *Video) Save() {

}
