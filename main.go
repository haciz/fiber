package fiber

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"pkg"
)

type JSONTextResponse struct {
	Message string
}

var (
	entitiesRepo pkg.EntityRepository = pkg.NewEntityMemoryRepository()
)

func main() {
	entitiesRepo.Init()

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(JSONTextResponse{Message: "Hello Fiber"})
	})

	entitiesAPI := app.Group("/entities")
	entitiesAPI.Get("/", entitiesList)

	err := app.Listen(":8080")
	if err != nil {
		log.Info().Msg("Error starting server")
	}
}

func entitiesList(c fiber.Ctx) error {
	entities, _ := entitiesRepo.List(1, 10)

	return c.JSON(entities)
}
