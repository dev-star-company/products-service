// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"products-service/internal/app/ent/brand"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Brand is the model entity for the Brand schema.
type Brand struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy int `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy int `json:"updated_by,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy *int `json:"deleted_by,omitempty"`
	// Name holds the value of the "name" field.
	Name *string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BrandQuery when eager-loading is set.
	Edges        BrandEdges `json:"edges"`
	selectValues sql.SelectValues
}

// BrandEdges holds the relations/edges for other nodes in the graph.
type BrandEdges struct {
	// Products holds the value of the products edge.
	Products []*Products `json:"products,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProductsOrErr returns the Products value or an error if the edge
// was not loaded in eager-loading.
func (e BrandEdges) ProductsOrErr() ([]*Products, error) {
	if e.loadedTypes[0] {
		return e.Products, nil
	}
	return nil, &NotLoadedError{edge: "products"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Brand) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case brand.FieldID, brand.FieldCreatedBy, brand.FieldUpdatedBy, brand.FieldDeletedBy:
			values[i] = new(sql.NullInt64)
		case brand.FieldName:
			values[i] = new(sql.NullString)
		case brand.FieldCreatedAt, brand.FieldUpdatedAt, brand.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Brand fields.
func (b *Brand) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case brand.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case brand.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case brand.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case brand.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				b.DeletedAt = new(time.Time)
				*b.DeletedAt = value.Time
			}
		case brand.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				b.CreatedBy = int(value.Int64)
			}
		case brand.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				b.UpdatedBy = int(value.Int64)
			}
		case brand.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				b.DeletedBy = new(int)
				*b.DeletedBy = int(value.Int64)
			}
		case brand.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				b.Name = new(string)
				*b.Name = value.String
			}
		default:
			b.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Brand.
// This includes values selected through modifiers, order, etc.
func (b *Brand) Value(name string) (ent.Value, error) {
	return b.selectValues.Get(name)
}

// QueryProducts queries the "products" edge of the Brand entity.
func (b *Brand) QueryProducts() *ProductsQuery {
	return NewBrandClient(b.config).QueryProducts(b)
}

// Update returns a builder for updating this Brand.
// Note that you need to call Brand.Unwrap() before calling this method if this Brand
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Brand) Update() *BrandUpdateOne {
	return NewBrandClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Brand entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Brand) Unwrap() *Brand {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Brand is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Brand) String() string {
	var builder strings.Builder
	builder.WriteString("Brand(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("created_at=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := b.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", b.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", b.UpdatedBy))
	builder.WriteString(", ")
	if v := b.DeletedBy; v != nil {
		builder.WriteString("deleted_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := b.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Brands is a parsable slice of Brand.
type Brands []*Brand
