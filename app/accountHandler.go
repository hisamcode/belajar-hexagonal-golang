package app

import (
  "net/http"
  "encoding/json"

  "github.com/hisamcode/belajar-hexagonal-golang/service"
  "github.com/hisamcode/belajar-hexagonal-golang/dto"
)

type AccountHandler struct {
  service service.AccountService
}

func (h AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {
  var request dto.NewAccountRequest
  err := json.NewDecoder(r.Body).Decode(&request)
  if err != nil {
    writeResponse(w, http.StatusBadRequest, err.Error())
  } else {
    account, appError := h.service.NewAccount(request)
    if appError != nil {
      writeResponse(w, appError.Code, appError.Message)
    } else {
      writeResponse(w, http.StatusCreated, account)
    }
  }

}