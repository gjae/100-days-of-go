package purchase

import (
	"context"
	"errors"
	coffeeco "example/coffeeco/internal"
	"example/coffeeco/internal/loyalty"
	"example/coffeeco/internal/payment"
	"example/coffeeco/internal/store"
	"fmt"
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

type StoreService interface {
	GetStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID) (float32, error)
}

type Service struct {
	cardService  CardChargeService
	purchaseRepo Repository
	storeService StoreService
}

var ErrPaymentMeansUnknow = errors.New("unknow payment type")
var ErrFailedToStorePurchase = errors.New("failed to store purchase")

func (s Service) CompletePurchase(ctx context.Context, storeID uuid.UUID, purchase *Purchase, coffeeBuxCard *loyalty.CoffeeBux) error {

	if err := purchase.validateAndEnrich(); err != nil {
		return err
	}

	if err := s.calculateStoreSpecificDiscount(ctx, storeID, purchase); err != nil {
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

	if err := s.purchaseRepo.Store(ctx, *purchase); err != nil {
		return ErrFailedToStorePurchase
	}

	if coffeeBuxCard != nil {
		coffeeBuxCard.AddStamp()
	}

	return nil
}

func (s *Service) calculateStoreSpecificDiscount(ctx context.Context, storeID uuid.UUID, purchase *Purchase) error {
	discount, err := s.storeService.GetStoreSpecificDiscount(ctx, storeID)
	if err != nil && err != store.ErrNoDiscount {
		return fmt.Errorf("failed to get discount: %w", err)
	}

	purchasePrice := purchase.total

	if discount > 0 {
		purchase.total = *purchasePrice.Multiply(int64(100 - discount))
	}

	return nil
}

func NewService(cardService CardChargeService, purchaseRepo Repository, storeService StoreService) *Service {
	return &Service{cardService: cardService, purchaseRepo: purchaseRepo, storeService: storeService}
}
