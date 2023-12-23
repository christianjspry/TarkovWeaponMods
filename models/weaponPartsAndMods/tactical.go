package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type Tactical struct {
    gorm.Model
    TacticalID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod 

    // Parts
}
func (tactical *Tactical) TableName() string {
    return "public.Tacticals"
}


