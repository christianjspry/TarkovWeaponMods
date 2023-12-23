package api

import (
  "github.com/gin-gonic/gin"

  //"TarkovWeaponMods/orm"
  //"net/http"
)

func InitAPI() *gin.Engine {
    r := gin.Default()
    
    InitControllers(r)

    r.Run()
    return r
}
