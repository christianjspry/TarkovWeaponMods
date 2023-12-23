package weaponPartsAndMods

import (
    models "TarkovWeaponMods/models"
)

type WeaponPartAndMod struct {
    models.Item
    Category string // Vital Part, Functional Mod, etc TODO: struct?
    Manufacturer        string
    SizeChangeDirection string
    SizeChangeMagnitude int16
    ModdedInRaid        bool
    Recoil              float32
    Ergonomics float32
    Heat float32
    Cooling float32
    DurabilityBurn float32
    //CompatibleWith []GearModPart?
}
