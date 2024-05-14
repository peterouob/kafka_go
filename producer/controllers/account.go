package controllers

import (
	"github.com/gofiber/fiber/v2"
	"producer/commands"
	"producer/services"
)

type AccountControllers interface {
	OpenAccount(c *fiber.Ctx) error
	DepositFund(c *fiber.Ctx) error
	WithdrawFund(c *fiber.Ctx) error
	CloseAccount(c *fiber.Ctx) error
}

type accountControllers struct {
	accountService services.AccountService
}

func NewAccountControllers(accountService services.AccountService) AccountControllers {
	return accountControllers{accountService}
}

func (obj accountControllers) OpenAccount(c *fiber.Ctx) error {
	command := commands.OpenAccountCommand{}
	if err := c.BodyParser(&command); err != nil {
		return err
	}
	id, err := obj.accountService.OpenAccountCommand(command)
	if err != nil {
		return err
	}
	c.Status(fiber.StatusCreated)

	return c.JSON(fiber.Map{
		"Message": "success create account",
		"id":      id,
	})
}

func (obj accountControllers) DepositFund(c *fiber.Ctx) error {
	command := commands.DepositFundCommand{}
	if err := c.BodyParser(&command); err != nil {
		return err
	}
	if err := obj.accountService.DepositFundCommand(command); err != nil {
		return err
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"Message": "deposit fund success ",
	})
}
func (obj accountControllers) WithdrawFund(c *fiber.Ctx) error {
	command := commands.WithdrawFundCommand{}
	if err := c.BodyParser(&command); err != nil {
		return err
	}
	if err := obj.accountService.WithdrawFundCommand(command); err != nil {
		return err
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"Message": "withdraw fund success ",
	})
}
func (obj accountControllers) CloseAccount(c *fiber.Ctx) error {
	command := commands.CloseAccountCommand{}
	if err := c.BodyParser(&command); err != nil {
		return err
	}
	if err := obj.accountService.CloseAccountCommand(command); err != nil {
		return err
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"Message": "close account success ",
	})
}
