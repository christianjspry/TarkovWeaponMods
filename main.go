package main

import (
  "fmt"
  //orm "TarkovWeaponMods/orm"
  api "TarkovWeaponMods/api"
  //models "TarkovWeaponMods/models"
)

func main() {
    fmt.Println("Beep, boop. API started.")
    //orm.InitORM()
    api.InitAPI()
}
