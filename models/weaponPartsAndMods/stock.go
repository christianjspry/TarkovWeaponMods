package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type Stock struct {
    gorm.Model
    StockID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
    Stock *Stock
    PistolGrip PistolGrip
    Tactical Tactical
}
func (stock *Stock) TableName() string {
    return "public.Stocks"
}


