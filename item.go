// Package entities holds all the entities that are shared across subdomains
package tavern

import "github.com/google/uuid"

// Item is an entity that represents a item in all domains
type Item struct {
	// ID is an identifier of the entity
	ID          uuid.UUID
	Name        string
	Description string
}
