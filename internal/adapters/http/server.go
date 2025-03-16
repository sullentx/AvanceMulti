package http

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/MarcosBrindis/hexagonal-api/internal/app"
)

func StartServer() {
    r := gin.Default()
    eventService := app.NewEventService()

    r.GET("/detect-movement", func(c *gin.Context) {
        event := eventService.DetectMovement()
        c.JSON(http.StatusOK, event)
    })

    r.GET("/detect-gas-leak", func(c *gin.Context) {
        event := eventService.DetectGasLeak()
        c.JSON(http.StatusOK, event)
    })

    r.GET("/detect-window-open", func(c *gin.Context) {
        event := eventService.DetectWindowOpen()
        c.JSON(http.StatusOK, event)
    })

    go func() {
        for {
            time.Sleep(10 * time.Second)
            event := eventService.DetectMovement()
            // Aquí puedes enviar la notificación del evento
            println(event.Type, event.Timestamp.String())
        }
    }()

    go func() {
        for {
            time.Sleep(20 * time.Second)
            event := eventService.DetectGasLeak()
            // Aquí puedes enviar la notificación del evento
            println(event.Type, event.Timestamp.String())
        }
    }()

    go func() {
        for {
            time.Sleep(30 * time.Second)
            event := eventService.DetectWindowOpen()
            // Aquí puedes enviar la notificación del evento
            println(event.Type, event.Timestamp.String())
        }
    }()

    r.Run(":8080")
}