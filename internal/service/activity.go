package service

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/terrapi-solution/controller/internal/database"
	"github.com/terrapi-solution/controller/internal/models"
	"gorm.io/gorm"
)

type Activity struct {
}

type ActivityRequest struct {
	DeploymentID uint
	Pointer      string
	Message      string
}

func NewActivityService() *Activity {
	return &Activity{}
}

func (a *Activity) Create(ctx context.Context, request ActivityRequest) (*models.Activity, error) {
	// Get the database instance
	conn := database.GetInstance()
	if conn == nil {
		return nil, errors.New("database instance is not initialized")
	}

	// Convert the request to a model
	activity := models.Activity{
		DeploymentID: request.DeploymentID,
		Pointer:      request.Pointer,
		Message:      request.Message,
	}

	// Create the activity to the database
	if err := conn.WithContext(ctx).Create(&activity).Error; err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, gorm.ErrDuplicatedKey
			}
		}
	}

	// Return the created activity
	return &activity, nil
}

func (a *Activity) List(ctx context.Context, deploymentId, page, pageSize int) ([]models.Activity, error) {
	// Define the list of activities
	var entities []models.Activity

	// Get the database instance
	conn := database.GetInstance()
	if conn == nil {
		return entities, errors.New("database instance is not initialized")
	}

	// Get the list of activities

	if err := conn.WithContext(ctx).
		Scopes(database.Paginate(page, pageSize)).
		Find(&entities, "deployment_id = ?", deploymentId).Error; err != nil {
		return entities, err
	}

	// Return the list of activities
	return entities, nil
}

func (a *Activity) Delete(ctx context.Context, id uint) error {
	// Get the database instance
	conn := database.GetInstance()
	if conn == nil {
		return errors.New("database instance is not initialized")
	}

	// Delete the activity from the database
	deleteRes := conn.WithContext(ctx).Delete(&models.Activity{}, id)
	if err := deleteRes.Error; err != nil {
		return err
	}

	// Check if the activity was deleted
	if deleteRes.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	} else {
		return nil
	}
}
