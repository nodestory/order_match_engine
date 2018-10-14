package matching

import "github.com/HuKeping/rbtree"

type PriceLevel struct {
	Price  int64
	Orders []Order
}

// We will order the node by `Time`
func (pl *PriceLevel) Less(than rbtree.Item) bool {
	return pl.Price < than.(*PriceLevel).Price
}

func NewPriceLevel(price int64, order Order) *PriceLevel {
	return &PriceLevel{
		Price:  price,
		Orders: []Order{order},
	}
}

func (pl *PriceLevel) Top() Order {
	if len(pl.Orders) == 0 {
		return nil
	}
	return pl.Orders[0]
}

func (pl *PriceLevel) Empty() bool {
	return len(pl.Orders) == 0
}

func (pl *PriceLevel) Add(order Order) {
	pl.Orders = append(pl.Orders, order)
}

func (pl *PriceLevel) Remove(oid string) bool {
	for i, o := range pl.Orders {
		if o.(*LimitOrder).Id == oid {
			pl.Orders = append(pl.Orders[:i], pl.Orders[i+1:]...)
			return len(pl.Orders) == 0
		}
	}
	return false
}

func (pl *PriceLevel) Find(id string) Order {
	for _, o := range pl.Orders {
		if o.(*LimitOrder).Id == id {
			return o
		}
	}

	return nil
}
