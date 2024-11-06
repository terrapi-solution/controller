package user

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	domainErrors "github.com/terrapi-solution/controller/domain/errors"
	"gorm.io/gorm"
	"time"
)

// Store struct manages interactions with authors store
type Store struct {
	db *gorm.DB
}

// New creates a new Store struct
func New(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

// Create creates a new user entry in the database
func (s *Store) Create(user User) (User, error) {
	// Update object
	user.Status = true
	user.LastActive = time.Time.UTC(time.Now())

	// Create the user in the database
	if err := s.db.Create(&user).Error; err != nil {
		var duplicateEntryError = &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			return User{}, domainErrors.NewConflict(err, fmt.Sprintf("Error creating User with the email: %s", user.Email), "UserStore.Create")
		}
		return User{}, domainErrors.NewInternal(err, "Error executing SQL query", "UserStore.Create")
	}
	return user, nil
}

// Read retrieves a user entry from the database
func (s *Store) Read(id int) (User, error) {
	var user User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, domainErrors.NewNotFound(err, fmt.Sprintf("Error obtaining User with the Id: %d", id), "UserStore.Read")
		}
		return User{}, domainErrors.NewInternal(err, "Error executing SQL query", "UserStore.Read")
	}
	return user, nil
}

// ReadBySubject retrieves a user entry from the database by subject
func (s *Store) ReadBySubject(subject string) (User, error) {
	var user User
	if err := s.db.Where("subject = ?", subject).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return User{}, domainErrors.NewNotFound(err, fmt.Sprintf("Error obtaining User with the Subject: %s", subject), "UserStore.ReadBySubject")
		}
		return User{}, domainErrors.NewInternal(err, "Error executing SQL query", "UserStore.ReadBySubject")
	}
	return user, nil
}

// Delete deletes a user entry from the database
func (s *Store) Delete(id int) error {
	if err := s.db.Delete(&User{}, id).Error; err != nil {
		return domainErrors.NewInternal(err, "Error executing SQL query", "UserStore.Delete")
	}
	return nil
}

// List retrieves all user entries from the database
func (s *Store) List() ([]User, error) {
	var users []User
	if err := s.db.Find(&users).Error; err != nil {
		return nil, domainErrors.NewInternal(err, "Error executing SQL query", "UserStore.List")
	}
	return users, nil
}

// Exists if a user exists in the database
func (s *Store) Exists(id int) bool {
	var count int64
	if err := s.db.Model(&User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false
	}
	return count > 0
}

// ReadByEmail retrieves a user entry from the database by email
func (s *Store) ReadByEmail(email string) (User, error) {
	user := &User{}
	result := s.db.Where("email = ?", email).First(user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return User{}, domainErrors.NewNotFound(result.Error, fmt.Sprintf("Error obtaining User with the email: %s", email), "UserStore.ReadByEmail")
		}
		return User{}, domainErrors.NewInternal(result.Error, "Error executing SQL query", "UserStore.ReadByEmail")
	}
	return *user, nil
}

// UpdateLastActive updates the last active field of a user entry in the database
func (s *Store) UpdateLastActive(id int) error {
	if err := s.db.Model(&User{}).Where("id = ?", id).Update("last_active", time.Time.UTC(time.Now())).Error; err != nil {
		return domainErrors.NewInternal(err, "Error updating User with the Id: %d", "UserStore.UpdateLastActive")
	}
	return nil
}

// UpdateStatus updates the is active field of a user entry in the database
func (s *Store) UpdateStatus(id int, isActive bool) error {
	if err := s.db.Model(&User{}).Where("id = ?", id).Update("is_active", isActive).Error; err != nil {
		return domainErrors.NewInternal(err, "Error updating User with the Id: %d", "UserStore.UpdateStatus")
	}
	return nil
}
