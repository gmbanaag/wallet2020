package handler

import (
	"net/http"

	"github.com/gmbanaag/wallet2020/internal/app/logger"
	"github.com/gmbanaag/wallet2020/internal/app/metrics"

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

	m := &metrics.Metrics{}
	router.Get("/metrics", wrapHandler(m.PrometheusHandler()))

	logger.LogInfo("Loading routes are done")

	return router
}

func wrapHandler(h http.Handler) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(rw, req)
	}
}
