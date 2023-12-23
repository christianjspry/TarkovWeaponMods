package functionMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Tactical struct {
    gorm.Model
    TacticalID uint    `gorm:"primaryKey"`
    GearModPart GearModPart

    // Parts
}
func (tactical *Tactical) TableName() string {
    return "public.Tacticals"
}


