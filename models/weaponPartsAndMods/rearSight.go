package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type RearSight struct {
    gorm.Model
    RearSightID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod
    SightingRange int16

    // Parts
    Mount Mount
    Scope Scope
}
func (rearSight *RearSight) TableName() string {
    return "public.RearSights"
}


