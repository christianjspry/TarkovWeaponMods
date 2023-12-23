package functionMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type FrontSight struct {
    gorm.Model
    FrontSightID uint    `gorm:"primaryKey"`
    GearModPart GearModPart
    SightingRange int16

    // Parts
    Mount Mount
    Scope Scope
}
func (frontSight *FrontSight) TableName() string {
    return "public.FrontSights"
}


