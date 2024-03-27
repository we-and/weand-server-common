package servers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ListenHTTPServer(app *fiber.App, address string, serverDesc string) error {
	fmt.Printf("[%v] API server HTTP listen to %v\n", serverDesc, address)
	return app.Listen(address)
}
