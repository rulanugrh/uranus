package entity

import (
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentSandbox struct {
	StatusCode         string                      `json:"status_code"`
	Token              string                      `json:"token"`
	RedirectURL        string                      `json:"redirect_url"`
	TransactionDetails midtrans.TransactionDetails `json:"transcation_details"`
	CustomerDetails    *midtrans.CustomerDetails   `json:"customer_details"`
	ItemsDetails       *[]midtrans.ItemDetails     `json:"items"`
	PaymentType        []snap.SnapPaymentType      `json:"payment_type"`
}
