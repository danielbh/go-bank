package transactions

// 0
import (
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// 1
var (
	mu           sync.Mutex
	transactions = []*Transaction{}
)

// Register register route handlers with main router
// When * is put in front of a type, e.g. *string, it becomes part of the type declaration, so you can say "this variable holds a pointer to a string".

// 2
func Register(router *gin.RouterGroup) {
	router.GET("/balance/:accountNumber", handleGetBalance)
	router.POST("/withdraw", handleWithdraw)
	router.POST("/deposit", handleDeposit)
	router.GET("/", handleGetTransactions)
}

// 3
func handleDeposit(ctx *gin.Context) {
	transactionBody, err := validateTransactionRequest(ctx)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	createTransaction(transactionBody.Amount, 0, transactionBody.AccountNumber)
	balance := getBalance(transactionBody.AccountNumber)

	ctx.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}

func handleWithdraw(ctx *gin.Context) {

	transactionBody, err := validateTransactionRequest(ctx)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	createTransaction(0, transactionBody.Amount, transactionBody.AccountNumber)
	balance := getBalance(transactionBody.AccountNumber)

	ctx.JSON(http.StatusOK, gin.H{
		"balance": balance,
	})
}

// 4
func createTransaction(debit float64, credit float64, accountNumber int64) {
	created := int32(time.Now().Unix())

	// Store credits as negatives
	if credit > 0 {
		credit = -credit
	}

	// Avoid accidental copies
	newTransaction := &Transaction{
		Debit:         debit,
		Credit:        credit,
		AccountNumber: accountNumber,
		Created:       created,
	}

	mu.Lock()
	transactions = append(transactions, newTransaction)
	mu.Unlock()
}

/*
A lock occurs when multiple processes try to access the same resource at the same time.
One process loses out and must wait for the other to finish.
A deadlock occurs when the waiting process is still holding on to another resource that the first needs before it can finish.
*/

/*
A race condition occurs when two or more threads can access shared data and they try to change it at the same time. Because the thread scheduling algorithm can swap between threads at any time, you don't know the order in which the threads will attempt to access the shared data. Therefore, the result of the change in data is dependent on the thread scheduling algorithm, i.e. both threads are "racing" to access/change the data.

Problems often occur when one thread does a "check-then-act" (e.g. "check" if the value is X, then "act" to do something that depends on the value being X) and another thread does something to the value in between the "check" and the "act". E.g:

https://stackoverflow.com/questions/34510/what-is-a-race-condition
*/

// 5
func handleGetTransactions(ctx *gin.Context) {
	// TODO: only limit writes
	mu.Lock()
	defer mu.Unlock()
	// Avoid validation edge cases with this check
	if len(transactions) == 0 {
		ctx.JSON(200, gin.H{"transactions": transactions})
		return
	}

	// emphasize don't use outside development
	start, _ := strconv.ParseInt(ctx.Query("start"), 10, 0)
	finish, _ := strconv.ParseInt(ctx.Query("finish"), 10, 0)

	if start > int64(len(transactions)) {
		start = int64(len(transactions))
	}

	if finish > int64(len(transactions)) {
		finish = int64(len(transactions))
	}
	// https://play.golang.org/p/tfrjKguadPO//
	transactionSubset := transactions[start:finish]
	ctx.JSON(200, gin.H{"transactions": transactionSubset})
}

func handleGetBalance(ctx *gin.Context) {

	accountNumber, err := strconv.ParseInt(ctx.Param("accountNumber"), 10, 0)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	balance := getBalance(accountNumber)

	ctx.JSON(200, gin.H{"balance": balance})
}

// 5
func getBalance(accountNumber int64) float64 {
	creditsChan := make(chan float64)
	debitsChan := make(chan float64)

	go aggregateDebits(accountNumber, transactions, debitsChan)
	go aggregateCredits(accountNumber, transactions, creditsChan)

	credits, debits := <-creditsChan, <-debitsChan
	balance := debits + credits

	return balance
}

func aggregateDebits(accountNumber int64, transactions []*Transaction, ch chan float64) {
	var sum float64
	for _, transaction := range transactions {
		if transaction.AccountNumber == accountNumber {
			sum += transaction.Debit
		}
	}
	ch <- sum
}

func aggregateCredits(accountNumber int64, transactions []*Transaction, ch chan float64) {
	var sum float64
	for _, transaction := range transactions {
		if transaction.AccountNumber == accountNumber {
			sum += transaction.Credit
		}
	}
	ch <- sum
}
