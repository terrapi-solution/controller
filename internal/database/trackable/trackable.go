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
	ctx.Statement.SetColumn("created_by", t.CreatedBy)
	return nil
}

// BeforeUpdate sets the updated by field
func (t *UpdatedBy) BeforeUpdate(ctx *gorm.DB) (err error) {
	ctx.Statement.SetColumn("updated_by", t.UpdatedBy)
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
