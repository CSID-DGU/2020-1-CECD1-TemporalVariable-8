// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"lightServer/ent/file"
	"lightServer/ent/menu"
	"lightServer/ent/restaurant"
	"strings"

	"github.com/Rhymond/go-money"
	"github.com/facebook/ent/dialect/sql"
)

// Menu is the model entity for the Menu schema.
type Menu struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// IsOption holds the value of the "isOption" field.
	IsOption bool `json:"isOption,omitempty"`
	// Price holds the value of the "price" field.
	Price *money.Money `json:"price,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MenuQuery when eager-loading is set.
	Edges       MenuEdges `json:"edges"`
	menu_owner  *int
	menu_images *int
}

// MenuEdges holds the relations/edges for other nodes in the graph.
type MenuEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Restaurant
	// Category holds the value of the category edge.
	Category []*Category
	// Images holds the value of the images edge.
	Images *File
	// Options holds the value of the options edge.
	Options []*Menu
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) OwnerOrErr() (*Restaurant, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// The edge owner was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: restaurant.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// CategoryOrErr returns the Category value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) CategoryOrErr() ([]*Category, error) {
	if e.loadedTypes[1] {
		return e.Category, nil
	}
	return nil, &NotLoadedError{edge: "category"}
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e MenuEdges) ImagesOrErr() (*File, error) {
	if e.loadedTypes[2] {
		if e.Images == nil {
			// The edge images was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: file.Label}
		}
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// OptionsOrErr returns the Options value or an error if the edge
// was not loaded in eager-loading.
func (e MenuEdges) OptionsOrErr() ([]*Menu, error) {
	if e.loadedTypes[3] {
		return e.Options, nil
	}
	return nil, &NotLoadedError{edge: "options"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Menu) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // description
		&sql.NullBool{},   // isOption
		&[]byte{},         // price
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Menu) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // menu_owner
		&sql.NullInt64{}, // menu_images
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Menu fields.
func (m *Menu) assignValues(values ...interface{}) error {
	if m, n := len(values), len(menu.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	m.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		m.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field description", values[1])
	} else if value.Valid {
		m.Description = value.String
	}
	if value, ok := values[2].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field isOption", values[2])
	} else if value.Valid {
		m.IsOption = value.Bool
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field price", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &m.Price); err != nil {
			return fmt.Errorf("unmarshal field price: %v", err)
		}
	}
	values = values[4:]
	if len(values) == len(menu.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field menu_owner", value)
		} else if value.Valid {
			m.menu_owner = new(int)
			*m.menu_owner = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field menu_images", value)
		} else if value.Valid {
			m.menu_images = new(int)
			*m.menu_images = int(value.Int64)
		}
	}
	return nil
}

// QueryOwner queries the owner edge of the Menu.
func (m *Menu) QueryOwner() *RestaurantQuery {
	return (&MenuClient{config: m.config}).QueryOwner(m)
}

// QueryCategory queries the category edge of the Menu.
func (m *Menu) QueryCategory() *CategoryQuery {
	return (&MenuClient{config: m.config}).QueryCategory(m)
}

// QueryImages queries the images edge of the Menu.
func (m *Menu) QueryImages() *FileQuery {
	return (&MenuClient{config: m.config}).QueryImages(m)
}

// QueryOptions queries the options edge of the Menu.
func (m *Menu) QueryOptions() *MenuQuery {
	return (&MenuClient{config: m.config}).QueryOptions(m)
}

// Update returns a builder for updating this Menu.
// Note that, you need to call Menu.Unwrap() before calling this method, if this Menu
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Menu) Update() *MenuUpdateOne {
	return (&MenuClient{config: m.config}).UpdateOne(m)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (m *Menu) Unwrap() *Menu {
	tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Menu is not a transactional entity")
	}
	m.config.driver = tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Menu) String() string {
	var builder strings.Builder
	builder.WriteString("Menu(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteString(", name=")
	builder.WriteString(m.Name)
	builder.WriteString(", description=")
	builder.WriteString(m.Description)
	builder.WriteString(", isOption=")
	builder.WriteString(fmt.Sprintf("%v", m.IsOption))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", m.Price))
	builder.WriteByte(')')
	return builder.String()
}

// Menus is a parsable slice of Menu.
type Menus []*Menu

func (m Menus) config(cfg config) {
	for _i := range m {
		m[_i].config = cfg
	}
}
