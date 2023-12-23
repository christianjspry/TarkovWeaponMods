package gearMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Launcher struct {
    gorm.Model
    LauncherID uint    `gorm:"primaryKey"`
    GearModPart GearModPart
    Caliber Caliber

    // Parts
}
func (launcher *Launcher) TableName() string {
    return "public.Launchers"
}


