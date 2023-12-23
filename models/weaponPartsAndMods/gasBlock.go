package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type GasBlock struct {
    gorm.Model
    GasBlockID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
    Scopes []Scope
    Tacticals []Tactical
    Mounts []Mount
}
func (gasBlock *GasBlock) TableName() string {
    return "public.GasBlock"
}
