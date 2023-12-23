package weaponPartsAndMods

import (
    "gorm.io/gorm"
)

type RecieverAndSlide struct {
    gorm.Model
    RecieverAndSlideID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
    Tacticals []Tactical
    Mounts []Mount

    Handguard Handguard
    FrontSight FrontSight
    RearSight RearSight
    Barrel Barrel
    Scopes Scope
}
func (recieverAndSlide *RecieverAndSlide) TableName() string {
    return "public.RecieverAndSlide"
}
