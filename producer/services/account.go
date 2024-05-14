package services

import (
	"errors"
	"events"
	"github.com/google/uuid"
	"log"
	"producer/commands"
)

type AccountService interface {
	OpenAccountCommand(command commands.OpenAccountCommand) (id string, err error)
	DepositFundCommand(command commands.DepositFundCommand) error
	WithdrawFundCommand(command commands.WithdrawFundCommand) error
	CloseAccountCommand(command commands.CloseAccountCommand) error
}

type accountService struct {
	eventProducer EventProducer
}

func NewAccountService(eventProducer EventProducer) AccountService {
	return accountService{eventProducer}
}

func (obj accountService) OpenAccountCommand(command commands.OpenAccountCommand) (id string, err error) {

	if command.AccountHolder == "" || command.AccountType == 0 || command.OpeningBalance == 0 {
		return "", errors.New("bad request")
	}

	event := events.OpenAccountEvent{
		ID:             uuid.NewString(),
		AccountHolder:  command.AccountHolder,
		AccountType:    command.AccountType,
		OpeningBalance: command.OpeningBalance,
	}

	log.Printf("%#v", event)
	return event.ID, obj.eventProducer.Produce(event)
}

func (obj accountService) DepositFundCommand(command commands.DepositFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}

	event := events.DepositFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}

func (obj accountService) WithdrawFundCommand(command commands.WithdrawFundCommand) error {
	if command.ID == "" || command.Amount == 0 {
		return errors.New("bad request")
	}

	event := events.WithdrawFundEvent{
		ID:     command.ID,
		Amount: command.Amount,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}

func (obj accountService) CloseAccountCommand(command commands.CloseAccountCommand) error {
	if command.ID == "" {
		return errors.New("bad request")
	}

	event := events.WithdrawFundEvent{
		ID: command.ID,
	}

	log.Printf("%#v", event)
	return obj.eventProducer.Produce(event)
}
