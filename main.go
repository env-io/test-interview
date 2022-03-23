package main

import "fmt"

func main() {
	m := new(Cart)

	bakso := &Item{
		ID:        1,
		Quantity:  2,
		UnitPrice: 1000,
	}

	m.AddItem(bakso)

	bakso2 := &Item{
		ID:        1,
		Quantity:  3,
		UnitPrice: 1000,
	}

	m.AddItem(bakso2)

	m.UpdateQuantity(1, 3)

	bihun := &Item{
		ID:        2,
		Quantity:  10,
		UnitPrice: 6000,
	}

	m.AddItem(bihun)

	saus := &Item{
		ID:        3,
		Quantity:  3,
		UnitPrice: 2000,
	}

	m.AddItem(saus)

	minyak := &Item{
		ID:        4,
		Quantity:  10,
		UnitPrice: 60000,
	}

	m.AddItem(minyak)

	m.RemoveItem(minyak)

	m.UpdateQuantity(3, 2)

	// keluarkan println ketika item tersebut tidak tersedia
	m.RemoveItem(minyak)

	// discount percentage
	SetDiscount(m, 50)

	m.Calculate()

	m.Checkout()

	fmt.Println("======Hasil Total Charges=======", m.TotalCharge == 33500)
	fmt.Println("======Hasil Total item=======", m.TotalItem == 3)
	fmt.Println("======Hasil Total Quantity=======", m.TotalQuantity == 15)

}
