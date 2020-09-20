package model

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

// Base contains common fields for all tables
type Base struct {
	ID        int       `pg:",pk" json:"id"`
	CreatedAt time.Time `pg:"created_at" json:"created_at"`
	UpdatedAt time.Time `pg:"updated_at" json:"updated_at"`
	// DeletedAt *time.Time `json:"deleted_at,omitempty" pg:",soft_delete"`
	//DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

var _ orm.BeforeInsertHook = (*Base)(nil)
var _ orm.BeforeUpdateHook = (*Base)(nil)

// ListQuery holds company/location data used for list db queries
type ListQuery struct {
	Query string
	ID    int
}

// FilterQuery holds company/location data used for list db queries
type FilterQuery struct {
	Query  string
	Params []interface{}
}

// BeforeQuery hooks into insert operations, setting createdAt and updatedAt to current time
func (b *Base) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	fmt.Println(q.FormattedQuery())
	return c, nil
}

// BeforeInsert hooks into insert operations, setting createdAt and updatedAt to current time
func (b *Base) BeforeInsert(c context.Context) (context.Context, error) {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	return c, nil
}

// BeforeUpdate hooks into update operations, setting updatedAt to current time
func (b *Base) BeforeUpdate(c context.Context) (context.Context, error) {
	b.UpdatedAt = time.Now()
	return c, nil
}

// ToString func
func (b *Base) ToString() string {
	str, _ := json.Marshal(b)
	return string(str)
}
