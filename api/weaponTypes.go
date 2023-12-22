package api

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func getWeaponTypes(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "getWeaponTypes",
    })
}

func getWeaponTypeById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "getWeaponTypeById",
    })
}
func postWeaponType(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "postWeaponType",
    })
}
func putWeaponTypeById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "putWeaponTypeById",
    })
}
func deleteWeaponTypeById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "deleteWeaponTypeById",
    })
}

