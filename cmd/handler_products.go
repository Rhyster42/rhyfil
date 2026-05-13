package main

import (
	"context"
	"database/sql"
	"fmt"
	"rhyfil/internal/database"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func handlerAddProduct(s *state, cmd command) error {
	if len(cmd.Args) < 3 {
		return fmt.Errorf("usage: %s <Product Name> <Quantity> <Price>", cmd.Name)
	}

	productName := cmd.Args[0]
	quantity, err := strconv.ParseInt(cmd.Args[1], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid quantity: %v", err)
	}
	price := cmd.Args[2]

	_, err = s.db.CreateProduct(context.Background(),
		database.CreateProductParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      productName,
			Count: sql.NullInt32{
				Int32: int32(quantity),
				Valid: true,
			},
			Price: sql.NullString{
				String: price,
				Valid:  true,
			},
		})
	if err != nil {
		return fmt.Errorf("error adding product: %v", err)
	}

	fmt.Printf("Product '%s' added successfully with price %s and quantity %v\n", productName, price, quantity)

	return nil
}

func handlerClearProducts(s *state, cmd command) error {
	err := s.db.DeleteAllProducts(context.Background())
	if err != nil {
		return fmt.Errorf("error clearing products: %v", err)
	}
	fmt.Println("All products cleared!")
	return nil
}
