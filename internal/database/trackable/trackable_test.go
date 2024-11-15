package trackable

import (
	"context"
	"gorm.io/gorm"
	"testing"
)

func TestBeforeCreate_FieldIsSet(t *testing.T) {
	ctx := context.WithValue(context.Background(), "user_id", "user123")
	db := &gorm.DB{Statement: &gorm.Statement{Context: ctx}}
	instance := &CreatedBy{}

	err := instance.BeforeCreate(db)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if instance.CreatedBy != "user123" {
		t.Errorf("Expected CreatedBy to be 'user123', got %v", instance.CreatedBy)
	}
}

func TestBeforeCreate_FieldIsUnknownWhenUserIDNotSet(t *testing.T) {
	ctx := context.WithValue(context.Background(), "user_id", "")
	db := &gorm.DB{Statement: &gorm.Statement{Context: ctx}}
	instance := &CreatedBy{}

	err := instance.BeforeCreate(db)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if instance.CreatedBy != "unknown" {
		t.Errorf("Expected CreatedBy to be 'unknown', got %v", instance.CreatedBy)
	}
}

func TestBeforeUpdate_FieldIsSet(t *testing.T) {
	ctx := context.WithValue(context.Background(), "user_id", "user123")
	db := &gorm.DB{Statement: &gorm.Statement{Context: ctx}}
	instance := &UpdatedBy{}

	err := instance.BeforeUpdate(db)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if instance.UpdatedBy != "user123" {
		t.Errorf("Expected UpdatedBy to be 'user123', got %v", instance.UpdatedBy)
	}
}

func TestBeforeUpdate_FieldIsUnknownWhenUserIDNotSet(t *testing.T) {
	ctx := context.WithValue(context.Background(), "user_id", "")
	db := &gorm.DB{Statement: &gorm.Statement{Context: ctx}}
	instance := &UpdatedBy{}

	err := instance.BeforeUpdate(db)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if instance.UpdatedBy != "unknown" {
		t.Errorf("Expected UpdatedBy to be 'unknown', got %v", instance.UpdatedBy)
	}
}
