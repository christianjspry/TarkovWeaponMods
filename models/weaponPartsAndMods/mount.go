package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type Mount struct {
    gorm.Model
    MountID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
    Tacticals []Tactical
    Foregrip Foregrip
    Scope Scope
    Mount *Mount
    
}
func (mount *Mount) TableName() string {
    return "public.Mounts"
}


