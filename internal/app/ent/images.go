// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"products-service/internal/app/ent/imagefolderpath"
	"products-service/internal/app/ent/images"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Images is the model entity for the Images schema.
type Images struct {
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
	// ImageFolderPathID holds the value of the "image_folder_path_id" field.
	ImageFolderPathID *int `json:"image_folder_path_id,omitempty"`
	// Content holds the value of the "content" field.
	Content *string `json:"content,omitempty"`
	// Path holds the value of the "path" field.
	Path *string `json:"path,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImagesQuery when eager-loading is set.
	Edges           ImagesEdges `json:"edges"`
	products_images *int
	selectValues    sql.SelectValues
}

// ImagesEdges holds the relations/edges for other nodes in the graph.
type ImagesEdges struct {
	// ImageFolderPath holds the value of the image_folder_path edge.
	ImageFolderPath *ImageFolderPath `json:"image_folder_path,omitempty"`
	// ProductHasImage holds the value of the product_has_image edge.
	ProductHasImage []*ProductHasImage `json:"product_has_image,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ImageFolderPathOrErr returns the ImageFolderPath value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImagesEdges) ImageFolderPathOrErr() (*ImageFolderPath, error) {
	if e.ImageFolderPath != nil {
		return e.ImageFolderPath, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: imagefolderpath.Label}
	}
	return nil, &NotLoadedError{edge: "image_folder_path"}
}

// ProductHasImageOrErr returns the ProductHasImage value or an error if the edge
// was not loaded in eager-loading.
func (e ImagesEdges) ProductHasImageOrErr() ([]*ProductHasImage, error) {
	if e.loadedTypes[1] {
		return e.ProductHasImage, nil
	}
	return nil, &NotLoadedError{edge: "product_has_image"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Images) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case images.FieldID, images.FieldCreatedBy, images.FieldUpdatedBy, images.FieldDeletedBy, images.FieldImageFolderPathID:
			values[i] = new(sql.NullInt64)
		case images.FieldContent, images.FieldPath:
			values[i] = new(sql.NullString)
		case images.FieldCreatedAt, images.FieldUpdatedAt, images.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case images.ForeignKeys[0]: // products_images
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Images fields.
func (i *Images) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case images.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case images.FieldCreatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[j])
			} else if value.Valid {
				i.CreatedAt = value.Time
			}
		case images.FieldUpdatedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[j])
			} else if value.Valid {
				i.UpdatedAt = value.Time
			}
		case images.FieldDeletedAt:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[j])
			} else if value.Valid {
				i.DeletedAt = new(time.Time)
				*i.DeletedAt = value.Time
			}
		case images.FieldCreatedBy:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[j])
			} else if value.Valid {
				i.CreatedBy = int(value.Int64)
			}
		case images.FieldUpdatedBy:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[j])
			} else if value.Valid {
				i.UpdatedBy = int(value.Int64)
			}
		case images.FieldDeletedBy:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[j])
			} else if value.Valid {
				i.DeletedBy = new(int)
				*i.DeletedBy = int(value.Int64)
			}
		case images.FieldImageFolderPathID:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field image_folder_path_id", values[j])
			} else if value.Valid {
				i.ImageFolderPathID = new(int)
				*i.ImageFolderPathID = int(value.Int64)
			}
		case images.FieldContent:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[j])
			} else if value.Valid {
				i.Content = new(string)
				*i.Content = value.String
			}
		case images.FieldPath:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[j])
			} else if value.Valid {
				i.Path = new(string)
				*i.Path = value.String
			}
		case images.ForeignKeys[0]:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field products_images", value)
			} else if value.Valid {
				i.products_images = new(int)
				*i.products_images = int(value.Int64)
			}
		default:
			i.selectValues.Set(columns[j], values[j])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Images.
// This includes values selected through modifiers, order, etc.
func (i *Images) Value(name string) (ent.Value, error) {
	return i.selectValues.Get(name)
}

// QueryImageFolderPath queries the "image_folder_path" edge of the Images entity.
func (i *Images) QueryImageFolderPath() *ImageFolderPathQuery {
	return NewImagesClient(i.config).QueryImageFolderPath(i)
}

// QueryProductHasImage queries the "product_has_image" edge of the Images entity.
func (i *Images) QueryProductHasImage() *ProductHasImageQuery {
	return NewImagesClient(i.config).QueryProductHasImage(i)
}

// Update returns a builder for updating this Images.
// Note that you need to call Images.Unwrap() before calling this method if this Images
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Images) Update() *ImagesUpdateOne {
	return NewImagesClient(i.config).UpdateOne(i)
}

// Unwrap unwraps the Images entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Images) Unwrap() *Images {
	_tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Images is not a transactional entity")
	}
	i.config.driver = _tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Images) String() string {
	var builder strings.Builder
	builder.WriteString("Images(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("created_at=")
	builder.WriteString(i.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(i.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := i.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", i.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", i.UpdatedBy))
	builder.WriteString(", ")
	if v := i.DeletedBy; v != nil {
		builder.WriteString("deleted_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := i.ImageFolderPathID; v != nil {
		builder.WriteString("image_folder_path_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := i.Content; v != nil {
		builder.WriteString("content=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := i.Path; v != nil {
		builder.WriteString("path=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// ImagesSlice is a parsable slice of Images.
type ImagesSlice []*Images
