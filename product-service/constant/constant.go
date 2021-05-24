package constant

// Product struct
type Product struct {
	ID       string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
}

const (
	dbName    = "awesomeApp"
	tableName = "products"
)
