package gearMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Magazine struct {
    gorm.Model
    MagazineID uint    `gorm:"primaryKey"`
    GearModPart GearModPart
    MaxCount int
    CheckSpeedModifier float32
    FailureToFeedChance string

    // Parts
}
func (magazine *Magazine) TableName() string {
    return "public.Magazines"
}


