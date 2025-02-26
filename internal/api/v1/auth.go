package v1

import "github.com/gin-gonic/gin"

// в gin функция v1 должна иметь указатель на c *gin.Context

func (h *Handler) signUp(c *gin.Context) {}

func (h *Handler) signIn(c *gin.Context) {}
