package payment

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/rulanugrh/uranus/configs"
	"github.com/rulanugrh/uranus/internal/domain/entity"
	"github.com/rulanugrh/uranus/internal/repository/port"
	portthirdparty "github.com/rulanugrh/uranus/third_party/midtrans/port"
)

type payment struct {
	user   port.UserInterfaceRepository
	order  port.OrderInterfaceRepository
	course port.CourseInterfaceRepository
}

func NewPayment(user port.UserInterfaceRepository, order port.OrderInterfaceRepository, course port.CourseInterfaceRepository) portthirdparty.PortSandbox {
	return &payment{
		user:   user,
		order:  order,
		course: course,
	}
}

func (pay *payment) Checkout(id uint, order entity.Order) (*entity.PaymentSandbox, error) {
	course, errFound := pay.course.FindById(uint(order.CourseID))
	if errFound != nil {
		return nil, errFound
	}

	user, errUser := pay.user.FindByID(uint(order.UserID))
	if errUser != nil {
		return nil, errUser
	}

	conf := configs.GetConfig()
	res, err := paymentCharge(conf.Midtrans.Sandbox, *course, id, *user, order)
	if err != nil {
		return nil, err
	}

	return res, nil

}

func paymentCharge(serverkey string, course entity.Course, id uint, user entity.User, order entity.Order) (*entity.PaymentSandbox, error) {
	configs.Core.New(serverkey, midtrans.Sandbox)
	chargeReq := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("%d", id),
			GrossAmt: int64(course.Price),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: user.Name,
			Email: user.Email,
			Phone: user.Address,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:       fmt.Sprintf("%d", course.ID),
				Price:    int64(course.Price),
				Qty:      int32(order.Quantity),
				Name:     course.Name,
				Category: course.Categories.Name,
			},
		},
		EnabledPayments: snap.AllSnapPaymentType,
	}

	res, errCreate := snap.CreateTransaction(chargeReq)
	if errCreate != nil {
		return nil, errCreate.RawError
	}

	// get token for transactions
	token, err := snap.CreateTransactionToken(chargeReq)
	if err != nil {
		return nil, err.RawError
	}

	res.Token = token

	var items []entity.ItemDetail
	var paymentsType []entity.PaymentTypes

	for _, data := range *chargeReq.Items {
		item := entity.ItemDetail{
			ID: data.ID,
			Name: data.Name,
			Brand: data.Brand,
			Qty: data.Qty,
			Price: data.Price,
			Category: data.Category,
			MerchantName: data.MerchantName,
		}

		items = append(items, item)
	}

	copier.Copy(&paymentsType, &chargeReq.EnabledPayments)

	pay := entity.PaymentSandbox {
		TransactionDetails: entity.TransactionDetail{
			OrderID: chargeReq.TransactionDetails.OrderID,
			GrossAmount: chargeReq.TransactionDetails.GrossAmt,
		},
		CustomerDetails: entity.CustomerDetail{
			Name: chargeReq.CustomerDetail.FName,
			Email: chargeReq.CustomerDetail.Email,
			Address: chargeReq.CustomerDetail.Phone,
		},
		ItemsDetails: items,
		PaymentType: paymentsType,
	}
	
	return &pay, nil
}
