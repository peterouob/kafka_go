package commands

type OpenAccountCommand struct {
	AccountHolder  string  `json:"accountHolder,omitempty"`
	AccountType    int     `json:"accountType,omitempty"`
	OpeningBalance float64 `json:"openingBalance,omitempty"`
}

type DepositFundCommand struct {
	ID     string  `json:"ID,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}
type WithdrawFundCommand struct {
	ID     string  `json:"ID,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}
type CloseAccountCommand struct {
	ID string `json:"ID,omitempty"`
}
