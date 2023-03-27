package purchase

import (
	"context"
	"errors"
	coffeeco "example/coffeeco/internal"
	"example/coffeeco/internal/payment"
	"example/coffeeco/internal/store"
	"time"

	"github.com/Rhymond/go-money"
	"github.com/google/uuid"
)

type Purchase struct {
	ID                 uuid.UUID
	Store              store.Store
	ProductsToPurchase []coffeeco.Product
	total              money.Money
	PaymentMeans       payment.Means
	timeOfPurchase     time.Time
	CardToken          *string
}

type CardChargeService interface {
	ChargeCard(ctx context.Context, amout money.Money, cardTGoken string) error
}

type Service struct {
	cardService  CardChargeService
	purchaseRepo Repository
}

var ErrPaymentMeansUnknow = errors.New("unknow payment type")

func (s Service) CompletePurchase(ctx context.Context, purchase *Purchase) error {

	if err := purchase.validateAndEnrich(); err != nil {
		return err
	}

	switch purchase.PaymentMeans {
	case payment.MEANS_CARD:
		if err := s.cardService.ChargeCard(ctx, purchase.total, *purchase.CardToken); err != nil {
			return err
		}
	case payment.MEANS_CASH:
		// TODO: for the reader to add
	default:
		return ErrPaymentMeansUnknow
	}

	if err := s.purchaseRepo.Store(ctx, *purchase); err != nil {
		return err
	}

	return nil
}
