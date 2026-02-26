package models

type PaymentMethod struct{
	PaymentMethodID int `json:"payment_id"`
	PaymentName string `json:"payment_name"`
}