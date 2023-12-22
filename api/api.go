package api

import (
  "github.com/gin-gonic/gin"
  //"net/http"
)

const (
    WEAPONS_ENDPOINT    =     "/weapons"
    WEAPONS_ENDPOINT_ID =     "/weapons/:id"
    WEAPONTYPES_ENDPOINT =    "/weapontypes"
    WEAPONTYPES_ENDPOINT_ID = "/weapontypes/:id"
    TYPESOFFIRE_ENDPOINT =    "/typesoffire"
    TYPESOFFIRE_ENDPOINT_ID = "/typesoffire/:id"
    CALIBERS_ENDPOINT =       "/calibers"
    CALIBERS_ENDPOINT_ID =    "/calibers/:id"
)

func WeaponsController(r *gin.Engine) {
    r.GET(WEAPONS_ENDPOINT, getWeapons)
    r.GET(WEAPONS_ENDPOINT_ID, getWeaponById)
    r.POST(WEAPONS_ENDPOINT, postWeapon) // TODO: allow for Weapon[]?
    r.PUT(WEAPONS_ENDPOINT_ID, putWeaponById)
    r.DELETE(WEAPONS_ENDPOINT_ID, deleteWeaponById)
}

func WeaponTypesController(r *gin.Engine) {
    r.GET(WEAPONTYPES_ENDPOINT, getWeaponTypes)
    r.GET(WEAPONTYPES_ENDPOINT_ID, getWeaponTypeById)
    r.POST(WEAPONTYPES_ENDPOINT, postWeaponType) // TODO: allow for WeaponType[]?
    r.PUT(WEAPONTYPES_ENDPOINT_ID, putWeaponTypeById)
    r.DELETE(WEAPONTYPES_ENDPOINT_ID, deleteWeaponTypeById)
}

func TypesOfFireController(r *gin.Engine) {
    r.GET(TYPESOFFIRE_ENDPOINT, getTypesOfFire)
    r.GET(TYPESOFFIRE_ENDPOINT_ID, getTypeOfFireById)
    r.POST(TYPESOFFIRE_ENDPOINT, postTypeOfFire) // TODO: allow for TypeOfFire[]?
    r.PUT(TYPESOFFIRE_ENDPOINT_ID, putTypeOfFireById)
    r.DELETE(TYPESOFFIRE_ENDPOINT_ID, deleteTypeOfFireById)
}

func CalibersController(r *gin.Engine) {
    r.GET(CALIBERS_ENDPOINT, getCalibers)
    r.GET(CALIBERS_ENDPOINT_ID, getCaliberById)
    r.POST(CALIBERS_ENDPOINT, postCaliber) // TODO: allow for Caliber[]?
    r.PUT(CALIBERS_ENDPOINT_ID, putCaliberById)
    r.DELETE(CALIBERS_ENDPOINT_ID, deleteCaliberById)
}

func InitAPI() *gin.Engine {
    r := gin.Default()

    WeaponsController(r) 
    WeaponTypesController(r)
    TypesOfFireController(r)
    CalibersController(r)

    r.Run()
    return r
}


