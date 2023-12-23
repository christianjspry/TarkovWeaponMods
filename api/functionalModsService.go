package api

import (
    "fmt"
    orm "TarkovWeaponMods/orm"
    "github.com/gin-gonic/gin"
    "net/http"
    mods "TarkovWeaponMods/models/weaponPartsAndMods"
    //models "TarkovWeaponMods/models"
)

func getForegrips(c *gin.Context) {
    db := orm.ConnectDB()
    
    fmt.Println(db) 
    
    c.JSON(http.StatusOK, gin.H{
        "message": "getForegrips",
    })
}

func getForegripById(c *gin.Context) {
    //id := c.Param("id")

    c.JSON(http.StatusOK, gin.H{
        "message": "getForegripById",
    })
}

func postForegrip(c *gin.Context) {

    var foregrip mods.Foregrip
    if err := c.BindJSON(&foregrip); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    fmt.Println(foregrip)
    
    db := orm.ConnectDB()
    db.Create(&foregrip)

    c.JSON(http.StatusCreated, foregrip)
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
