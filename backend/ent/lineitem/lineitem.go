// Code generated by entc, DO NOT EDIT.

package lineitem

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the lineitem type in the database.
	Label = "line_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"

	// EdgeItem holds the string denoting the item edge name in mutations.
	EdgeItem = "item"
	// EdgeCart holds the string denoting the cart edge name in mutations.
	EdgeCart = "cart"

	// Table holds the table name of the lineitem in the database.
	Table = "line_items"
	// ItemTable is the table the holds the item relation/edge.
	ItemTable = "line_items"
	// ItemInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemInverseTable = "items"
	// ItemColumn is the table column denoting the item relation/edge.
	ItemColumn = "line_item_item"
	// CartTable is the table the holds the cart relation/edge.
	CartTable = "line_items"
	// CartInverseTable is the table name for the Cart entity.
	// It exists in this package in order to avoid circular dependency with the "cart" package.
	CartInverseTable = "carts"
	// CartColumn is the table column denoting the cart relation/edge.
	CartColumn = "cart_line_items"
)

// Columns holds all SQL columns for lineitem fields.
var Columns = []string{
	FieldID,
	FieldQuantity,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the LineItem type.
var ForeignKeys = []string{
	"cart_line_items",
	"line_item_item",
}

var (
	// QuantityValidator is a validator for the "quantity" field. It is called by the builders before save.
	QuantityValidator func(int) error
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)