package bank

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

// Deposit updated the deposit value
func Deposit(amount int) { deposits <- amount }

// Balance reads the balance
func Balance() int { return <-balances }

// teller is the monitor goroutine
func teller() {
	var balance int // balance is confined to teller goroutine

	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}
