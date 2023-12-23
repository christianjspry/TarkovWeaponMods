package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type Foregrip struct {
    gorm.Model
    //ForegripID uint    `gorm:"primaryKey"`
    WeaponPartAndMod

    // Parts
}
func (forgrip *Foregrip) TableName() string {
    return "public.Foregrips"
}


