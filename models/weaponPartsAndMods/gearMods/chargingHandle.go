package gearMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type ChargingHandle struct {
    gorm.Model
    ChargingHandleID uint    `gorm:"primaryKey"`
    GearModPart GearModPart

    // Parts
}
func (chargingHandle *ChargingHandle) TableName() string {
    return "public.ChargingHandles"
}


