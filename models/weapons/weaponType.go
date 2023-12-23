package weapons

import "gorm.io/gorm"

type WeaponType struct {
    gorm.Model
    WeaponTypeID    uint    `gorm:"primaryKey"`
    Name            string
}
func (weaponType *WeaponType) TableName() string {
    return "public.WeaponTypes"
}
