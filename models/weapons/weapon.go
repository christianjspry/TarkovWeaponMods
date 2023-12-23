package weapons

import (
    "gorm.io/gorm"
    "TarkovWeaponMods/models"
)

type Weapon struct {
    gorm.Model
    WeaponID            uint        `gorm:"primaryKey"`
    Item Item
    Manufacturer        string
    Ergonomics          float32
    Accuracy            float32
    SightingRange       int32
    VerticalRecoil      float32
    HorizontalRecoil    float32
    MuzzleVelocity      int32
    FireRate            int32
    EffectiveDistance   int32

    // AR, Carbine, SMG, etc.. TODO: field validation?
    WeaponType          string      //`gorm:"embedded"`

    // Type of Fire
    SingleFire bool
    BurstFire  bool
    FullAuto   bool

    // Embeds
    Caliber Caliber     `gorm:"embedded"`
    // m4
    PistolGrip PistolGrip  `gorm:"embedded"`
    Magazine Magazine  `gorm:"embedded"`
    PistolGrip PistolGrip  `gorm:"embedded"`
    Receiver Receiver
    Stock Stock
    ChargingHandle ChargingHandle
    
    ////aks74u
    Muzzle Muzzle
    GasBlock GasBlock

    //ak-74u
    UnderbarrelGrenadeLauncher UnderbarrelGrenadeLauncher
    RearSight RearSight
    Mount     Mount
    //VPO-209
    FrontSight FrontSight
    //M1A
    Barrel Barrel
    //P226R
    Tactical Tactical
    //RFB
    Handguard Handguard

    //Chamber?
}
func (weapon *Weapon) TableName() string {
    return "public.Weapons"
}

//type WeaponType struct {
//    AssaultCarbine string
//    AssaultRifle string
//    BoltActionRifle string
//    GrenadeLauncher string
//    MachineGun string
//    MarksmanRifle string
//    MeleeWeapon string
//    Pistol string
//    SMG string
//    Shotgun string
//    SpecialWeapon string
//    Throwable string
//}
//
//const WeaponTypes WeaponType = WeaponType{
//    "Assault Carbine",
//    "Assault Rifle",
//    "Bolt-action Rifle",
//    "Grenade Launcher",
//    "Machine Gun",
//    "Marksman Rifle",
//    "Melee Weapon",
//    "Pistol",
//    "SMG",
//    "Shotgun",
//    "Special Weapon",
//    "Throwable",
//}
