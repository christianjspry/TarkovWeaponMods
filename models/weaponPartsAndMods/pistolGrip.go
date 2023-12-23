package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type PistolGrip struct {
    gorm.Model
    PistolGripID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
    Stock *Stock
}
func (pistolGrip *PistolGrip) TableName() string {
    return "public.PistolGrip"
}
