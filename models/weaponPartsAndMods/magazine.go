package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type Magazine struct {
    gorm.Model
    MagazineID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod
    MaxCount int
    CheckSpeedModifier float32
    FailureToFeedChance string

    // Parts
}
func (magazine *Magazine) TableName() string {
    return "public.Magazines"
}


