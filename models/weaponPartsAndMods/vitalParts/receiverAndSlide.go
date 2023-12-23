package vitalParts

import (
    "gorm.io/gorm"

    "TarkovWeaponMods/models"
    "TarkovWeaponMods/models/weaponPartsAndMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    "TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
)

type RecieverAndSlide struct {
    gorm.Model
    RecieverAndSlideID uint    `gorm:"primaryKey"`
    WeaponPartAndMod WeaponPartAndMod

    // Parts
    Tacticals []Tactical
    Mounts []Mounts

    Handguard Handguard
    FrontSight FrontSight
    RearSight RearSight
    Barrel Barrel
    Scopes Scope
}
func (recieverAndSlide *RecieverAndSlide) TableName() string {
    return "public.RecieverAndSlide"
}
