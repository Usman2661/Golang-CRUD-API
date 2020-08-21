package models

// Todo Struct
type Todo struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"  gorm:"type: boolean; default:false"`
}
