package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pws-backend/domain"
)

type FilesController struct {
	FileUsecase domain.FileUsecase
}

func (pc *FilesController) Index(c *gin.Context) {
	projects, err := pc.FileUsecase.Fetch(c, 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, projects)
}

func (pc *FilesController) Details(c *gin.Context) {
	fileId, err := strconv.ParseUint(c.Param("fileId"), 10, 32)
	fmt.Println(fileId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	file, err := pc.FileUsecase.GetByID(c, uint(fileId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, file)
}
func (pc *FilesController) Delete(c *gin.Context) {
	fileId, err := strconv.ParseUint(c.Param("fileId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	file, err := pc.FileUsecase.GetByID(c, uint(fileId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	pc.FileUsecase.Delete(c, &file)
	c.JSON(http.StatusOK, file)
}
func (pc *FilesController) Create(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err)
		return
	}
	var fileE domain.File

	errUpload := pc.FileUsecase.Upload(c, file, &fileE)
	if errUpload != nil {
		log.Fatal(errUpload)
		c.JSON(http.StatusInternalServerError, errUpload)
		return
	}
	errCreate := pc.FileUsecase.Create(c, &fileE)
	if errCreate != nil {
		log.Fatal(errCreate)
		c.JSON(http.StatusInternalServerError, errCreate)
		return
	}
	c.JSON(http.StatusOK, fileE)
}
