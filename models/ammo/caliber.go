package ammo

import "gorm.io/gorm"

type Caliber struct {
    gorm.Model
    CaliberID   uint    `gorm:"primaryKey"` 
    Name        string
}
func (caliber *Caliber) TableName() string {
    return "public.Calibers"
}
