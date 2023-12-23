package functionMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Muzzle struct {
    gorm.Model
    MuzzleID uint    `gorm:"primaryKey"`
    GearModPart GearModPart
    Loudness int16

    // Parts
    Muzzle Muzzle
}
func (muzzle *Muzzle) TableName() string {
    return "public.Muzzles"
}


