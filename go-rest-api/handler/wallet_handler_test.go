package handler_test

import (
	"encoding/json"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"wallet-service/handler"
	"wallet-service/mock"
	"wallet-service/model"
)

func TestWalletHandler_GetAll(t *testing.T) {
	srv := mock.NewMockIWalletService(gomock.NewController(t))
	srvResponse := map[string]int{
		"Furkan":   4500,
		"Ahmet":    0,
		"Muhammed": 2000,
	}
	srv.EXPECT().
		GetAll().
		Return(&srvResponse, nil).
		Times(1)

	hndlr := handler.NewWalletHandler(srv)
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	hndlr.HandleRequest(w, r)

	var actual map[string]int
	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, srvResponse, actual)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestWalletHandler_Get(t *testing.T) {
	srv := mock.NewMockIWalletService(gomock.NewController(t))
	username := "Furkan"
	srvResponse := model.ResponseModel{
		Username: username,
		Balance:  4500,
	}
	srv.EXPECT().
		Get(username).
		Return(srvResponse, nil).
		Times(1)

	hndlr := handler.NewWalletHandler(srv)
	r := httptest.NewRequest(http.MethodGet, "/"+username, nil)
	w := httptest.NewRecorder()
	hndlr.HandleRequest(w, r)

	actual := model.ResponseModel{}
	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, srvResponse, actual)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}

func TestWalletHandler_Get_NotExistsWallet(t *testing.T) {
	srv := mock.NewMockIWalletService(gomock.NewController(t))
	username := "Mahmut"
	emptyBody := model.ResponseModel{}
	errResponse := errors.New("wallet does not exists")
	srv.EXPECT().
		Get(username).
		Return(emptyBody, errResponse).
		Times(1)

	hndlr := handler.NewWalletHandler(srv)
	r := httptest.NewRequest(http.MethodGet, "/"+username, nil)
	w := httptest.NewRecorder()
	hndlr.HandleRequest(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}

func TestNewWalletHandler_Put(t *testing.T) {
	srv := mock.NewMockIWalletService(gomock.NewController(t))
	username := "Mahmut"
	srv.EXPECT().
		CreateWallet(username).
		Return(nil).
		Times(1)

	hndlr := handler.NewWalletHandler(srv)
	r := httptest.NewRequest(http.MethodPut, "/"+username, nil)
	w := httptest.NewRecorder()
	hndlr.HandleRequest(w, r)

	assert.Equal(t, http.StatusCreated, w.Result().StatusCode)
}

func TestNewWalletHandler_Put_Error(t *testing.T) {
	srv := mock.NewMockIWalletService(gomock.NewController(t))
	username := "Furkan"
	errResponse := errors.New("wallet already exists")
	srv.EXPECT().
		CreateWallet(username).
		Return(errResponse).
		Times(1)

	hndlr := handler.NewWalletHandler(srv)
	r := httptest.NewRequest(http.MethodPut, "/"+username, nil)
	w := httptest.NewRecorder()
	hndlr.HandleRequest(w, r)

	assert.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}

func TestNewWalletHandler_Post(t *testing.T) {
	srv := mock.NewMockIWalletService(gomock.NewController(t))
	username := "Ahmet"
	requestBody := strings.NewReader(`{ "balance": 300 }`)
	responseModel := model.ResponseModel{
		Username: username,
		Balance:  0,
	}
	srv.EXPECT().
		ChangeBalance(username, 300).
		Return(responseModel, nil).
		Times(1)

	hndlr := handler.NewWalletHandler(srv)
	r := httptest.NewRequest(http.MethodPost, "/"+username, requestBody)
	r.Header.Set("content-type", "application/json")
	w := httptest.NewRecorder()
	hndlr.HandleRequest(w, r)

	var actual model.ResponseModel
	json.Unmarshal(w.Body.Bytes(), &actual)

	assert.Equal(t, responseModel, actual)
	assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	assert.Equal(t, "application/json", w.Header().Get("content-type"))
}
