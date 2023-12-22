package api

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func getCalibers(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "getCalibers",
    })
}

func getCaliberById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "getCaliberById",
    })
}
func postCaliber(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "postCaliber",
    })
}
func putCaliberById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "putCaliberById",
    })
}
func deleteCaliberById(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "deleteCaliberById",
    })
}

