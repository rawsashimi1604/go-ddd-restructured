// Package entities holds all the value objects that are shared across subdomains
package valueobject

import (
	"time"

	"github.com/google/uuid"
)

// Transaction is a valueobject because it has no identifer and immutable
type Transaction struct {
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
