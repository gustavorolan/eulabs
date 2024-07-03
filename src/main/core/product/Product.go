package product

type Product struct {
	ID          int     `gorm:"primaryKey" json:"id"`
	Name        string  `gorm:"unique;not null" json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}
