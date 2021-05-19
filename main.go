package main

import (
	"github.com/danielberigoi/GoTodoList/todo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func view (view string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id");
		if id != "" {
			return c.Render(view, fiber.Map{"todo": todo.Read(id)})
		}
		return c.Render(view, fiber.Map{"todos": todo.Read()})
	}
}

func apiGet () fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id");
        return c.JSON(todo.Read(id))
    }
}

func apiGetAll () fiber.Handler {
	return func(c *fiber.Ctx) error {
        return c.JSON(todo.Read())
    }
}

func apiCreate () fiber.Handler {
	return func(c *fiber.Ctx) error {
		item := new(todo.Todo)
		c.BodyParser(item)
		todo.Create(item)
        return c.JSON(item)
    }
}

func apiUpdate () fiber.Handler {
	return func(c *fiber.Ctx) error {
		item := new(todo.Todo)
		item.Id = c.Params("id")
		c.BodyParser(item)
        return c.JSON(todo.Update(item))
    }
}

func main() {
	engine := html.New("./views", ".html")
    app := fiber.New(fiber.Config{Views: engine})

	app.Get("/", view("index"))

	api := app.Group("/api")
	api.Get("/", apiGetAll())
    api.Post("/", apiCreate())
    api.Get("/:id", apiGet())
    api.Put("/:id", apiUpdate())

	app.Static("/static", "./static")

    app.Listen(":3000")
}