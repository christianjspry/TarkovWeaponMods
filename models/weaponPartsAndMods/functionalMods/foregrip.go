package functionMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Foregrip struct {
    gorm.Model
    ForegripID uint    `gorm:"primaryKey"`
    GearModPart GearModPart

    // Parts
}
func (forgrip *Foregrip) TableName() string {
    return "public.Foregrips"
}


