// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"products-service/internal/app/ent/producthasproductreference"
	"products-service/internal/app/ent/productreferences"
	"products-service/internal/app/ent/products"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ProductHasProductReference is the model entity for the ProductHasProductReference schema.
type ProductHasProductReference struct {
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
	// ProductReferenceID holds the value of the "product_reference_id" field.
	ProductReferenceID *int `json:"product_reference_id,omitempty"`
	// ProductID holds the value of the "product_id" field.
	ProductID *int `json:"product_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProductHasProductReferenceQuery when eager-loading is set.
	Edges        ProductHasProductReferenceEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProductHasProductReferenceEdges holds the relations/edges for other nodes in the graph.
type ProductHasProductReferenceEdges struct {
	// ProductReference holds the value of the product_reference edge.
	ProductReference *ProductReferences `json:"product_reference,omitempty"`
	// Product holds the value of the product edge.
	Product *Products `json:"product,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProductReferenceOrErr returns the ProductReference value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductHasProductReferenceEdges) ProductReferenceOrErr() (*ProductReferences, error) {
	if e.ProductReference != nil {
		return e.ProductReference, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: productreferences.Label}
	}
	return nil, &NotLoadedError{edge: "product_reference"}
}

// ProductOrErr returns the Product value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProductHasProductReferenceEdges) ProductOrErr() (*Products, error) {
	if e.Product != nil {
		return e.Product, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: products.Label}
	}
	return nil, &NotLoadedError{edge: "product"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ProductHasProductReference) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case producthasproductreference.FieldID, producthasproductreference.FieldCreatedBy, producthasproductreference.FieldUpdatedBy, producthasproductreference.FieldDeletedBy, producthasproductreference.FieldProductReferenceID, producthasproductreference.FieldProductID:
			values[i] = new(sql.NullInt64)
		case producthasproductreference.FieldCreatedAt, producthasproductreference.FieldUpdatedAt, producthasproductreference.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ProductHasProductReference fields.
func (phpr *ProductHasProductReference) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case producthasproductreference.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			phpr.ID = int(value.Int64)
		case producthasproductreference.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				phpr.CreatedAt = value.Time
			}
		case producthasproductreference.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				phpr.UpdatedAt = value.Time
			}
		case producthasproductreference.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				phpr.DeletedAt = new(time.Time)
				*phpr.DeletedAt = value.Time
			}
		case producthasproductreference.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				phpr.CreatedBy = int(value.Int64)
			}
		case producthasproductreference.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				phpr.UpdatedBy = int(value.Int64)
			}
		case producthasproductreference.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				phpr.DeletedBy = new(int)
				*phpr.DeletedBy = int(value.Int64)
			}
		case producthasproductreference.FieldProductReferenceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_reference_id", values[i])
			} else if value.Valid {
				phpr.ProductReferenceID = new(int)
				*phpr.ProductReferenceID = int(value.Int64)
			}
		case producthasproductreference.FieldProductID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field product_id", values[i])
			} else if value.Valid {
				phpr.ProductID = new(int)
				*phpr.ProductID = int(value.Int64)
			}
		default:
			phpr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ProductHasProductReference.
// This includes values selected through modifiers, order, etc.
func (phpr *ProductHasProductReference) Value(name string) (ent.Value, error) {
	return phpr.selectValues.Get(name)
}

// QueryProductReference queries the "product_reference" edge of the ProductHasProductReference entity.
func (phpr *ProductHasProductReference) QueryProductReference() *ProductReferencesQuery {
	return NewProductHasProductReferenceClient(phpr.config).QueryProductReference(phpr)
}

// QueryProduct queries the "product" edge of the ProductHasProductReference entity.
func (phpr *ProductHasProductReference) QueryProduct() *ProductsQuery {
	return NewProductHasProductReferenceClient(phpr.config).QueryProduct(phpr)
}

// Update returns a builder for updating this ProductHasProductReference.
// Note that you need to call ProductHasProductReference.Unwrap() before calling this method if this ProductHasProductReference
// was returned from a transaction, and the transaction was committed or rolled back.
func (phpr *ProductHasProductReference) Update() *ProductHasProductReferenceUpdateOne {
	return NewProductHasProductReferenceClient(phpr.config).UpdateOne(phpr)
}

// Unwrap unwraps the ProductHasProductReference entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (phpr *ProductHasProductReference) Unwrap() *ProductHasProductReference {
	_tx, ok := phpr.config.driver.(*txDriver)
	if !ok {
		panic("ent: ProductHasProductReference is not a transactional entity")
	}
	phpr.config.driver = _tx.drv
	return phpr
}

// String implements the fmt.Stringer.
func (phpr *ProductHasProductReference) String() string {
	var builder strings.Builder
	builder.WriteString("ProductHasProductReference(")
	builder.WriteString(fmt.Sprintf("id=%v, ", phpr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(phpr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(phpr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := phpr.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", phpr.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", phpr.UpdatedBy))
	builder.WriteString(", ")
	if v := phpr.DeletedBy; v != nil {
		builder.WriteString("deleted_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := phpr.ProductReferenceID; v != nil {
		builder.WriteString("product_reference_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := phpr.ProductID; v != nil {
		builder.WriteString("product_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// ProductHasProductReferences is a parsable slice of ProductHasProductReference.
type ProductHasProductReferences []*ProductHasProductReference
