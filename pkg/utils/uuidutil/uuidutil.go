package uuidutil

import "github.com/google/uuid"

// GenerateUUID generates a new unique identifier (UUID)
func GenerateUUID() string {
	return uuid.NewString()
}
