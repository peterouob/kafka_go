package events

import "reflect"

var Topics = []string{
	reflect.TypeOf(OpenAccountEvent{}).Name(),
	reflect.TypeOf(DepositFundEvent{}).Name(),
	reflect.TypeOf(WithdrawFundEvent{}).Name(),
	reflect.TypeOf(CloseAccountEvent{}).Name(),
}

type Event interface{}

type OpenAccountEvent struct {
	ID             string  `json:"ID"`
	AccountHolder  string  `json:"accountHolder"`
	AccountType    int     `json:"accountType"`
	OpeningBalance float64 `json:"openingBalance"`
}

type DepositFundEvent struct {
	ID     string  `json:"ID"`
	Amount float64 `json:"amount"`
}

type WithdrawFundEvent struct {
	ID     string  `json:"ID"`
	Amount float64 `json:"amount"`
}

type CloseAccountEvent struct {
	ID string `json:"ID"`
}
