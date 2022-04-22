package handlers

import (
	"go-clean-arch/internal/business"

	"github.com/gin-gonic/gin"
)

type businessHandler struct {
	useCase business.BusinessUseCaseCRUD
}

func NewBusinessHandler(useCase business.BusinessUseCaseCRUD) *businessHandler {
	return &businessHandler{useCase}
}

func (b businessHandler) HandleCreate(c *gin.Context) {
	
}
