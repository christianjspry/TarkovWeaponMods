package api

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func getTypesOfFire(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "getTypesOfFire",
    })
}

func getTypeOfFireById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "getTypeOfFireById",
    })
}

func postTypeOfFire(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "postTypeOfFire",
    })
}

func putTypeOfFireById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "putTypeOfFireById",
    })
}

func deleteTypeOfFireById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "deleteTypeOfFireById",
    })
}
