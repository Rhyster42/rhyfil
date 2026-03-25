package main

import (
	"rhyfil/internal/config"
	"rhyfil/internal/database"

	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("Rhys")
	if err != nil {
		log.Fatalf("error setting user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config after setting user: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Printf("Failed to open connection to database: %s", err)
		os.Exit(1)
	}
	dbQueuries := database.New(db)

	newState := state{
		db:  dbQueuries,
		cfg: &cfg,
	}

	cmds := commands{}

	cmds.register("add", HandlerAddProduct)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(&newState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
