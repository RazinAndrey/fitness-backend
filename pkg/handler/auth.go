package handler

import "github.com/gin-gonic/gin"

// в gin функция handler должна иметь указатель на c *gin.Context

func (h *Handler) signUp(c *gin.Context) {}

func (h *Handler) signIn(c *gin.Context) {}
