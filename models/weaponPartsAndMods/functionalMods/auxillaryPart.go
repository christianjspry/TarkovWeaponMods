package functionMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type AuxillaryPart struct {
    gorm.Model
    AuxillaryPartID uint    `gorm:"primaryKey"`
    GearModPart GearModPart

    // Parts
    Tactical Tactical
}
func (auxillaryPart *AuxillaryPart) TableName() string {
    return "public.AuxillaryParts"
}


