package main

import (
	"context"
	"fmt"
	"rhyfil/internal/database"

	"github.com/google/uuid"
)

func handlerAddModGroup(s *state, cmd command) error {
	if len(cmd.Args) < 1 {
		return fmt.Errorf("usage: %s <Modifier Group Name> for example: Size, Milk...", cmd.Name)
	}

	groupName := cmd.Args[0]

	_, err := s.db.CreateModifierGroup(context.Background(),
		database.CreateModifierGroupParams{
			ID:   uuid.New(),
			Name: groupName,
		},
	)
	if err != nil {
		return fmt.Errorf("error add product: %v", err)
	}

	fmt.Printf("Modifier group '%s' added successfully \n", groupName)

	return nil
}

func handlerAddModOption(s *state, cmd command) error {
	if len(cmd.Args) < 3 {
		return fmt.Errorf("usage: %s <Modifier Option Name> <Price Adjustment> <Modifier Group ID OR Modifier Group Name>", cmd.Name)
	}
	isUUID, parsedID := helperResolveUUID(cmd.Args[2])
	if !isUUID {
		retrievedID, err := s.db.GetIdForGroup(context.Background(), cmd.Args[2])
		if err != nil {
			return fmt.Errorf("Failed to retrieve ID from database: %s", err)
		}
		parsedID = retrievedID
	}
	price := cmd.Args[1]
	optionName := cmd.Args[0]

	_, err := s.db.CreateModifierOption(context.Background(), database.CreateModifierOptionParams{
		ID:              uuid.New(),
		Name:            optionName,
		PriceAdjustment: price,
		ModifierGroupID: parsedID,
	})
	if err != nil {
		return fmt.Errorf("Error adding Modifier Option to Modifier Group: %v", err)
	}
	fmt.Printf("Modifier option %s added to modifier group %s \n", optionName, cmd.Args[2])
	return nil
}

func handlerLinkModifier(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("usage: %s <Product Name or ID> <Modifier Group Name or ID>", cmd.Name)
	}

	isUUID, parsedProductID := helperResolveUUID(cmd.Args[0])
	if !isUUID {
		retrievedID, err := s.db.GetIdForProduct(context.Background(), cmd.Args[0])
		if err != nil {
			return fmt.Errorf("Failed to retreive Product ID from database: %v", err)
		}
		parsedProductID = retrievedID
	}

	isUUID, parsedModGroupID := helperResolveUUID(cmd.Args[1])
	if !isUUID {
		retrievedID, err := s.db.GetIdForGroup(context.Background(), cmd.Args[1])
		if err != nil {
			return fmt.Errorf("Failed to retrieve Modifier Group ID from database: %v", err)
		}
		parsedModGroupID = retrievedID
	}

	err := s.db.LinkModifierGroupToProduct(context.Background(), database.LinkModifierGroupToProductParams{
		ProductID:       parsedProductID,
		ModifierGroupID: parsedModGroupID,
	})
	if err != nil {
		return fmt.Errorf("Failed to link modifier group to product in database: %v", err)
	}

	fmt.Printf("Product %s linked to Modifer Group %s \n", cmd.Args[0], cmd.Args[1])
	return nil
}

func helperResolveUUID(value string) (bool, uuid.UUID) { //helper to check if value is a valid UUID or a string
	parsedID, err := uuid.Parse(value)
	if err != nil {
		return false, uuid.Nil
	}
	return true, parsedID
}
