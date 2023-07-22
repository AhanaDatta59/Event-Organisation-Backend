package Controllers

import (
  "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

  "main/Database"
	"main/Models"
)

func AllEvent(c *fiber.Ctx)  error {

  cookie:=c.Cookies("jwt")

  _, err:=jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token)(interface{}, error) {
    return[]byte(SecretKey),nil
  })
  
   if err!=nil{
     c.Status(fiber.StatusUnauthorized)
     return c.JSON(fiber.Map{
       "message":"Unauthorized",
     })
   }

    if 

  
  
  
}
