package vitalParts

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models"
    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type GasBlock struct {
    gorm.Model
    GasBlockID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
    Scopes []Scope
    Tacticals []Tactical
    Mounts []Mounts
}
func (gasBlock *GasBlock) TableName() string {
    return "public.GasBlock"
}
