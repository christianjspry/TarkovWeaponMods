package functionMods

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type Scope struct {
    gorm.Model
    ScopeID uint    `gorm:"primaryKey"`
    GearModPart GearModPart
    SightingRange int16

    // Parts
    Tacticals []Tactical
    Mounts []Mount
    Scope Scope
}
func (scope *Scope) TableName() string {
    return "public.Scopes"
}


