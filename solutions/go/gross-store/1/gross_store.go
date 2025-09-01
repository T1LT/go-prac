package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unitsMap := map[string]int{}
    unitsMap["quarter_of_a_dozen"] = 3
    unitsMap["half_of_a_dozen"] = 6
    unitsMap["dozen"] = 12
    unitsMap["small_gross"] = 120
    unitsMap["gross"] = 144
    unitsMap["great_gross"] = 1728

    return unitsMap
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	u, prs := units[unit]
    if !prs {
        return false
    } else {
        bill[item] += u
        return true
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, itemPresent := bill[item]
    _, unitPresent := units[unit]
    
    if !itemPresent || !unitPresent {
        return false
    } else {
     	newVal := bill[item] - units[unit]
		if newVal < 0 {
            return false
        } else if newVal == 0 {
            delete(bill, item)
            return true
        } else {
            bill[item] = newVal
            return true
        }
    }
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, prs := bill[item]
    if !prs {
        return 0, false
    } else {
        return qty, true
    }
}
