package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pws-backend/domain"
)

type ProjectsController struct {
	ProjectUsecase domain.ProjectUsecase
}

func (pc *ProjectsController) Index(c *gin.Context) {
	projects, err := pc.ProjectUsecase.Fetch(c, 5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, projects)
}

func (pc *ProjectsController) Details(c *gin.Context) {
	c.JSON(http.StatusUnprocessableEntity, nil)
}

func (pc *ProjectsController) Update(c *gin.Context) {
	c.JSON(http.StatusUnprocessableEntity, nil)
}

func (pc *ProjectsController) Create(c *gin.Context) {
	var project domain.Project
	err := c.ShouldBind(&project)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = pc.ProjectUsecase.Create(c, &project)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, project)
}
