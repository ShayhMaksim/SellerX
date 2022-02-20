package users

import(
	"error"
)

type Users struct {
	id int
	balance float64
} 

func (d * Users) SetBalance(new_balance float64) error {
	if (new_balance<0.) {
		return errors.New("Balance can't be a negative number")
	}
	d.balance = new_balance
	return nil
}