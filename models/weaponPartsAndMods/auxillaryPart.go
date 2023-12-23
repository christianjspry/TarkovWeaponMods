package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type AuxillaryPart struct {
    gorm.Model
    AuxillaryPartID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod 

    // Parts
    Tactical Tactical
}
func (auxillaryPart *AuxillaryPart) TableName() string {
    return "public.AuxillaryParts"
}


