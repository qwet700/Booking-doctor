package model

type User struct {
	ID        string  `json:"id"`
	FullName  string  `json:"fullname"`
	Phone     float64 `json:"phone"`
	OrderNumb string  `json:"ordernumb"`
}
