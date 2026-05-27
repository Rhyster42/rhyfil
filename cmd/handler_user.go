package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"rhyfil/internal/database"
	"rhyfil/server"
	"time"

	"github.com/google/uuid"
)

type login struct {
	Name string `json:"name"`
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}
	name := cmd.Args[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("error setting user: %v", err)
	}

	fmt.Println("User Created!")
	return nil
}

func printUser(user database.User) {
	fmt.Printf(" * ID: %s\n", user.ID)
	fmt.Printf(" * Name: %s\n", user.Name)
}

func handlerClearUsers(s *state, cmd command) error {
	err := s.db.ClearUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error clearing users: %v", err)
	}
	fmt.Println("All users cleared!")
	return nil
}

func handlerHTTPLogin(s *state) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var login login
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&login); err != nil {
			server.RespondWithError(w, http.StatusBadRequest, "failed to decode JSON: ", err)
			return
		}

		loginData, err := s.db.GetUser(context.Background(), login.Name)
		if err != nil {
			server.RespondWithError(w, http.StatusUnauthorized, "failed to retrieve login data: ", err)
			return
		}

		server.RespondWithJSON(w, http.StatusOK, loginData)
	}
}
