package models

import "gorm.io/gorm"

type Weapon struct {
    gorm.Model
    WeaponID            uint        `gorm:"primaryKey"`
    Name                string
    WeaponType          WeaponType  `gorm:"embedded"`
    Caliber             Caliber     `gorm:"embedded"`
    Weight              float32      
    Ergonomics          float32
    Accuracy            float32
    SightingRange       int32
    VerticalRecoil      float32
    HorizontalRecoil    float32
    MuzzleVelocity      int32
    TypeOfFire          TypeOfFire  `gorm:"embedded"`
    FireRate            int32
    EffectiveDistance   int32
}
func (weapon *Weapon) TableName() string {
    return "public.Weapons"
}
