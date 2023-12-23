package weaponPartsAndMods

import (
    "gorm.io/gorm"
    ammo "TarkovWeaponMods/models/ammo"
)

type Barrel struct {
    gorm.Model
    BarrelID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod 
    Caliber ammo.Caliber
    MuzzleVelocity float32
    Accuracy float32

    // Parts
    // Muzzle Muzzle
    GasBlock GasBlock
}
func (barrel *Barrel) TableName() string {
    return "public.Barrels"
}


