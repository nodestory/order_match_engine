package matching

import (
	"time"
)

type LimitOrder struct {
	Id        string
	Type      string
	Timestamp time.Time
	Volume    int64
	Price     int64
	Market    string
}

func (lm *LimitOrder) Fill(tradeVolume int64) {
	if tradeVolume <= lm.Volume {
		lm.Volume -= tradeVolume
	}
}

func (lm *LimitOrder) Filled() bool {
	return lm.Volume == 0
}

func (lm *LimitOrder) TradeWith(counterOrder Order, counterBook OrderBook) *Trade {
	if lm.crossed(counterOrder.(*LimitOrder).Price) {
		price := counterOrder.(*LimitOrder).Price
		volume := min(lm.Volume, counterOrder.(*LimitOrder).Volume)
		funds := price * volume
		return &Trade{price, volume, funds}
	}
	return nil
}

func (lm *LimitOrder) crossed(price int64) bool {
	switch lm.Type {
	case "ask": // buy
		return price <= lm.Price // if people offer price higher or equal than ask limit
	case "bid": // sell
		return price >= lm.Price // if people offer price lower or equal than bid limit
	default:
		return false
	}
}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
