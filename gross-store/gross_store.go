package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	// panic("Please implement the Units() function")

	return map[string]int {
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	// panic("Please implement the NewBill() function")

	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the AddItem() function")

	value, unitExists := units[unit]

	if !unitExists {
		return false
	}

	_, billExists := bill[item]

	if billExists {
		bill[item] += value
	} else {
		bill[item] = value
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// panic("Please implement the RemoveItem() function")
	if _, ok := bill[item]; !ok {
		return false
	}

	if _, ok := units[unit]; !ok {
		return false
	}

	newUnit := bill[item] - units[unit]
	switch {
		case newUnit < 0:
			return false
		case newUnit == 0:
			delete(bill, item)
		default:
			bill[item] -= units[unit]
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	// panic("Please implement the GetItem() function")
	if _, ok := bill[item]; !ok {
		return 0, false
	}

	return bill[item], true
}
