package entity

type PaymentSandbox struct {
	StatusCode         string            `json:"status_code"`
	Token              string            `json:"token"`
	RedirectURL        string            `json:"redirect_url"`
	TransactionDetails TransactionDetail `json:"transcation_details"`
	CustomerDetails    CustomerDetail    `json:"customer_details"`
	ItemsDetails       []ItemDetail      `json:"items"`
	PaymentType        []PaymentTypes    `json:"payment_type"`
}

type TransactionDetail struct {
	OrderID     string `json:"order_id"`
	GrossAmount int64  `json:"gross_amount"`
}

type CustomerDetail struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

type ItemDetail struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Qty          int32  `json:"quantity"`
	Brand        string `json:"brand,omitempty"`
	Category     string `json:"category,omitempty"`
	MerchantName string `json:"merchant_na me,omitempty"`
}

type PaymentTypes string