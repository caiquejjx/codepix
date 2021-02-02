package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

//PixKeyRepositoryInterface is the pixkey repo interface
type PixKeyRepositoryInterface interface {
	RegisterKey(PixKey *PixKey) (*PixKey, error)
	FindKeyByKind(key string, kind string) (*PixKey, error)
	AddBank(bank *Bank) error
	AddAcount(account *Account) error
	FindAccount(id string) (*Account, error)
}

//PixKey is the pixkey struct
type PixKey struct {
	Base      `valid:"require"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	AccountID string   `json:"acccount_id" valid:"notnull"`
	Account   *Account `valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {

	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != "email" && pixKey.Kind != "cpf" {
		return errors.New("invalide key type")
	}

	if pixKey.Status != "active" && pixKey.Status != "inactive" {
		return errors.New("invalide status")
	}

	if err != nil {
		return err
	}
	return nil
}

//NewPixKey is the pixKey creation func
func NewPixKey(account *Account, kind string, key string) (*PixKey, error) {
	pixkey := PixKey{
		Account: account,
		Kind:    kind,
		Key:     key,
		Status:  "active",
	}

	pixkey.ID = uuid.NewV4().String()
	pixkey.CreatedAt = time.Now()

	err := pixkey.isValid()

	if err != nil {
		return nil, err
	}

	return &pixkey, nil
}
