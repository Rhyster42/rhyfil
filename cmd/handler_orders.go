package main

import (
	"context"
	"encoding/json"
	"net/http"
	"rhyfil/internal/database"
	"rhyfil/server"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type order struct {
	Total float64 `json:"total"`
	Items []item  `json:"items"`
}

type item struct {
	ProductId uuid.UUID `json:"productId"`
	ItemTotal float64   `json:"itemTotal"`
	ItemMods  []itemMod `json:"itemMods"`
	Quantity  int       `json:"quantity"`
}

type itemMod struct {
	ModId           uuid.UUID `json:"modId"`
	PriceAdjustment float64   `json:"priceAdjustment"`
}

func handlerCheckOut(s *state) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var order order //contains json values
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&order); err != nil {
			server.RespondWithError(w, http.StatusInternalServerError, "failed to decode JSON: ", err)
			return
		}

		//orderInfo - returned information from database used for order_items and order_modifiers tables
		orderInfo, err := s.db.CreateOrder(context.Background(), database.CreateOrderParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			Total:     strconv.FormatFloat(order.Total, 'f', -1, 64),
		})
		if err != nil {
			server.RespondWithError(w, http.StatusInternalServerError, "failed to add order to database: ", err)
			return
		}

		for _, item := range order.Items {
			itemInfo, err := s.db.AddItemToOrder(context.Background(), database.AddItemToOrderParams{
				ID:              uuid.New(),
				OrderID:         orderInfo.ID,
				ProductID:       item.ProductId,
				Quantity:        int32(item.Quantity),
				PriceAtPurchase: strconv.FormatFloat(item.ItemTotal, 'f', -1, 64),
			})
			if err != nil {
				server.RespondWithError(w, http.StatusInternalServerError, "failed to add item to order in database: ", err)
				return
			}

			for _, modifier := range item.ItemMods {
				err := s.db.AddModifierToOrderItem(context.Background(), database.AddModifierToOrderItemParams{
					OrderItemID:               itemInfo.ID,
					ModifierOptionID:          modifier.ModId,
					PriceAdjustmentAtPurchase: strconv.FormatFloat(modifier.PriceAdjustment, 'f', -1, 64),
				})

				if err != nil {
					server.RespondWithError(w, http.StatusInternalServerError, "failed to add mod to order item in database: ", err)
					return
				}
			}
		}
		server.RespondWithJSON(w, http.StatusOK, orderInfo)
	}
}
