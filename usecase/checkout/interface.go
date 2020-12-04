package checkout

import (
	"github.com/vitoraalmeida/silkroad/entity"
)

type UseCase interface {
	Checkout(c Cart, customerID uint) (entity.Sale, error)
}
