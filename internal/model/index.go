package model

type Index struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (i *Index) TableName() string {
	return "index"
}
