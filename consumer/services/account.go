package services

import (
	"consumer/repositories"
	"encoding/json"
	"events"
	"log"
	"reflect"
)

type EventHandler interface {
	Handle(topic string, eventBytes []byte)
}

type accountEventHandler struct {
	accountRepo repositories.AccountRepository
}

func NewAccountEventHandler(accountRepo repositories.AccountRepository) EventHandler {
	return accountEventHandler{accountRepo}
}

func (obj accountEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(events.OpenAccountEvent{}).Name():
		event := &events.OpenAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Fatalln("json unmarshal error : ", err)
			return
		}
		bankAccount := repositories.BankAccount{
			ID:            event.ID,
			AccountHolder: event.AccountHolder,
			AccountType:   event.AccountType,
			Balance:       event.OpeningBalance,
		}
		if err = obj.accountRepo.Save(bankAccount); err != nil {
			log.Println("cannot save the bank account : ", err)
			return
		}
		log.Printf("[%v] %v", topic, event)

	case reflect.TypeOf(events.DepositFundEvent{}).Name():
		event := &events.DepositFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Fatalln("json unmarshal error : ", err)
			return
		}
		bankAccount, err := obj.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Fatalln("error to find account by id : ", err)
			return
		}
		bankAccount.Balance += event.Amount
		if err = obj.accountRepo.Save(bankAccount); err != nil {
			log.Println("cannot save the bank account : ", err)
			return
		}
		log.Printf("[%v] %v", topic, event)

	case reflect.TypeOf(events.WithdrawFundEvent{}).Name():
		event := &events.WithdrawFundEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Fatalln("json unmarshal error : ", err)
			return
		}
		bankAccount, err := obj.accountRepo.FindByID(event.ID)
		if err != nil {
			log.Fatalln("error to find account by id : ", err)
			return
		}
		bankAccount.Balance -= event.Amount
		if err = obj.accountRepo.Save(bankAccount); err != nil {
			log.Println("cannot save the bank account : ", err)
			return
		}
		log.Printf("[%v] %v", topic, event)

	case reflect.TypeOf(events.CloseAccountEvent{}).Name():
		event := &events.CloseAccountEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			log.Fatalln("json unmarshal error : ", err)
			return
		}
		if err = obj.accountRepo.Delete(event.ID); err != nil {
			log.Fatalln("error to delete account : ", err)
			return
		}
		log.Printf("[%v] %v", topic, event)
	default:
		log.Println("no event handler :" + topic)
	}
}
