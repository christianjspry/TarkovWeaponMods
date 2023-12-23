package gearMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Stock struct {
    gorm.Model
    StockID uint    `gorm:"primaryKey"`
    GearModPart GearModPart

    // Parts
    Stock Stock
    PistolGrip PistolGrip
    Tactical Tactical
}
func (stock *Stock) TableName() string {
    return "public.Stocks"
}


