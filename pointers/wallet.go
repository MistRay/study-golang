package pointers

import (
	"errors"
	"fmt"
)

// 比特币
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// 某人钱包中拥有比特币的数量
type wallet struct {
	balance Bitcoin
}

// 存款操作 将一些比特币存入钱包中
func (w *wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// 钱包中没有足够的比特币来提款
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// 提现操作 将一些比特币从钱包中扣除,如果无法执行则返回错误
func (w *wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// 查看余额
func (w *wallet) Balance() Bitcoin {
	return w.balance
}
