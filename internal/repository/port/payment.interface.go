package port

import "github.com/rulanugrh/uranus/internal/domain/entity"

type PaymentRepository interface {
	Save(req entity.PaymentSandbox) (*entity.PaymentSandbox, error)
	History(id uint) (*entity.PaymentSandbox, error)
}
