package models

type User struct {
	UUID       string
	Username   string
	Email      string
	Password   string
	Permission int
}

type ShoppingList struct {
	UUID    string
	Name    string
	OwnerID int
	Shared  bool
}
