package functionMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type RearSight struct {
    gorm.Model
    RearSightID uint    `gorm:"primaryKey"`
    GearModPart GearModPart
    SightingRange int16

    // Parts
    Mount Mount
    Scope Scope
}
func (rearSight *RearSight) TableName() string {
    return "public.RearSights"
}


