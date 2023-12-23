package weaponPartsAndMods

import (
    "gorm.io/gorm"
    ammo "TarkovWeaponMods/models/ammo"
)

type Launcher struct {
    gorm.Model
    LauncherID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod
    Caliber ammo.Caliber

    // Parts
}
func (launcher *Launcher) TableName() string {
    return "public.Launchers"
}


