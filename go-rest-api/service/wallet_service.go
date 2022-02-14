package service

import (
	"errors"
	"wallet-service/config"
	"wallet-service/model"
	"wallet-service/repository"
)

type IWalletService interface {
	GetAll() (*map[string]int, error)
	Get(string) (model.ResponseModel, error)
	CreateWallet(string) error
	ChangeBalance(string, int) (model.ResponseModel, error)
}

type WalletService struct {
	walletRepository repository.IWalletRepository
}

func (ws *WalletService) GetAll() (*map[string]int, error) {
	response := ws.walletRepository.GetAll()
	return response, nil
}

func (ws *WalletService) Get(username string) (model.ResponseModel, error) {
	if exists := ws.walletRepository.Exists(username); !exists {
		return model.ResponseModel{}, errors.New("wallet does not exists")
	}

	return ws.walletRepository.GetBalance(username), nil
}

func (ws *WalletService) CreateWallet(username string) error {
	if exists := ws.walletRepository.Exists(username); exists {
		return errors.New("wallet already exists")
	}

	ws.walletRepository.CreateWallet(username, config.GetConfig().InitialAmount)
	return nil
}

func (ws *WalletService) ChangeBalance(username string, amount int) (model.ResponseModel, error) {
	if exists := ws.walletRepository.Exists(username); !exists {
		return model.ResponseModel{}, errors.New("wallet does not exists")
	}

	response := ws.walletRepository.GetBalance(username)
	balance := response.Balance
	newBalance := balance + amount
	if newBalance < config.GetConfig().MinimumAmount {
		return response, errors.New("not enough amount")
	} else {
		ws.walletRepository.ChangeBalance(username, amount)
		return response, nil
	}
}

func NewWalletService(repository repository.IWalletRepository) IWalletService {
	return &WalletService{
		walletRepository: repository,
	}
}
