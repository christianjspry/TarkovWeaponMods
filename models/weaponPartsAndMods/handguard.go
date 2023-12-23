package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type Handguard struct {
    gorm.Model
    HandguardID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
    Tacticals []Tactical
    Mounts []Mount

    FrontSight FrontSight
    Scopes Scope
    Foregrip Foregrip
}
func (handguard *Handguard) TableName() string {
    return "public.Handguard"
}
