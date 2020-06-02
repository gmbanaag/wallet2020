package handler

import (
	"fmt"
	"net/http"

	"github.com/gmbanaag/wallet2020/internal/app/logger"
	"github.com/gmbanaag/wallet2020/internal/app/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

//GetSentTransactions handler
func (h *Handler) GetSentTransactions(w http.ResponseWriter, r *http.Request) {
	reqID = r.Context().Value(middleware.RequestIDKey)
	user := r.Context().Value(h.cfg.UserCtxKey)
	retrieveUser(user)

	transactions := []model.Transaction{}
	transaction := model.Transaction{}
	transactions, err := transaction.GetTransactionBySourceUserID(userToken.UserID)

	if err != nil {
		logger.LogError(fmt.Sprintf("%s | %s", err.Error(), reqID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.DefaultResponder(w, r, transactions)
	return
}

//GetReceivedTransactions handler
func (h *Handler) GetReceivedTransactions(w http.ResponseWriter, r *http.Request) {
	reqID = r.Context().Value(middleware.RequestIDKey)
	user := r.Context().Value(h.cfg.UserCtxKey)
	retrieveUser(user)

	transactions := []model.Transaction{}
	transaction := model.Transaction{}
	transactions, err := transaction.GetTransactionByDestinationUserID(userToken.UserID)

	if err != nil {
		logger.LogError(fmt.Sprintf("%s [%s]", err.Error(), reqID))
		w.WriteHeader(http.StatusNotFound)
	} else {
		render.Status(r, http.StatusOK)
		render.DefaultResponder(w, r, transactions)
	}
	return
}

//GetAllTransactions handler
func (h *Handler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	reqID = r.Context().Value(middleware.RequestIDKey)
	user := r.Context().Value(h.cfg.UserCtxKey)
	retrieveUser(user)

	if userToken.Scope != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	transactions := []model.Transaction{}
	transaction := model.Transaction{}
	transactions, err := transaction.GetAllTransactions()

	if err != nil {
		logger.LogError(fmt.Sprintf("%s [%s]", err.Error(), reqID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.DefaultResponder(w, r, transactions)
	return

}

//GetTransaction handler
func (h *Handler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	reqID = r.Context().Value(middleware.RequestIDKey)
	transactionID := chi.URLParam(r, "transactionID")

	transaction := model.Transaction{}
	err := transaction.GetTransactionByID(transactionID)

	if err != nil {
		logger.LogError(fmt.Sprintf("%s [%s]", err.Error(), reqID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	user := r.Context().Value(h.cfg.UserCtxKey)
	retrieveUser(user)

	if transaction.SourceUserID != userToken.UserID {
		//user cannot retrieve other user's transactions
		logger.LogDebug(fmt.Sprintf("user %s cannot retrieve transaction %s [%s]", userToken.UserID, transactionID, reqID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.DefaultResponder(w, r, transaction)
	return
}
