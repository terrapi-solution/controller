package users

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/terrapi-solution/controller/data/user"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"github.com/terrapi-solution/controller/internal/filter"
	"gorm.io/gorm"
)

// Service struct manages interactions with activities store
type Service struct {
	db   *gorm.DB
	user *user.Store
}

// New creates a new Service struct
func New(db *gorm.DB) Service {
	return Service{
		db:   db,
		user: user.New(db),
	}
}

// Read retrieves a user entry from the database
func (s *Service) Read(id int) (user.User, error) {
	dbUser, err := s.user.Read(id)
	if err != nil {
		return user.User{}, err
	}

	return dbUser, nil
}

// Delete deletes a user entry from the database
func (s *Service) Delete(id int) error {
	if !s.user.Exists(id) {
		return domainErrors.NewNotFound(nil, "User not found", "UserService.Delete")
	}

	if err := s.user.Delete(id); err != nil {
		return err
	}

	log.Info().Msgf("User %d deleted successfully", id)
	return nil
}

// List retrieves all user entries from the database
func (s *Service) List() ([]user.User, error) {
	users, err := s.user.List()
	if err != nil {
		return nil, err
	}

	return users, nil
}

// PaginateList retrieves a paginated list of user entries from the database
func (s *Service) PaginateList(ctx *gin.Context) (*[]user.User, error) {
	var entries []user.User
	err := s.db.Model(&user.User{}).Scopes(
		filter.FilterByQuery(ctx, filter.Paginate|filter.OrderBy|filter.Search|filter.Filter),
	).Find(&entries).Error

	if err != nil {
		return nil, domainErrors.NewInternal(err, "Error executing SQL query", "UserService.PaginateList")
	}

	return &entries, nil
}

// UpdateLastLogin updates the last login of a user in the database
func (s *Service) UpdateLastLogin(id int) error {
	return s.user.UpdateLastActive(id)
}

// UpdateStatus updates the is active field of a user in the database
func (s *Service) UpdateStatus(id int, isActive bool) error {
	return s.user.UpdateStatus(id, isActive)
}

// Me retrieves the user information of the currently logged-in user
func (s *Service) Me(ctx *gin.Context) (user.User, error) {
	// Get the user claims from the context
	subject, exists := ctx.Get("subject")
	if !exists {
		return user.User{}, domainErrors.NewInternal(nil, "Unable to retrieve subject from context", "UserService.Me")
	}

	// Get the user from the database
	dbUser, err := s.user.ReadBySubject(subject.(string))
	if err != nil {
		return user.User{}, err
	}

	// Return current user information
	return dbUser, nil
}
