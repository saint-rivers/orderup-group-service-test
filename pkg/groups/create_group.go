package groups

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/saint-rivers/test-api/pkg/common/models"
)

type GroupRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

func (h handler) CreateGroup(c *gin.Context) {

	body := GroupRequest{}

	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var group models.Group

	group.Name = body.Name
	group.Description = body.Description
	group.CreatedAt = body.CreatedAt
	println(&group)

	// saves the group into the database
	// then check if an error is present
	// 		if yes, return an error
	if result := h.DB.Create(&group); result.Error != nil {
		c.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	c.JSON(http.StatusCreated, group)
}

func (h handler) GetGroup(c *gin.Context) {
	id := c.Param("id")

	var group models.Group

	if result := h.DB.Find(&group, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, group)
}
