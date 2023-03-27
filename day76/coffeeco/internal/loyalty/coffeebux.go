package loyalty

import (
	"github.com/google/uuid"
	coffeeco "example/coffeeco/internal"
	"example/coffeeco/internal/store"
)

type CoffeeBux struct {
	ID uuid.UUID
	store store.Store
	coffeeLover coffeeco.CoffeeLover
	FreeDrinksAvailable int
	RemainingDrinkPurchasesUntilFreeDrink int
}
