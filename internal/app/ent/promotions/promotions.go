// Code generated by ent, DO NOT EDIT.

package promotions

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the promotions type in the database.
	Label = "promotions"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldStartingDatetime holds the string denoting the starting_datetime field in the database.
	FieldStartingDatetime = "starting_datetime"
	// FieldEndingDatetime holds the string denoting the ending_datetime field in the database.
	FieldEndingDatetime = "ending_datetime"
	// EdgePromotionHasProduct holds the string denoting the promotion_has_product edge name in mutations.
	EdgePromotionHasProduct = "promotion_has_product"
	// Table holds the table name of the promotions in the database.
	Table = "promotions"
	// PromotionHasProductTable is the table that holds the promotion_has_product relation/edge.
	PromotionHasProductTable = "promotion_has_products"
	// PromotionHasProductInverseTable is the table name for the PromotionHasProduct entity.
	// It exists in this package in order to avoid circular dependency with the "promotionhasproduct" package.
	PromotionHasProductInverseTable = "promotion_has_products"
	// PromotionHasProductColumn is the table column denoting the promotion_has_product relation/edge.
	PromotionHasProductColumn = "promotions_promotion_has_product"
)

// Columns holds all SQL columns for promotions fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldCreatedBy,
	FieldUpdatedBy,
	FieldDeletedBy,
	FieldName,
	FieldStartingDatetime,
	FieldEndingDatetime,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// CreatedByValidator is a validator for the "created_by" field. It is called by the builders before save.
	CreatedByValidator func(int) error
	// UpdatedByValidator is a validator for the "updated_by" field. It is called by the builders before save.
	UpdatedByValidator func(int) error
)

// OrderOption defines the ordering options for the Promotions queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByStartingDatetime orders the results by the starting_datetime field.
func ByStartingDatetime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStartingDatetime, opts...).ToFunc()
}

// ByEndingDatetime orders the results by the ending_datetime field.
func ByEndingDatetime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEndingDatetime, opts...).ToFunc()
}

// ByPromotionHasProductCount orders the results by promotion_has_product count.
func ByPromotionHasProductCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPromotionHasProductStep(), opts...)
	}
}

// ByPromotionHasProduct orders the results by promotion_has_product terms.
func ByPromotionHasProduct(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPromotionHasProductStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newPromotionHasProductStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PromotionHasProductInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, PromotionHasProductTable, PromotionHasProductColumn),
	)
}
