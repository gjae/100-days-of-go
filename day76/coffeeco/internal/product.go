package coffeeco

import (
	"errors"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Product struct {
	ItemName  string
	BasePrice money.Money
}

var ErrPurchaseQuantity = errors.New("purchase must consist of at least one product")
var ErrTotalIsZero = errors.New("likely mistake; purchase should never be 0. Please validate")

func (p *Purchase) validateAndEnrich() error {

	if len(p.ProductToPurchase) == 0 {
		return ErrPurchaseQuantity
	}

	p.Total = *money.New(0, "USD")
	for _, v := range p.ProductsToPurchase {
		newTotal, _ := p.total.Add(&v.BasePrice)

		p.total = *newTotal
	}

	if p.total.IsZero() {
		return ErrTotalIsZero
	}

	p.id = uuid.New()

	p.timeOfPurchase = time.Now()

	return nil
}
