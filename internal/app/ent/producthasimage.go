// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"products-service/internal/app/ent/images"
	"products-service/internal/app/ent/producthasimage"
	"products-service/internal/app/ent/products"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ProductHasImage is the model entity for the ProductHasImage schema.
type ProductHasImage struct {
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
	// ImageID holds the value of the "image_id" field.
	ImageID *int `json:"image_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID *int `json:"product_id,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority int `json:"priority,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductHasImageQuery when eager-loading is set.
	Edges        ProductHasImageEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductHasImageEdges holds the relations/edges for other nodes in the graph.
type ProductHasImageEdges struct {
	// Product holds the value of the product edge.
	Product *Products `json:"product,omitempty"`
	// Image holds the value of the image edge.
	Image *Images `json:"image,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductHasImageEdges) ProductOrErr() (*Products, error) {
	if e.Product != nil {
		return e.Product, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: products.Label}
	}
	return nil, &NotLoadedError{edge: "product"}
}

// ImageOrErr returns the Image value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductHasImageEdges) ImageOrErr() (*Images, error) {
	if e.Image != nil {
		return e.Image, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: images.Label}
	}
	return nil, &NotLoadedError{edge: "image"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductHasImage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case producthasimage.FieldID, producthasimage.FieldCreatedBy, producthasimage.FieldUpdatedBy, producthasimage.FieldDeletedBy, producthasimage.FieldImageID, producthasimage.FieldProductID, producthasimage.FieldPriority:
			values[i] = new(sql.NullInt64)
		case producthasimage.FieldCreatedAt, producthasimage.FieldUpdatedAt, producthasimage.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductHasImage fields.
func (phi *ProductHasImage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case producthasimage.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			phi.ID = int(value.Int64)
		case producthasimage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				phi.CreatedAt = value.Time
			}
		case producthasimage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				phi.UpdatedAt = value.Time
			}
		case producthasimage.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				phi.DeletedAt = new(time.Time)
				*phi.DeletedAt = value.Time
			}
		case producthasimage.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				phi.CreatedBy = int(value.Int64)
			}
		case producthasimage.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				phi.UpdatedBy = int(value.Int64)
			}
		case producthasimage.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				phi.DeletedBy = new(int)
				*phi.DeletedBy = int(value.Int64)
			}
		case producthasimage.FieldImageID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field image_id", values[i])
			} else if value.Valid {
				phi.ImageID = new(int)
				*phi.ImageID = int(value.Int64)
			}
		case producthasimage.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				phi.ProductID = new(int)
				*phi.ProductID = int(value.Int64)
			}
		case producthasimage.FieldPriority:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value.Valid {
				phi.Priority = int(value.Int64)
			}
		default:
			phi.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductHasImage.
// This includes values selected through modifiers, order, etc.
func (phi *ProductHasImage) Value(name string) (ent.Value, error) {
	return phi.selectValues.Get(name)
}

// QueryProduct queries the "product" edge of the ProductHasImage entity.
func (phi *ProductHasImage) QueryProduct() *ProductsQuery {
	return NewProductHasImageClient(phi.config).QueryProduct(phi)
}

// QueryImage queries the "image" edge of the ProductHasImage entity.
func (phi *ProductHasImage) QueryImage() *ImagesQuery {
	return NewProductHasImageClient(phi.config).QueryImage(phi)
}

// Update returns a builder for updating this ProductHasImage.
// Note that you need to call ProductHasImage.Unwrap() before calling this method if this ProductHasImage
// was returned from a transaction, and the transaction was committed or rolled back.
func (phi *ProductHasImage) Update() *ProductHasImageUpdateOne {
	return NewProductHasImageClient(phi.config).UpdateOne(phi)
}

// Unwrap unwraps the ProductHasImage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (phi *ProductHasImage) Unwrap() *ProductHasImage {
	_tx, ok := phi.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductHasImage is not a transactional entity")
	}
	phi.config.driver = _tx.drv
	return phi
}

// String implements the fmt.Stringer.
func (phi *ProductHasImage) String() string {
	var builder strings.Builder
	builder.WriteString("ProductHasImage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", phi.ID))
	builder.WriteString("created_at=")
	builder.WriteString(phi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(phi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := phi.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", phi.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", phi.UpdatedBy))
	builder.WriteString(", ")
	if v := phi.DeletedBy; v != nil {
		builder.WriteString("deleted_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := phi.ImageID; v != nil {
		builder.WriteString("image_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := phi.ProductID; v != nil {
		builder.WriteString("product_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", phi.Priority))
	builder.WriteByte(')')
	return builder.String()
}

// ProductHasImages is a parsable slice of ProductHasImage.
type ProductHasImages []*ProductHasImage
