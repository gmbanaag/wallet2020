package handler

import (
	"github.com/gmbanaag/wallet2020/internal/app/logger"

	"github.com/go-chi/chi"
)

//LoadRouters warms-up all the available route for the service
func LoadRouters(router *chi.Mux, hd *Handler) *chi.Mux {
	router.Route("/v1/admin", func(r chi.Router) {
		r.Get("/wallets", hd.GetAllWallets)
		r.Get("/transactions", hd.GetAllTransactions)
	})

	router.Route("/v1", func(r chi.Router) {
		r.Route("/wallets", func(r chi.Router) {
			r.Get("/", hd.GetWallets)
			r.Get("/{walletID}", hd.GetWallet)
		})

		r.Post("/transfer", hd.Transfer)

		r.Route("/transactions", func(r chi.Router) {
			r.Get("/sent", hd.GetSentTransactions)
			r.Get("/received", hd.GetReceivedTransactions)
			r.Get("/{transactionID}", hd.GetTransaction)
		})
	})

	logger.LogInfo("Loading routes are done")

	return router
}
