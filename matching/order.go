package matching

type Order interface {
	Fill(tradeVolume int64)
	Filled() bool
	TradeWith(counterOrder Order, counterBook OrderBook) *Trade
}
