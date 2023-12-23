package orm

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    mods "TarkovWeaponMods/models/weaponPartsAndMods"
)

const (
    HOST =           "localhost"
    DB_USER =        "postgres"
    DB_PASSWORD =    "postgres"
    DB_NAME =        "EFT_DB"
    DB_PORT_NUMBER = "5432"
    SSL_MODE =       "disable"          // Doubt this will change
    TIME_ZONE =      "America/New_York" // Doubt this will change 
)

func ConnectDB() (db *gorm.DB) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        HOST,
        DB_USER,
        DB_PASSWORD,
        DB_NAME,
        DB_PORT_NUMBER,
        SSL_MODE,
        TIME_ZONE,
    )
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    //db
    if err != nil {
        panic(err)
    }
    
    //db.AutoMigrate(&models.TypeOfFire{})
    return
}

func InitORM() {
    db := ConnectDB()

    db.AutoMigrate(&mods.Foregrip{})

    // TODO: There has to be a better way than this
    hasTable := db.Migrator().HasTable(&mods.Foregrip{})
    fmt.Println(hasTable)

    db.AutoMigrate(&mods.Foregrip{})

}
