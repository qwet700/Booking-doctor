package models

type User struct {
	ID        string `json:"id"`
	FullName  string `json:"fullname"`
	Phone     string `json:"phone"`
	OrderNumb string `json:"ordernumb"`
}
