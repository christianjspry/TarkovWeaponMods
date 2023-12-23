package api

import (
  "github.com/gin-gonic/gin"

  //"TarkovWeaponMods/orm"
  //"net/http"
)

const (
    ID_PATH = "/:id"

    WEAPONS_ENDPOINT = "/weapons"
    WEAPONS_ENDPOINT_ID = WEAPONS_ENDPOINT + ID_PATH

    /* ----  Weapon parts & mods ----- */
    WEAPONPARTSANDMODS_ENDPOINT = "/mods"

    FOREGRIPS_ENDPOINT = WEAPONPARTSANDMODS_ENDPOINT+ "/foregrips"
    FOREGRIPS_ENDPOINT_ID = FOREGRIPS_ENDPOINT + ID_PATH

    CALIBERS_ENDPOINT = "/calibers"
    CALIBERS_ENDPOINT_ID = CALIBERS_ENDPOINT + ID_PATH
)

func WeaponsController(r *gin.Engine) {
    r.GET(WEAPONS_ENDPOINT, getWeapons)
    r.GET(WEAPONS_ENDPOINT_ID, getWeaponById)
    r.POST(WEAPONS_ENDPOINT, postWeapon) // TODO: allow for Weapon[]?
    r.PUT(WEAPONS_ENDPOINT_ID, putWeaponById)
    r.DELETE(WEAPONS_ENDPOINT_ID, deleteWeaponById)
}

func ForegripsController(r *gin.Engine) {
    r.GET(FOREGRIPS_ENDPOINT, getForegrips)
    r.GET(FOREGRIPS_ENDPOINT_ID, getForegripById)
    r.POST(FOREGRIPS_ENDPOINT, postForegrip) // TODO: allow for Foregrip[]?
    r.PUT(FOREGRIPS_ENDPOINT_ID, putForegripById)
    r.DELETE(FOREGRIPS_ENDPOINT_ID, deleteForegripById)
}

func CalibersController(r *gin.Engine) {
    r.GET(CALIBERS_ENDPOINT, getCalibers)
    r.GET(CALIBERS_ENDPOINT_ID, getCaliberById)
    r.POST(CALIBERS_ENDPOINT, postCaliber) // TODO: allow for Caliber[]?
    r.PUT(CALIBERS_ENDPOINT_ID, putCaliberById)
    r.DELETE(CALIBERS_ENDPOINT_ID, deleteCaliberById)
}

func InitControllers(r *gin.Engine) {
    ForegripsController(r) 
}
