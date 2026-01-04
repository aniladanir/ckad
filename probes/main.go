package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gofiber/fiber/v2"
)

var isReady = atomic.Bool{}

func Startup() {
	// simulate initializing external dependencies
	time.Sleep(time.Second * 15)
	isReady.Store(true)
}

func IsReady() bool {
	return isReady.Load()
}

func main() {
	wg := sync.WaitGroup{}

	wg.Go(func() {
		Startup()
	})

	wg.Go(func() {
		app := fiber.New()

		app.Get("/health", Health)
		app.Get("/ready", Ready)

		if err := app.Listen(":3000"); err != nil {
			fmt.Println(err)
		}
	})

	wg.Wait()
	os.Exit(1)
}

func Health(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusOK)
}

func Ready(ctx *fiber.Ctx) error {
	if !IsReady() {
		return ctx.SendStatus(http.StatusServiceUnavailable)
	}
	return ctx.SendStatus(http.StatusOK)
}
