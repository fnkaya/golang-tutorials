package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"wallet-service/model"
	"wallet-service/service"
)

type IWalletHandler interface {
	HandleRequest(w http.ResponseWriter, r *http.Request)
}

type WalletHandler struct {
	walletService service.IWalletService
}

func (wh *WalletHandler) HandleRequest(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		wh.handleGetRequest(w, r)
	case http.MethodPut:
		wh.handlePutRequest(w, r)
	case http.MethodPost:
		wh.handlePostRequest(w, r)
	}
}

func (wh *WalletHandler) handleGetRequest(w http.ResponseWriter, r *http.Request) {
	var (
		response interface{}
		err      error
	)

	param := strings.TrimPrefix(r.URL.Path, "/")
	if len(param) == 0 {
		response, err = wh.walletService.GetAll()
	} else {
		response, err = wh.walletService.Get(param)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(jsonResponse)
}

func (wh *WalletHandler) handlePutRequest(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/")
	err := wh.walletService.CreateWallet(param)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (wh *WalletHandler) handlePostRequest(w http.ResponseWriter, r *http.Request) {
	param := strings.TrimPrefix(r.URL.Path, "/")
	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	requestModel := &model.RequestModel{}
	err = json.Unmarshal(b, requestModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	response, err := wh.walletService.ChangeBalance(param, requestModel.Balance)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Write(jsonResponse)
}

func NewWalletHandler(service service.IWalletService) IWalletHandler {
	return &WalletHandler{
		walletService: service,
	}
}
