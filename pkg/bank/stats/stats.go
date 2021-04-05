package stats

import (
	"github.com/Alexanders95/bank/pkg/bank/types"
)

// Avg calculate avgerage summ to a payment
func Avg(payments []types.Payment) (types.Money) {

	var sum = types.Money(0)
	for _, card := range payments {

		sum += card.Amount
	}
	return sum / types.Money(len(payments))
}

// TotalInCategory calculate summ payments for category
func TotalInCategory(payments []types.Payment, category types.Category) (types.Money) {
	sum := types.Money(0)

	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}
 