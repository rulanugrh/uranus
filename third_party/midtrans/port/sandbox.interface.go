package portthirdparty

import "github.com/rulanugrh/uranus/internal/domain/entity"

type PortSandbox interface {
	Checkout(id uint, order entity.Order) (*entity.PaymentSandbox, error)
	History(id uint) (*entity.PaymentSandbox, error)
}
