package api

import (
    //"TarkovWeaponMods/models/weaponPartsAndMods/functionalMods"
    //"TarkovWeaponMods/models/weaponPartsAndMods/gearMods"
    //"TarkovWeaponMods/models/weaponPartsAndMods/vitalParts"
  orm "TarkovWeaponMods/orm"
  "github.com/gin-gonic/gin"
  "net/http"
)

func getForegrips(c *gin.Context) {
    //db := orm.ConnectDB()
    orm.ConnectDB()
    
    c.JSON(http.StatusOK, gin.H{
        "message": "getForegrips",
    })
}

func getForegripById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "getForegripById",
    })
}

func postForegrip(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "postForegrip",
    })
}

func putForegripById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "putForegripById",
    })
}

func deleteForegripById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "deleteForegripById",
    })
}
