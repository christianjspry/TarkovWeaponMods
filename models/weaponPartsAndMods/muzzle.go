package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type Muzzle struct {
    gorm.Model
    MuzzleID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod
    Loudness int16

    // Parts
    Muzzle *Muzzle
}
func (muzzle *Muzzle) TableName() string {
    return "public.Muzzles"
}


