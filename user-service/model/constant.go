package model

// User struct
type User struct {
	ID      string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string  `json:"name"`
	Product string  `json:"product"`
	Age     float64 `json:"age"`
}

const (
	dbName    = "awesomeApp"
	tableName = "users"
)
