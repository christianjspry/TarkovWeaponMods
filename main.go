package main

import (
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "github.com/gin-gonic/gin"
  "net/http"
  "fmt"
  "models"
)

const WEAPONS_ENDPOINT = "/weapons"
const WEAPONS_ENDPOINT_ID = "/weapons/:id"

func init_orm() *gorm.DB {
    dsn := "host=localhost user=postgres password=postgres dbname=EFT_DB port=5432 sslmode=disable TimeZone=America/New_York"
    db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    return db
}

func init_api() *gin.Engine {
    r := gin.Default()
    
    //Endpoints

    // CRUD for /Weapons endpoint
    r.GET(WEAPONS_ENDPOINT, get_Weapons)
    r.GET(WEAPONS_ENDPOINT_ID, get_WeaponById)
    r.POST(WEAPONS_ENDPOINT, post_Weapon) // TODO: allow for Weapon[]?
    r.PUT(WEAPONS_ENDPOINT_ID, put_WeaponById)
    r.DELETE(WEAPONS_ENDPOINT_ID, delete_WeaponById)


    r.Run()
    return r
}

func get_Weapons(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "get_Weapons",
    })
}

func get_WeaponById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "get_WeaponById",
    })
}
func post_Weapon(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "post_Weapon",
    })
}
func put_WeaponById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "put_WeaponById",
    })
}
func delete_WeaponById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "delete_WeaponById",
    })
}


func main() {

    fmt.Println("Hello, World!")
    //init_orm()
    init_api()

    //db.AutoMigrate(&TypeOfFire{})

    //db.Create(&TypeOfFire{
    //    SingleFire: true,
    //    FullAuto: false,
    //    BurstFire: false,
    //})

    //typeOfFire := TypeOfFire{}

    //db.First(&typeOfFire, 1)

    //fmt.Println(typeOfFire)

}

//type Weapon struct {
//    gorm.Model
//    WeaponID            uint        `gorm:"primaryKey"`
//    Name                string
//    WeaponType          WeaponType  `gorm:"embedded"`
//    Caliber             Caliber     `gorm:"embedded"`
//    Weight              float32      
//    Ergonomics          float32
//    Accuracy            float32
//    SightingRange       int32
//    VerticalRecoil      float32
//    HorizontalRecoil    float32
//    MuzzleVelocity      int32
//    TypeOfFire          TypeOfFire  `gorm:"embedded"`
//    FireRate            int32
//    EffectiveDistance   int32
//}
//func (weapon *Weapon) TableName() string {
//    return "public.Weapons"
//}
//
//type WeaponType struct {
//    gorm.Model
//    WeaponTypeID    uint    `gorm:"primaryKey"`
//    Name            string
//}
//func (weaponType *WeaponType) TableName() string {
//    return "public.WeaponTypes"
//}
//
//type TypeOfFire struct {
//    gorm.Model
//    TypeOfFireID    uint    `gorm:"primaryKey"`
//    SingleFire      bool 
//    FullAuto        bool
//    BurstFire       bool
//}
//func (typeOfFire *TypeOfFire) TableName() string {
//    return "public.TypesOfFire"
//}
//
//type Caliber struct {
//    gorm.Model
//    CaliberID   uint    `gorm:"primaryKey"` 
//    Name        string
//}
//func (caliber *Caliber) TableName() string {
//    return "public.Calibers"
//}
