package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type Bipod struct {
    gorm.Model
    BipodID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
}
func (bipod *Bipod) TableName() string {
    return "public.Bipods"
}


