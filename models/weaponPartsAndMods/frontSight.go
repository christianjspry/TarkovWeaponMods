package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type FrontSight struct {
    gorm.Model
    FrontSightID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod
    SightingRange int16

    // Parts
    Mount Mount
    Scope Scope
}
func (frontSight *FrontSight) TableName() string {
    return "public.FrontSights"
}


