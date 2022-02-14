package service_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"wallet-service/config"
	"wallet-service/mock"
	"wallet-service/model"
	"wallet-service/service"
)

func TestWalletService_GetAll(t *testing.T) {
	response := map[string]int{
		"Furkan":   4500,
		"Muhammed": 20000,
		"Ahmet":    0,
	}

	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	mockRepository.EXPECT().
		GetAll().
		Return(&response).
		Times(1)

	srv := service.NewWalletService(mockRepository)
	wallets, err := srv.GetAll()

	assert.Equal(t, &response, wallets)
	assert.Nil(t, err)
}

func TestWalletService_Get_ExistsWallet(t *testing.T) {
	username := "Furkan"
	response := model.ResponseModel{
		Username: username, Balance: 4500,
	}

	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	mockRepository.EXPECT().
		GetBalance(username).
		Return(response).
		Times(1)
	mockRepository.EXPECT().
		Exists(username).
		Return(true).
		Times(1)

	srv := service.NewWalletService(mockRepository)
	wallet, err := srv.Get(username)

	assert.Equal(t, response, wallet)
	assert.Nil(t, err)
}

func TestWalletService_Get_NotExistsWallet(t *testing.T) {
	username := "Kübra"
	err := errors.New("wallet does not exists")

	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	mockRepository.EXPECT().
		Exists(username).
		Return(false).
		Times(1)

	srv := service.NewWalletService(mockRepository)
	_, respErr := srv.Get(username)

	assert.Equal(t, err, respErr)
}

func TestWalletService_CreateWallet_WithValidUsername(t *testing.T) {
	username := "Kübra"

	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	mockRepository.EXPECT().
		Exists(username).
		Return(false).
		Times(1)
	mockRepository.EXPECT().
		CreateWallet(username, config.GetConfig().InitialAmount).
		Times(1)

	srv := service.NewWalletService(mockRepository)
	err := srv.CreateWallet(username)

	assert.Nil(t, err)
}

func TestWalletService_CreateWallet_WithNotValidUsername(t *testing.T) {
	username := "Furkan"

	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	mockRepository.EXPECT().
		Exists(username).
		Return(true).
		Times(1)

	srv := service.NewWalletService(mockRepository)
	err := srv.CreateWallet(username)

	assert.Equal(t, errors.New("wallet already exists"), err)
}

func TestNewWalletService_ChangeBalance_NotValidOperation(t *testing.T) {
	responseModel := model.ResponseModel{Username: "Ahmet", Balance: 0}

	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	mockRepository.EXPECT().
		Exists(responseModel.Username).
		Return(true).
		Times(1)
	mockRepository.EXPECT().
		GetBalance(responseModel.Username).
		Return(responseModel).
		Times(1)

	srv := service.NewWalletService(mockRepository)
	_, err := srv.ChangeBalance(responseModel.Username, -2000)

	assert.Equal(t, errors.New("not enough amount"), err)
}

func TestNewWalletService_ChangeBalance_ValidOperation(t *testing.T) {
	responseModel := model.ResponseModel{Username: "Ahmet", Balance: 0}

	mockRepository := mock.NewMockIWalletRepository(gomock.NewController(t))
	mockRepository.EXPECT().
		Exists(responseModel.Username).
		Return(true).
		Times(1)
	mockRepository.EXPECT().
		GetBalance(responseModel.Username).
		Return(responseModel).
		Times(1)
	mockRepository.EXPECT().
		ChangeBalance(responseModel.Username, responseModel.Balance).
		Times(1)

	srv := service.NewWalletService(mockRepository)
	_, err := srv.ChangeBalance(responseModel.Username, 0)

	assert.Nil(t, err)
}
