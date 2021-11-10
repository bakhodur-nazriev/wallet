package types

// Money represent a sum of money in minimum units (cents, penny, dirams etc.).
type Money int64

// Currency represent code currency.
type Currency string

// PAN represent amount numbers of card.
type PAN string

// Category represnt the category in which the payments was made (auto, chemists, restaurants etc.).
type Category string

// Status introduce info about statuses.
type Status string

// Predefined payment statuses.
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Code currency.
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

// Payment introduce info about payments.
type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

// Card introduce info about cards.
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}
