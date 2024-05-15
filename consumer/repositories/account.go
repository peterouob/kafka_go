package repositories

import "gorm.io/gorm"

type BankAccount struct {
	ID            string  `json:"ID"`
	AccountHolder string  `json:"accountHolder"`
	AccountType   int     `json:"accountType"`
	Balance       float64 `json:"balance"`
}

type AccountRepository interface {
	Save(bankAccount BankAccount) error
	Delete(id string) error
	FindAll() (bankAccounts []BankAccount, err error)
	FindByID(id string) (bankAccount BankAccount, err error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	db.Table("bond_banks").AutoMigrate(&BankAccount{})
	return accountRepository{db: db}
}

func (obj accountRepository) Save(bankAccount BankAccount) error {
	return obj.db.Table("bond_banks").Save(bankAccount).Error
}

func (obj accountRepository) Delete(id string) error {
	return obj.db.Table("bond_banks").Where("id=?", id).Delete(&BankAccount{}).Error
}

func (obj accountRepository) FindAll() (bankAccounts []BankAccount, err error) {
	err = obj.db.Table("bond_banks").Find(&bankAccounts).Error
	return bankAccounts, err
}

func (obj accountRepository) FindByID(id string) (bankAccount BankAccount, err error) {
	err = obj.db.Table("bond_banks").Where("id=?", id).First(&bankAccount).Error
	return bankAccount, err
}
