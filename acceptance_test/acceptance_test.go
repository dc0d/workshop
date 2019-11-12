package acceptance

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	startServer()

	exitVal := m.Run()

	os.Exit(exitVal)
}

func Test_statement(t *testing.T) {
	clientID := generateID()
	client := newClient(clientID)
	var bankStatement string

	Convey("Given a client", t, func() {
		Convey("When the client makes a deposit", func() { client.deposit(1000, depositDate1) })

		Convey("And makes another deposit", func() { client.deposit(2000, depositDate2) })

		Convey("And makes a withdrawal", func() { client.withdraw(500, withdrawDate1) })

		Convey("When the client prints the bank statement", func() { bankStatement = client.bankStatement() })

		Convey("Then the client sees the bank statement", func() {
			So(bankStatement, ShouldEqual, expectedBankStatement)
		})
	})
}

var (
	depositDate1          = parseDate("10-01-2012")
	depositDate2          = parseDate("13-01-2012")
	withdrawDate1         = parseDate("14-01-2012")
	expectedBankStatement = `date || credit || debit || balance
14/01/2012 || || 500.00 || 2500.00
13/01/2012 || 2000.00 || || 3000.00
10/01/2012 || 1000.00 || || 1000.00`
)
