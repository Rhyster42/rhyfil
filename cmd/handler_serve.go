package main

import (
	"context"
	"log"
	"net/http"
	"rhyfil/server"

	"github.com/google/uuid"
)

func handlerSpinServer(s *state, cmd command) error {

	// handles retrieving menu items from db //
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {

		menuItems, err := s.db.GetAllProducts(context.Background())
		if err != nil {
			server.RespondWithError(w, http.StatusInternalServerError, "failed to retrieve menu from database:", err)
			return
		}
		server.RespondWithJSON(w, http.StatusOK, menuItems)
	})

	// handles retrieving modifier groups //
	http.HandleFunc("/modifiers", func(w http.ResponseWriter, r *http.Request) {
		productIDString := r.URL.Query().Get("product_id")
		productID, err := uuid.Parse(productIDString)
		if err != nil {
			server.RespondWithError(w, http.StatusBadRequest, "failed to parse uuid for product modifier groups:", err)
			return
		}

		modifiers, err := s.db.GetAllModifierGroupsForProduct(context.Background(), productID)
		if err != nil {
			server.RespondWithError(w, http.StatusInternalServerError, "failed to retrieve product modifier groups from database:", err)
			return
		}

		server.RespondWithJSON(w, http.StatusOK, modifiers)
	})

	// handler retrieving modifier options //
	http.HandleFunc("/modifier_options", func(w http.ResponseWriter, r *http.Request) {
		modifierGroupString := r.URL.Query().Get("group_id")
		modifierGroupID, err := uuid.Parse(modifierGroupString)
		if err != nil {
			server.RespondWithError(w, http.StatusBadRequest, "failed to parse uuid for modifier group:", err)
			return
		}

		modifierOptions, err := s.db.GetAllOptionsForModifierGroup(context.Background(), modifierGroupID)
		if err != nil {
			server.RespondWithError(w, http.StatusInternalServerError, "failed to retrieve modifier options from database:", err)
			return
		}

		server.RespondWithJSON(w, http.StatusOK, modifierOptions)
	})

	http.HandleFunc("/checkout", handlerCheckOut(s))

	http.HandleFunc("/login", handlerHTTPLogin(s))

	http.Handle("/", http.FileServer(http.Dir("../frontend")))

	log.Print("Running server...")

	return http.ListenAndServe(":8080", nil)
}
