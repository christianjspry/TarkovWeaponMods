package vitalParts

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Barrel struct {
    gorm.Model
    BarrelID uint    `gorm:"primaryKey"`
    GearModPart GearModPart
    Caliber Caliber
    MuzzleVelocity float32
    Accuracy float32

    // Parts
    // Muzzle Muzzle
    GasBlock GasBlock
}
func (barrel *Barrel) TableName() string {
    return "public.Barrels"
}


