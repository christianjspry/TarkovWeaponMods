package vitalParts

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models"
    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Handguard struct {
    gorm.Model
    HandguardID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
    Tacticals []Tactical
    Mounts []Mounts

    FrontSight FrontSight
    Scopes Scope
    Foregrip Foregrip
}
func (handguard *Handguard) TableName() string {
    return "public.Handguard"
}
