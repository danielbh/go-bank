package transactions

// Transaction Exported from package so that other modules can access
type Transaction struct {
	// Keys must be capitalized so JSON package can read them
	// struct tags
	Debit         float64 `json:"debit"`
	Credit        float64 `json:"credit"`
	AccountNumber int64   `json:"account_number"`
	Created       int32   `json:"created"`
}
