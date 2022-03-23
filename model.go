package main

type Item struct {
	ID        int64
	Quantity  int
	UnitPrice int64
}

type Cart struct {
	Items         []Item
	TotalQuantity int32
	TotalItem     int
	Discount      float64
	TotalCharge   float64
}

func (c *Cart) AddItem(i *Item) {

}

func (c *Cart) RemoveItem(i *Item) {
}

func (c *Cart) Calculate() {

}

func (c *Cart) UpdateQuantity(itemId int, qty float64) {
}

func (c *Cart) Checkout() {
}

func SetDiscount(m Cart, discount int64) {
}
