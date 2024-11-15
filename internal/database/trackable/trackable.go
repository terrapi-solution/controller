package trackable

import "gorm.io/gorm"

// CreatedBy represents the created by field
type CreatedBy struct {
	CreatedBy string `gorm:"not null"`
}

// UpdatedBy represents the updated by field
type UpdatedBy struct {
	UpdatedBy string
}

// BeforeCreate sets the created by field
func (t *CreatedBy) BeforeCreate(ctx *gorm.DB) (err error) {
	t.CreatedBy = getUserID(ctx)
	return nil
}

// BeforeUpdate sets the updated by field
func (t *UpdatedBy) BeforeUpdate(ctx *gorm.DB) (err error) {
	t.UpdatedBy = getUserID(ctx)
	return nil
}

// getUserID returns the user id from the context
func getUserID(ctx *gorm.DB) string {
	userId, ok := ctx.Statement.Context.Value("user_id").(string)
	if !ok || userId == "" {
		return "unknown"
	}
	return userId
}
