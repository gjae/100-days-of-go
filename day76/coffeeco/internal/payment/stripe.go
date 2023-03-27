package payment

import (
	"context"
	"errors"
	"fmt"

	"github.com/Rhymond/go-money"
	"github.com/stripe/stripe-go/v73"
	"github.com/stripe/stripe-go/v73/charge"
	"github.com/stripe/stripe-go/v73/client"
)

type StripeService struct {
	stripeClient *client.API
}

var ErrStripeApiKeyNotFound = errors.New("API key cannot be nil")

func NewStripeService(apiKey string) (*StripeService, error) {
	if apiKey == "" {
		return nil, ErrStripeApiKeyNotFound
	}

	sc := &client.API{}

	sc.Init(apiKey, nil)

	return &StripeService{stripeClient: sc}, nil
}

func (s StripeService) ChargeCard(ctx context.Context, amount money.Money, cardToken string) error {
	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(amount.Amount()),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Source:   &stripe.PaymentSourceSourceParams{Token: stripe.String(cardToken)},
	}

	_, err := charge.New(params)

	if err != nil {
		return fmt.Errorf("ffailed to create charge: %v", err)
	}

	return nil
}
