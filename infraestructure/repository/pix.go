package repository

import (
	"github.com/caiquejjx/codepix/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (r PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := r.Db.Create(bank).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDb) AddAcount(account *model.Account) error {
	err := r.Db.Create(account).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) error {
	err := r.Db.Create(pixKey).Error

	if err != nil {
		return err
	}

	return nil
}

func (r PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	result := r.Db.Preload("Bank").First(&account, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &account, nil
}

func (r PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	result := r.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if result.Error != nil {
		return nil, result.Error
	}

	return &pixKey, nil
}

func (r PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	result := r.Db.First(&bank, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &bank, nil
}
