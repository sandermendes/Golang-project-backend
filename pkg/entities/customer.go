package entities

// Customer entity
type Customer struct {
	ID        uint   `gorm:"primaryKey"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// CustomerResponse entity
type CustomerResponse struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}
