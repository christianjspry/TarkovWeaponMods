package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type ChargingHandle struct {
    gorm.Model
    ChargingHandleID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
}
func (chargingHandle *ChargingHandle) TableName() string {
    return "public.ChargingHandles"
}


