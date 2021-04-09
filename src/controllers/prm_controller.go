package controllers

import (
	config "prmgateway/src/config"
	do_reqest "prmgateway/src/do_request"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func PrmApies(app *fiber.App) {
	prmApi := app.Group("/api/prm")
	prmApi.Get("/showText/:text", showText)
	prmApi.Post("/insert", prmInsert)
	prmApi.Post("/callUrl/", prmCallUrl)
}

func Get(app *fiber.App, path string, handler fiber.Handler) {
	app.Get(path, showText)
}

func showText(c *fiber.Ctx) error {
	return c.SendString("showText get method text:" + c.Params("text"))
}

func prmInsert(c *fiber.Ctx) error {
	return c.SendString("prmInsert post method body:" + string(c.Body()))
}

func prmCallUrl(c *fiber.Ctx) error {
	configs, _ := config.GetFromEnv()
	code, myString := do_reqest.DoRequest(fiber.MethodGet, configs.BusinessApp.TestPath)
	return c.SendString("prmCallUrl body:" + string(c.Body()) + " code:" + strconv.Itoa(code) + " Res:" + myString)
}