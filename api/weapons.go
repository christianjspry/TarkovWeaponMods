package api

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func getWeapons(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "getWeapons",
    })
}

func getWeaponById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "getWeaponById",
    })
}

func postWeapon(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "postWeapon",
    })
}

func putWeaponById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "putWeaponById",
    })
}

func deleteWeaponById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "deleteWeaponById",
    })
}

