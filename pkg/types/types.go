package types

// Money represent a sum of money in minimum units (cents, penny, dirams etc.).
type Money int64

// PaymentCategory represent the category in which the payments were made (auto, chemists, restaurants etc.).
type PaymentCategory string

// PaymentStatus introduce info about statuses.
type PaymentStatus string

// Predefined payment statuses.
const (
	PaymentStatusOk         PaymentStatus = "OK"
	PaymentStatusFail       PaymentStatus = "FAIL"
	PaymentStatusInProgress PaymentStatus = "INPROGRESS"
)

// Payment introduce info about payments.
type Payment struct {
	ID        string
	AccountID int64
	Amount    Money
	Category  PaymentCategory
	Status    PaymentStatus
}

type Phone string

// Account introduce info about bill of users
type Account struct {
	ID      int64
	Phone   Phone
	Balance Money
}
