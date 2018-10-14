package matching

import (
	"github.com/HuKeping/rbtree"
)

type OrderBook interface {
	AddOrder(order Order)
	RemoveOrder(order Order)
	Top() Order
	FillTop(*Trade)
}

type OrderBookManager struct {
	Market      string
	LimitOrders *rbtree.Rbtree
}

func (ob *OrderBookManager) AddOrder(order Order) {
	switch v := order.(type) {
	case *LimitOrder:
		level := ob.LimitOrders.Get(&PriceLevel{Price: v.Price})
		if level == nil {
			pl := NewPriceLevel(v.Price, order)
			ob.LimitOrders.Insert(pl)
			return
		}
		level.(*PriceLevel).Add(order)
	}
}

func (ob *OrderBookManager) RemoveOrder(order Order) {
	switch v := order.(type) {
	case *LimitOrder:
		level := ob.LimitOrders.Get(&PriceLevel{Price: v.Price})
		empty := level.(*PriceLevel).Remove(v.Id)
		if empty {
			ob.LimitOrders.Delete(&PriceLevel{Price: v.Price})
		}
	}
}

type AskOrderBook struct {
	OrderBookManager
}

func NewAskOrderBook(market string) *AskOrderBook {
	return &AskOrderBook{
		OrderBookManager{
			Market:      market,
			LimitOrders: rbtree.New(),
		},
	}
}

func (aob *AskOrderBook) Top() Order {
	pl := aob.LimitOrders.Max()
	if pl == nil {
		return nil
	}
	return pl.(*PriceLevel).Top()
}

func (aob *AskOrderBook) FillTop(trade *Trade) {
	pl := aob.LimitOrders.Max()
	order := pl.(*PriceLevel).Top()

	lm := order.(*LimitOrder)
	lm.Fill(trade.Volume)
	if lm.Filled() {
		aob.RemoveOrder(lm)
	} else {
		// broadcast
	}
}

type BidOrderBook struct {
	OrderBookManager
}

func NewBidOrderBook(market string) *BidOrderBook {
	return &BidOrderBook{
		OrderBookManager{
			Market:      market,
			LimitOrders: rbtree.New(),
		},
	}
}

func (bob *BidOrderBook) Top() Order {
	pl := bob.LimitOrders.Min()
	if pl == nil {
		return nil
	}
	return pl.(*PriceLevel).Top()
}

func (bob *BidOrderBook) FillTop(trade *Trade) {
	pl := bob.LimitOrders.Min()
	order := pl.(*PriceLevel).Top()

	lm := order.(*LimitOrder)
	lm.Fill(trade.Volume)
	if lm.Filled() {
		bob.RemoveOrder(lm)
	} else {
		// broadcast
	}
}
