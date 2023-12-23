package vitalParts

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models"
    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type PistolGrip struct {
    gorm.Model
    PistolGripID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
    Stock Stock
}
func (pistolGrip *PistolGrip) TableName() string {
    return "public.PistolGrip"
}
