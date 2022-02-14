package repository

import "wallet-service/model"

type IWalletRepository interface {
	CreateWallet(string, int)
	ChangeBalance(string, int)
	GetAll() *map[string]int
	GetBalance(string) model.ResponseModel
	Exists(string) bool
}

type WalletRepository struct {
	db map[string]int
}

func (wr *WalletRepository) CreateWallet(id string, amount int) {
	wr.db[id] = amount
}

func (wr *WalletRepository) ChangeBalance(id string, amount int) {
	wr.db[id] = amount
}

func (wr *WalletRepository) GetAll() *map[string]int {
	return &wr.db
}

func (wr *WalletRepository) GetBalance(username string) model.ResponseModel {
	return model.ResponseModel{
		Username: username,
		Balance:  wr.db[username],
	}
}

func (wr *WalletRepository) Exists(username string) bool {
	_, exists := wr.db[username]
	return exists
}

func NewWalletRepository() IWalletRepository {
	return &WalletRepository{db: getMockData()}
}

func getMockData() map[string]int {
	return map[string]int{
		"Furkan": 4500,
		"Said":   20000,
		"Ahmet":  0,
	}
}
