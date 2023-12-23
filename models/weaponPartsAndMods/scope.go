package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type Scope struct {
    gorm.Model
    ScopeID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod
    SightingRange int16

    // Parts
    Tacticals []Tactical
    Mounts []Mount
    Scope *Scope
}
func (scope *Scope) TableName() string {
    return "public.Scopes"
}


