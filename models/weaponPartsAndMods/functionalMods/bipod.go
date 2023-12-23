package functionMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Bipod struct {
    gorm.Model
    BipodID uint    `gorm:"primaryKey"`
    GearModPart GearModPart

    // Parts
}
func (bipod *Bipod) TableName() string {
    return "public.Bipods"
}


