package loyalty

import (
	"context"
	"errors"
	"example/coffeeco/internal/purchase"
	"fmt"
)

func (c *CoffeeBux) Pay(ctx context.Context, purchases []purchase.Purchase) error {
	lp := len(purchases)

	if lp == 0 {
		return errors.New("nothing to buy")
	}

	if c.FreeDrinksAvailable < lp {
		return fmt.Errorf("not enough coffeebux to cover entire purchase. Have %d, need %d,", len(purchases), c.FreeDrinksAvailable)
	}

	c.FreeDrinksAvailable = c.FreeDrinksAvailable - lp

	return nil
}
