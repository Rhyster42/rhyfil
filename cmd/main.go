package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"rhyfil/internal/config"
	"rhyfil/internal/database"

	"log"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read() //builds from config file
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DBURL) //grab database
	if err != nil {
		log.Fatalf("error opening database: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := commands{ // set up CLI commands
		registeredCommands: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("newproduct", handlerAddProduct)
	cmds.register("clearusers", handlerClearUsers)
	cmds.register("clearproducts", handlerClearProducts)
	cmds.register("serve", handlerSpinServer)
	cmds.register("addmodifiergroup", handlerAddModGroup)
	cmds.register("addgroupoption", handlerAddModOption)

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = cmds.run(programState, command{Name: cmdName, Args: cmdArgs})
	if err != nil {
		log.Fatal(err)
	}

}
