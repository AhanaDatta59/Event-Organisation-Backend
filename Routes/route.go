package Routes

import(
  "github.com/gofiber/fiber/v2"
  "main/Controllers"
)

func SetUp(app *fiber.App){
  app.Post("/kevents/register", Controllers.Register)
  app.Post("/kevents/login", Controllers.Login)
  app.Post("/kevents/logout", Controllers.Logout)
  app.Get("/kevnts/allevent", Controllers.AllEvent)
  app.Get("/kevents/yourevent", Controllers.YourEvent)
  // app.Post("/",(req,res) => {
    
  // })
}