package orm

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    models "TarkovWeaponMods/models"
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

//func init_orm() *gorm.DB {
func InitORM() {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        HOST,
        DB_USER,
        DB_PASSWORD,
        DB_NAME,
        DB_PORT_NUMBER,
        SSL_MODE,
        TIME_ZONE,
    )
    db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})


    db.AutoMigrate(&models.TypeOfFire{})

    db.Create(&models.TypeOfFire{
        SingleFire: true,
        FullAuto: false,
        BurstFire: false,
    })

    typeOfFire := models.TypeOfFire{}

    db.First(&typeOfFire, 1)

    //return db
}
