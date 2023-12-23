package models

import "gorm.io/gorm"

type Weapon struct {
    gorm.Model
    WeaponID            uint        `gorm:"primaryKey"`
    Manufacturer        string
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

    // Attachments
    // m4
    //PistolGrip          PistolGrip
    //Magazine            Magazine
    //Receiver            Receiver
    //Stock               Stock
    //ChargingHandle      ChargingHandle
    ////aks74u
    //Muzzle              Muzzle
    //GasBlock            GasBlock
    ////ak-74u
    //UnderbarrelGrenadeLauncher  UnderbarrelGrenadeLauncher
    //RearSight           RearSight
    //Mount               Mount
    ////VPO-209
    //FrontSight          FrontSight
    ////M1A
    //Barrel  Barrel
    ////P226R
    //Tactical Tactical
    ////RFB
    //Handguard Handguard

    //Chamber?
}
func (weapon *Weapon) TableName() string {
    return "public.Weapons"
}
