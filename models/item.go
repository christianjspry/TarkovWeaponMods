package models

//import ()

// Fields each item will have
type Item struct {
    Name string
    Weight float32
    Description string

    GridHeight int16
    GridWidth int16

    SoldPrice int
    SoldBy string // Trader?
    SoldLoyaltyLevel int
    SoldCurrency string // Currency enum?

    BoughtPrice int
    BoughtBy string // Trader
    BoughtCurrency string // always rubles?
}
