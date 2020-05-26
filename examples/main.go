package main

import (
	"os"

	"github.com/nodestory/order_match_engine/matching"
)

// ask - sell
// bid - buy
func main() {
	var orders []*matching.LimitOrder
	switch os.Args[1] {
	case "1":
		orders = []*matching.LimitOrder{
			{Id: "1", Type: "ask", Volume: 100, Price: 2030},
			{Id: "2", Type: "ask", Volume: 100, Price: 2025},
			{Id: "3", Type: "ask", Volume: 200, Price: 2030},
			{Id: "4", Type: "bid", Volume: 100, Price: 2015},
			{Id: "5", Type: "bid", Volume: 200, Price: 2020},
			{Id: "6", Type: "bid", Volume: 200, Price: 2015},
			{Id: "7", Type: "bid", Volume: 250, Price: 2035},
		}
	case "2":
		orders = []*matching.LimitOrder{
			{Id: "1", Type: "bid", Volume: 100, Price: 15},
			{Id: "2", Type: "bid", Volume: 200, Price: 17},
			{Id: "3", Type: "bid", Volume: 150, Price: 16},
			{Id: "4", Type: "ask", Volume: 400, Price: 14},
		}
	case "3":
		orders = []*matching.LimitOrder{
			{Id: "1", Type: "ask", Volume: 100, Price: 15},
			{Id: "2", Type: "ask", Volume: 150, Price: 16},
			{Id: "3", Type: "bid", Volume: 200, Price: 15},
			{Id: "4", Type: "ask", Volume: 150, Price: 14},
		}
	}

	engine := matching.NewEngine("TEST-2")
	for _, order := range orders {
		engine.Submit(order)
	}
}
