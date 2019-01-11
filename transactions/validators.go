package transactions

import (
	"github.com/gin-gonic/gin"
)

// TransactionRequestBody used to validate transactions
type TransactionRequestBody struct {
	Amount        float64 `form:"amount" binding:"required,min=0"`
	AccountNumber int64   `form:"account_number" binding:"required"`
}

func validateTransactionRequest(ctx *gin.Context) (*TransactionRequestBody, error) {
	// memory efficiency in speed by passing a reference
	transactionBody := &TransactionRequestBody{}

	if err := ctx.ShouldBind(transactionBody); err != nil {
		return nil, err
	}

	return transactionBody, nil
}
