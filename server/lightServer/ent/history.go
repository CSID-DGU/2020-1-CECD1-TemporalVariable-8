// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"lightServer/ent/history"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/sql"
)

// History is the model entity for the History schema.
type History struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OpenAt holds the value of the "open_at" field.
	OpenAt time.Time `json:"open_at,omitempty"`
	// CloseAt holds the value of the "close_at" field.
	CloseAt time.Time `json:"close_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the HistoryQuery when eager-loading is set.
	Edges HistoryEdges `json:"edges"`
}

// HistoryEdges holds the relations/edges for other nodes in the graph.
type HistoryEdges struct {
	// Of holds the value of the of edge.
	Of []*Restaurant
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OfOrErr returns the Of value or an error if the edge
// was not loaded in eager-loading.
func (e HistoryEdges) OfOrErr() ([]*Restaurant, error) {
	if e.loadedTypes[0] {
		return e.Of, nil
	}
	return nil, &NotLoadedError{edge: "of"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*History) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullTime{},  // open_at
		&sql.NullTime{},  // close_at
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the History fields.
func (h *History) assignValues(values ...interface{}) error {
	if m, n := len(values), len(history.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	h.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field open_at", values[0])
	} else if value.Valid {
		h.OpenAt = value.Time
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field close_at", values[1])
	} else if value.Valid {
		h.CloseAt = value.Time
	}
	return nil
}

// QueryOf queries the of edge of the History.
func (h *History) QueryOf() *RestaurantQuery {
	return (&HistoryClient{config: h.config}).QueryOf(h)
}

// Update returns a builder for updating this History.
// Note that, you need to call History.Unwrap() before calling this method, if this History
// was returned from a transaction, and the transaction was committed or rolled back.
func (h *History) Update() *HistoryUpdateOne {
	return (&HistoryClient{config: h.config}).UpdateOne(h)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (h *History) Unwrap() *History {
	tx, ok := h.config.driver.(*txDriver)
	if !ok {
		panic("ent: History is not a transactional entity")
	}
	h.config.driver = tx.drv
	return h
}

// String implements the fmt.Stringer.
func (h *History) String() string {
	var builder strings.Builder
	builder.WriteString("History(")
	builder.WriteString(fmt.Sprintf("id=%v", h.ID))
	builder.WriteString(", open_at=")
	builder.WriteString(h.OpenAt.Format(time.ANSIC))
	builder.WriteString(", close_at=")
	builder.WriteString(h.CloseAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Histories is a parsable slice of History.
type Histories []*History

func (h Histories) config(cfg config) {
	for _i := range h {
		h[_i].config = cfg
	}
}
