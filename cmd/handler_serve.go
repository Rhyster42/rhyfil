package main

import (
	"context"
	"net/http"
	"rhyfil/server"
)

func handlerSpinServer(s *state, cmd command) error {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {

		menuItems, err := s.db.GetAllProducts(context.Background())
		if err != nil {
			server.RespondWithError(w, http.StatusInternalServerError, "failed to retrieve menu from database:", err)
			return
		}
		server.RespondWithJSON(w, http.StatusOK, menuItems)
	})

	return http.ListenAndServe(":8080", nil)
}
