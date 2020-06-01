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

//GetWallets handler
func (h *Handler) GetWallets(w http.ResponseWriter, r *http.Request) {
	reqID = r.Context().Value(middleware.RequestIDKey)
	user := r.Context().Value(h.cfg.UserCtxKey)
	retrieveUser(user)

	wallets := []model.Wallet{}
	wallet := model.Wallet{}
	wallets, err := wallet.GetWalletsByUserID(userToken.UserID)

	if err != nil {
		logger.LogError(fmt.Sprintf("%s [%s]", err.Error(), reqID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.DefaultResponder(w, r, wallets)
	return
}

//GetAllWallets handler
func (h *Handler) GetAllWallets(w http.ResponseWriter, r *http.Request) {
	reqID = r.Context().Value(middleware.RequestIDKey)
	user := r.Context().Value(h.cfg.UserCtxKey)
	retrieveUser(user)

	if userToken.Scope != "admin" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	wallets := []model.Wallet{}
	wallet := model.Wallet{}
	wallets, err := wallet.GetAllWallets()

	if err != nil {
		logger.LogError(fmt.Sprintf("%s [%s]", err.Error(), reqID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.DefaultResponder(w, r, wallets)
	return

}

//GetWallet get wallet information
func (h *Handler) GetWallet(w http.ResponseWriter, r *http.Request) {
	walletID := chi.URLParam(r, "walletID")

	wallet := model.Wallet{}
	err := wallet.GetWalletByID(walletID)

	if err != nil {
		logger.LogError(fmt.Sprintf("%s [%s]", err.Error(), reqID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	reqID = r.Context().Value(middleware.RequestIDKey)
	user := r.Context().Value(h.cfg.UserCtxKey)
	retrieveUser(user)

	if wallet.UserID != userToken.UserID {
		//user cannot retrieve other user's transactions
		logger.LogDebug(fmt.Sprintf("user %s cannot retrieve wallet %s [%s]", userToken.UserID, wallet.ID, reqID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	render.Status(r, http.StatusOK)
	render.DefaultResponder(w, r, wallet)
	return
}
