package models

import "gorm.io/gorm"

type TypeOfFire struct {
    gorm.Model
    TypeOfFireID    uint    `gorm:"primaryKey"`
    SingleFire      bool 
    FullAuto        bool
    BurstFire       bool
}
func (typeOfFire *TypeOfFire) TableName() string {
    return "public.TypesOfFire"
}


