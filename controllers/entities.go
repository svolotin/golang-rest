// Copyright (c) 2023 ...

package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /entity
func GetEntity(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Attributes to return": c.Param("entityId")})
}
