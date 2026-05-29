package main

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*state, command) error
}

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]cliCommand
}

func (c *commands) register(name string, cmd cliCommand) {
	c.registeredCommands[name] = cmd
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.name]
	if !ok {
		return errors.New("command not found")
	}
	return f.callback(s, cmd)
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays all registered commands",
			callback:    handlerHelp,
		},
		"register": {
			name:        "register",
			description: "Registers a user to the database, requires <name> argument",
			callback:    handlerRegister,
		},
		"login": {
			name:        "login",
			description: "Logs in user, requires <name> argument",
			callback:    handlerLogin,
		},
		"new-product": {
			name:        "new-product",
			description: "Adds product to database, requires <Name> <Quantity> <Price> arguments",
			callback:    handlerAddProduct,
		},
		"clear-users": {
			name:        "clear-users",
			description: "removes all users from database",
			callback:    handlerClearUsers,
		},
		"clear-products": {
			name:        "clear-products",
			description: "removes all products from database",
			callback:    handlerClearProducts,
		},
		"serve": {
			name:        "serve",
			description: "initializes server to listen for http requests from the DOM",
			callback:    handlerSpinServer,
		},
		"add-modifier-group": {
			name:        "add-modifier-group",
			description: "Adds a modifier group to database: requires <name> argument",
			callback:    handlerAddModGroup,
		},
		"add-group-option": {
			name:        "add-group-option",
			description: "Adds an option to a preexisting modifier group. requires <name> <price of option> <modifier group name or id> arguments",
			callback:    handlerAddModOption,
		},
		"link-product-to-group": {
			name:        "link-product-to-group",
			description: "connects a product to a modifier group. The group will now display in the modifiers with modifier options when the product is selected: requires <product name or id> <modifer group name or id> arguments.",
			callback:    handlerLinkModifier,
		},
	}
}

func handlerHelp(s *state, cmd command) error {
	fmt.Println("Welcome to rhyfil!")
	fmt.Printf("Available commands:\n\n")

	cmds := getCommands()

	keys := make([]string, 0, len(getCommands()))
	for _, cmd := range cmds {
		keys = append(keys, cmd.name)
	}
	slices.SortFunc(keys, func(a, b string) int {
		return cmp.Compare(len(a), len(b))
	})

	for _, cmdName := range keys {
		fmt.Printf("%s: %s\n\n", cmdName, cmds[cmdName].description)
	}
	return nil
}
