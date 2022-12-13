package cdn

import ( 
  "github.com/gofiber/fiber/v2"
  "fmt"
  "os"
  "github.com/joho/godotenv"
  "github.com/cloudflare/cloudflare-go"
)

func main () {
  app := fiber.New()

  cf, err := cloudflare.New(os.Getenv("CLOUDFLARE_R2_TOKEN"), os.Getenv("CLOUDFLARE_API_EMAIL"))
}
