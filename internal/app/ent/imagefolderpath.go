// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"products-service/internal/app/ent/imagefolderpath"
	"products-service/internal/app/ent/imagefoldersource"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// ImageFolderPath is the model entity for the ImageFolderPath schema.
type ImageFolderPath struct {
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
	// ImageFolderSourceID holds the value of the "image_folder_source_id" field.
	ImageFolderSourceID *int `json:"image_folder_source_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ImageFolderPathQuery when eager-loading is set.
	Edges        ImageFolderPathEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ImageFolderPathEdges holds the relations/edges for other nodes in the graph.
type ImageFolderPathEdges struct {
	// ImageFolderSource holds the value of the image_folder_source edge.
	ImageFolderSource *ImageFolderSource `json:"image_folder_source,omitempty"`
	// Images holds the value of the images edge.
	Images []*Images `json:"images,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ImageFolderSourceOrErr returns the ImageFolderSource value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ImageFolderPathEdges) ImageFolderSourceOrErr() (*ImageFolderSource, error) {
	if e.ImageFolderSource != nil {
		return e.ImageFolderSource, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: imagefoldersource.Label}
	}
	return nil, &NotLoadedError{edge: "image_folder_source"}
}

// ImagesOrErr returns the Images value or an error if the edge
// was not loaded in eager-loading.
func (e ImageFolderPathEdges) ImagesOrErr() ([]*Images, error) {
	if e.loadedTypes[1] {
		return e.Images, nil
	}
	return nil, &NotLoadedError{edge: "images"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ImageFolderPath) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case imagefolderpath.FieldID, imagefolderpath.FieldCreatedBy, imagefolderpath.FieldUpdatedBy, imagefolderpath.FieldDeletedBy, imagefolderpath.FieldImageFolderSourceID:
			values[i] = new(sql.NullInt64)
		case imagefolderpath.FieldCreatedAt, imagefolderpath.FieldUpdatedAt, imagefolderpath.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ImageFolderPath fields.
func (ifp *ImageFolderPath) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case imagefolderpath.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ifp.ID = int(value.Int64)
		case imagefolderpath.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ifp.CreatedAt = value.Time
			}
		case imagefolderpath.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				ifp.UpdatedAt = value.Time
			}
		case imagefolderpath.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ifp.DeletedAt = new(time.Time)
				*ifp.DeletedAt = value.Time
			}
		case imagefolderpath.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				ifp.CreatedBy = int(value.Int64)
			}
		case imagefolderpath.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				ifp.UpdatedBy = int(value.Int64)
			}
		case imagefolderpath.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				ifp.DeletedBy = new(int)
				*ifp.DeletedBy = int(value.Int64)
			}
		case imagefolderpath.FieldImageFolderSourceID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field image_folder_source_id", values[i])
			} else if value.Valid {
				ifp.ImageFolderSourceID = new(int)
				*ifp.ImageFolderSourceID = int(value.Int64)
			}
		default:
			ifp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the ImageFolderPath.
// This includes values selected through modifiers, order, etc.
func (ifp *ImageFolderPath) Value(name string) (ent.Value, error) {
	return ifp.selectValues.Get(name)
}

// QueryImageFolderSource queries the "image_folder_source" edge of the ImageFolderPath entity.
func (ifp *ImageFolderPath) QueryImageFolderSource() *ImageFolderSourceQuery {
	return NewImageFolderPathClient(ifp.config).QueryImageFolderSource(ifp)
}

// QueryImages queries the "images" edge of the ImageFolderPath entity.
func (ifp *ImageFolderPath) QueryImages() *ImagesQuery {
	return NewImageFolderPathClient(ifp.config).QueryImages(ifp)
}

// Update returns a builder for updating this ImageFolderPath.
// Note that you need to call ImageFolderPath.Unwrap() before calling this method if this ImageFolderPath
// was returned from a transaction, and the transaction was committed or rolled back.
func (ifp *ImageFolderPath) Update() *ImageFolderPathUpdateOne {
	return NewImageFolderPathClient(ifp.config).UpdateOne(ifp)
}

// Unwrap unwraps the ImageFolderPath entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ifp *ImageFolderPath) Unwrap() *ImageFolderPath {
	_tx, ok := ifp.config.driver.(*txDriver)
	if !ok {
		panic("ent: ImageFolderPath is not a transactional entity")
	}
	ifp.config.driver = _tx.drv
	return ifp
}

// String implements the fmt.Stringer.
func (ifp *ImageFolderPath) String() string {
	var builder strings.Builder
	builder.WriteString("ImageFolderPath(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ifp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(ifp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(ifp.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := ifp.DeletedAt; v != nil {
		builder.WriteString("deleted_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", ifp.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", ifp.UpdatedBy))
	builder.WriteString(", ")
	if v := ifp.DeletedBy; v != nil {
		builder.WriteString("deleted_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ifp.ImageFolderSourceID; v != nil {
		builder.WriteString("image_folder_source_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// ImageFolderPaths is a parsable slice of ImageFolderPath.
type ImageFolderPaths []*ImageFolderPath
