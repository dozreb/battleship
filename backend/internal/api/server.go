package api

import (
	"net/http"

	"battleship/backend/internal/game"
	"github.com/gin-gonic/gin"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	r.Use(corsMiddleware())

	svc := game.NewService()

	api := r.Group("/api")
	{
		api.POST("/game/new", func(c *gin.Context) {
			c.JSON(http.StatusCreated, svc.NewGame())
		})

		api.GET("/game/:id", func(c *gin.Context) {
			view, err := svc.GetGameView(c.Param("id"))
			if err != nil {
				status, payload := mapError(err)
				c.JSON(status, payload)
				return
			}
			c.JSON(http.StatusOK, view)
		})

		api.POST("/game/:id/shot", func(c *gin.Context) {
			var req game.ShotRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
				return
			}
			res, err := svc.PlayerShot(c.Param("id"), req)
			if err != nil {
				status, payload := mapError(err)
				c.JSON(status, payload)
				return
			}
			c.JSON(http.StatusOK, res)
		})
	}

	return r
}

func mapError(err error) (int, gin.H) {
	switch err {
	case game.ErrNotFound:
		return http.StatusNotFound, gin.H{"error": err.Error()}
	case game.ErrGameOver:
		return http.StatusConflict, gin.H{"error": err.Error()}
	case game.ErrNotPlayersTurn:
		return http.StatusConflict, gin.H{"error": err.Error()}
	case game.ErrInvalidCoordinate:
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	case game.ErrAlreadyShot:
		return http.StatusConflict, gin.H{"error": err.Error()}
	default:
		return http.StatusInternalServerError, gin.H{"error": "internal server error"}
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
