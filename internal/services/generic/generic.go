package generic

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/internal/database"
	"github.com/terrapi-solution/controller/internal/filter"
	"gorm.io/gorm"
	"net/http"
)

// ServiceGeneric is a generic service.
type ServiceGeneric[Model any] struct {
	conn *gorm.DB
}

// NewGenericService creates a new generic service.
func NewGenericService[Model any]() *ServiceGeneric[Model] {
	return &ServiceGeneric[Model]{conn: database.GetInstance()}
}

// List is a generic method to list entities.
func (receiver *ServiceGeneric[Model]) List(ctx *gin.Context) {
	// Get the list of entities from the database
	var entities []Model
	err := receiver.conn.Model(new(Model)).Scopes(
		filter.FilterByQuery(ctx, filter.Paginate|filter.OrderBy|filter.Search),
	).Find(&entities).Error

	// Check if there was an error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		log.Error().Err(err).Msg("failed to list the entities")
		return
	}

	// Return the entities
	ctx.JSON(http.StatusOK, entities)
}

// GetOne is a generic method to get a specific entity.
func (receiver *ServiceGeneric[Model]) GetOne(ctx *gin.Context) {
	// Get the entity ID
	entityID := ctx.Param("id")

	// Get the entity from the database
	var entity Model
	err := receiver.conn.First(&entity, entityID).Error

	// Check if there was an error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		log.Error().Err(err).Msg("failed to get the entity")
		return
	}

	// Return the entity
	ctx.JSON(http.StatusOK, entity)
}

// Delete is a generic method to delete a specific entity.
func (receiver *ServiceGeneric[Model]) Delete(ctx *gin.Context) {
	// Get the entity ID
	entityID := ctx.Param("id")

	// Delete the entity from the database
	err := receiver.conn.Delete(new(Model), entityID).Error

	// Check if there was an error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		log.Error().Err(err).Msg("failed to delete the entity")
		return
	}

	// Return a success message
	ctx.JSON(http.StatusNoContent, "")
}
