package gearMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Mount struct {
    gorm.Model
    MountID uint    `gorm:"primaryKey"`
    GearModPart GearModPart

    // Parts
    Tacticals []Tactical
    Foregrip Foregrip
    Scope Scope
    Mount Mount
    
}
func (mount *Mount) TableName() string {
    return "public.Mounts"
}


