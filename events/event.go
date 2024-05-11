package event

import "reflect"

var Topics = []string{
	reflect.TypeOf(OpenAccountEvent{}).Name(),
	reflect.TypeOf(DepositFundEvent{}).Name(),
	reflect.TypeOf(WithdrawFunEvent{}).Name(),
	reflect.TypeOf(CloseAccountEvent{}).Name(),
}

type OpenAccountEvent struct {
	ID             string
	AccountHolder  string
	AccountType    int
	OpeningBalance float64
}

type DepositFundEvent struct {
	ID     string
	Amount float64
}

type WithdrawFunEvent struct {
	ID string
	Amount float64
}

type CloseAccountEvent struct {
	ID string
}
