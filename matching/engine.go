package matching

import "fmt"

type Engine struct {
	market       string
	askOrderBook OrderBook
	bidOrderBook OrderBook
}

func NewEngine(market string) *Engine {
	return &Engine{
		askOrderBook: NewAskOrderBook(market),
		bidOrderBook: NewBidOrderBook(market),
	}
}

func (e *Engine) Submit(order *LimitOrder) {
	switch order.Type {
	case "ask":
		e.match(order, e.bidOrderBook)
		e.add(order, e.askOrderBook)
	case "bid":
		e.match(order, e.askOrderBook)
		e.add(order, e.bidOrderBook)
	}
}

func (e *Engine) match(order Order, counterBook OrderBook) {
	if order.Filled() {
		return
	}

	counterOrder := counterBook.Top()
	if counterOrder == nil {
		return
	}

	trade := order.TradeWith(counterOrder, counterBook)
	if trade != nil {
		counterBook.FillTop(trade)
		order.Fill(trade.Volume)

		// TODO: publish order, counter_order, trade
		fmt.Println("matched", order, counterOrder, trade, order.Filled())

		e.match(order, counterBook)
	}
}

func (e *Engine) add(order Order, book OrderBook) {
	if !order.Filled() {
		book.AddOrder(order)
	}
}
