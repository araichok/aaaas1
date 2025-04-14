package domain

type Product struct {
	ID          string  `json:"id" bson:"id"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
	Stock       int     `json:"stock" bson:"stock"`
	Category    string  `json:"category" bson:"category"`
}
