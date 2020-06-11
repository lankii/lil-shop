// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/awolk/lil-shop/backend/ent/cart"
	"github.com/awolk/lil-shop/backend/ent/item"
	"github.com/awolk/lil-shop/backend/ent/lineitem"
	"github.com/google/uuid"

	"github.com/facebookincubator/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeCart     = "Cart"
	TypeItem     = "Item"
	TypeLineItem = "LineItem"
)

// CartMutation represents an operation that mutate the Carts
// nodes in the graph.
type CartMutation struct {
	config
	op                Op
	typ               string
	id                *uuid.UUID
	clearedFields     map[string]struct{}
	line_items        map[uuid.UUID]struct{}
	removedline_items map[uuid.UUID]struct{}
	done              bool
	oldValue          func(context.Context) (*Cart, error)
}

var _ ent.Mutation = (*CartMutation)(nil)

// cartOption allows to manage the mutation configuration using functional options.
type cartOption func(*CartMutation)

// newCartMutation creates new mutation for $n.Name.
func newCartMutation(c config, op Op, opts ...cartOption) *CartMutation {
	m := &CartMutation{
		config:        c,
		op:            op,
		typ:           TypeCart,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withCartID sets the id field of the mutation.
func withCartID(id uuid.UUID) cartOption {
	return func(m *CartMutation) {
		var (
			err   error
			once  sync.Once
			value *Cart
		)
		m.oldValue = func(ctx context.Context) (*Cart, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Cart.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withCart sets the old Cart of the mutation.
func withCart(node *Cart) cartOption {
	return func(m *CartMutation) {
		m.oldValue = func(context.Context) (*Cart, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m CartMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m CartMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Cart creation.
func (m *CartMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *CartMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// AddLineItemIDs adds the line_items edge to LineItem by ids.
func (m *CartMutation) AddLineItemIDs(ids ...uuid.UUID) {
	if m.line_items == nil {
		m.line_items = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.line_items[ids[i]] = struct{}{}
	}
}

// RemoveLineItemIDs removes the line_items edge to LineItem by ids.
func (m *CartMutation) RemoveLineItemIDs(ids ...uuid.UUID) {
	if m.removedline_items == nil {
		m.removedline_items = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.removedline_items[ids[i]] = struct{}{}
	}
}

// RemovedLineItems returns the removed ids of line_items.
func (m *CartMutation) RemovedLineItemsIDs() (ids []uuid.UUID) {
	for id := range m.removedline_items {
		ids = append(ids, id)
	}
	return
}

// LineItemsIDs returns the line_items ids in the mutation.
func (m *CartMutation) LineItemsIDs() (ids []uuid.UUID) {
	for id := range m.line_items {
		ids = append(ids, id)
	}
	return
}

// ResetLineItems reset all changes of the "line_items" edge.
func (m *CartMutation) ResetLineItems() {
	m.line_items = nil
	m.removedline_items = nil
}

// Op returns the operation name.
func (m *CartMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Cart).
func (m *CartMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *CartMutation) Fields() []string {
	fields := make([]string, 0, 0)
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *CartMutation) Field(name string) (ent.Value, bool) {
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *CartMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	return nil, fmt.Errorf("unknown Cart field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CartMutation) SetField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Cart field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *CartMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *CartMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *CartMutation) AddField(name string, value ent.Value) error {
	return fmt.Errorf("unknown Cart numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *CartMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *CartMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *CartMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Cart nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *CartMutation) ResetField(name string) error {
	return fmt.Errorf("unknown Cart field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *CartMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.line_items != nil {
		edges = append(edges, cart.EdgeLineItems)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *CartMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case cart.EdgeLineItems:
		ids := make([]ent.Value, 0, len(m.line_items))
		for id := range m.line_items {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *CartMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedline_items != nil {
		edges = append(edges, cart.EdgeLineItems)
	}
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *CartMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case cart.EdgeLineItems:
		ids := make([]ent.Value, 0, len(m.removedline_items))
		for id := range m.removedline_items {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *CartMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *CartMutation) EdgeCleared(name string) bool {
	switch name {
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *CartMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Cart unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *CartMutation) ResetEdge(name string) error {
	switch name {
	case cart.EdgeLineItems:
		m.ResetLineItems()
		return nil
	}
	return fmt.Errorf("unknown Cart edge %s", name)
}

// ItemMutation represents an operation that mutate the Items
// nodes in the graph.
type ItemMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	name          *string
	cost_cents    *int
	addcost_cents *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Item, error)
}

var _ ent.Mutation = (*ItemMutation)(nil)

// itemOption allows to manage the mutation configuration using functional options.
type itemOption func(*ItemMutation)

// newItemMutation creates new mutation for $n.Name.
func newItemMutation(c config, op Op, opts ...itemOption) *ItemMutation {
	m := &ItemMutation{
		config:        c,
		op:            op,
		typ:           TypeItem,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withItemID sets the id field of the mutation.
func withItemID(id uuid.UUID) itemOption {
	return func(m *ItemMutation) {
		var (
			err   error
			once  sync.Once
			value *Item
		)
		m.oldValue = func(ctx context.Context) (*Item, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Item.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withItem sets the old Item of the mutation.
func withItem(node *Item) itemOption {
	return func(m *ItemMutation) {
		m.oldValue = func(context.Context) (*Item, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ItemMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ItemMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on Item creation.
func (m *ItemMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *ItemMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the name field.
func (m *ItemMutation) SetName(s string) {
	m.name = &s
}

// Name returns the name value in the mutation.
func (m *ItemMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old name value of the Item.
// If the Item object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *ItemMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName reset all changes of the "name" field.
func (m *ItemMutation) ResetName() {
	m.name = nil
}

// SetCostCents sets the cost_cents field.
func (m *ItemMutation) SetCostCents(i int) {
	m.cost_cents = &i
	m.addcost_cents = nil
}

// CostCents returns the cost_cents value in the mutation.
func (m *ItemMutation) CostCents() (r int, exists bool) {
	v := m.cost_cents
	if v == nil {
		return
	}
	return *v, true
}

// OldCostCents returns the old cost_cents value of the Item.
// If the Item object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *ItemMutation) OldCostCents(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCostCents is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCostCents requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCostCents: %w", err)
	}
	return oldValue.CostCents, nil
}

// AddCostCents adds i to cost_cents.
func (m *ItemMutation) AddCostCents(i int) {
	if m.addcost_cents != nil {
		*m.addcost_cents += i
	} else {
		m.addcost_cents = &i
	}
}

// AddedCostCents returns the value that was added to the cost_cents field in this mutation.
func (m *ItemMutation) AddedCostCents() (r int, exists bool) {
	v := m.addcost_cents
	if v == nil {
		return
	}
	return *v, true
}

// ResetCostCents reset all changes of the "cost_cents" field.
func (m *ItemMutation) ResetCostCents() {
	m.cost_cents = nil
	m.addcost_cents = nil
}

// Op returns the operation name.
func (m *ItemMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Item).
func (m *ItemMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *ItemMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, item.FieldName)
	}
	if m.cost_cents != nil {
		fields = append(fields, item.FieldCostCents)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *ItemMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case item.FieldName:
		return m.Name()
	case item.FieldCostCents:
		return m.CostCents()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *ItemMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case item.FieldName:
		return m.OldName(ctx)
	case item.FieldCostCents:
		return m.OldCostCents(ctx)
	}
	return nil, fmt.Errorf("unknown Item field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *ItemMutation) SetField(name string, value ent.Value) error {
	switch name {
	case item.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case item.FieldCostCents:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCostCents(v)
		return nil
	}
	return fmt.Errorf("unknown Item field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *ItemMutation) AddedFields() []string {
	var fields []string
	if m.addcost_cents != nil {
		fields = append(fields, item.FieldCostCents)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *ItemMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case item.FieldCostCents:
		return m.AddedCostCents()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *ItemMutation) AddField(name string, value ent.Value) error {
	switch name {
	case item.FieldCostCents:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCostCents(v)
		return nil
	}
	return fmt.Errorf("unknown Item numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *ItemMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *ItemMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *ItemMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Item nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *ItemMutation) ResetField(name string) error {
	switch name {
	case item.FieldName:
		m.ResetName()
		return nil
	case item.FieldCostCents:
		m.ResetCostCents()
		return nil
	}
	return fmt.Errorf("unknown Item field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *ItemMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *ItemMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *ItemMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *ItemMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *ItemMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *ItemMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *ItemMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Item unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *ItemMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Item edge %s", name)
}

// LineItemMutation represents an operation that mutate the LineItems
// nodes in the graph.
type LineItemMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	quantity      *int
	addquantity   *int
	clearedFields map[string]struct{}
	item          *uuid.UUID
	cleareditem   bool
	cart          *uuid.UUID
	clearedcart   bool
	done          bool
	oldValue      func(context.Context) (*LineItem, error)
}

var _ ent.Mutation = (*LineItemMutation)(nil)

// lineitemOption allows to manage the mutation configuration using functional options.
type lineitemOption func(*LineItemMutation)

// newLineItemMutation creates new mutation for $n.Name.
func newLineItemMutation(c config, op Op, opts ...lineitemOption) *LineItemMutation {
	m := &LineItemMutation{
		config:        c,
		op:            op,
		typ:           TypeLineItem,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLineItemID sets the id field of the mutation.
func withLineItemID(id uuid.UUID) lineitemOption {
	return func(m *LineItemMutation) {
		var (
			err   error
			once  sync.Once
			value *LineItem
		)
		m.oldValue = func(ctx context.Context) (*LineItem, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().LineItem.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLineItem sets the old LineItem of the mutation.
func withLineItem(node *LineItem) lineitemOption {
	return func(m *LineItemMutation) {
		m.oldValue = func(context.Context) (*LineItem, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LineItemMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LineItemMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that, this
// operation is accepted only on LineItem creation.
func (m *LineItemMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the id value in the mutation. Note that, the id
// is available only if it was provided to the builder.
func (m *LineItemMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetQuantity sets the quantity field.
func (m *LineItemMutation) SetQuantity(i int) {
	m.quantity = &i
	m.addquantity = nil
}

// Quantity returns the quantity value in the mutation.
func (m *LineItemMutation) Quantity() (r int, exists bool) {
	v := m.quantity
	if v == nil {
		return
	}
	return *v, true
}

// OldQuantity returns the old quantity value of the LineItem.
// If the LineItem object wasn't provided to the builder, the object is fetched
// from the database.
// An error is returned if the mutation operation is not UpdateOne, or database query fails.
func (m *LineItemMutation) OldQuantity(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldQuantity is allowed only on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldQuantity requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldQuantity: %w", err)
	}
	return oldValue.Quantity, nil
}

// AddQuantity adds i to quantity.
func (m *LineItemMutation) AddQuantity(i int) {
	if m.addquantity != nil {
		*m.addquantity += i
	} else {
		m.addquantity = &i
	}
}

// AddedQuantity returns the value that was added to the quantity field in this mutation.
func (m *LineItemMutation) AddedQuantity() (r int, exists bool) {
	v := m.addquantity
	if v == nil {
		return
	}
	return *v, true
}

// ResetQuantity reset all changes of the "quantity" field.
func (m *LineItemMutation) ResetQuantity() {
	m.quantity = nil
	m.addquantity = nil
}

// SetItemID sets the item edge to Item by id.
func (m *LineItemMutation) SetItemID(id uuid.UUID) {
	m.item = &id
}

// ClearItem clears the item edge to Item.
func (m *LineItemMutation) ClearItem() {
	m.cleareditem = true
}

// ItemCleared returns if the edge item was cleared.
func (m *LineItemMutation) ItemCleared() bool {
	return m.cleareditem
}

// ItemID returns the item id in the mutation.
func (m *LineItemMutation) ItemID() (id uuid.UUID, exists bool) {
	if m.item != nil {
		return *m.item, true
	}
	return
}

// ItemIDs returns the item ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// ItemID instead. It exists only for internal usage by the builders.
func (m *LineItemMutation) ItemIDs() (ids []uuid.UUID) {
	if id := m.item; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetItem reset all changes of the "item" edge.
func (m *LineItemMutation) ResetItem() {
	m.item = nil
	m.cleareditem = false
}

// SetCartID sets the cart edge to Cart by id.
func (m *LineItemMutation) SetCartID(id uuid.UUID) {
	m.cart = &id
}

// ClearCart clears the cart edge to Cart.
func (m *LineItemMutation) ClearCart() {
	m.clearedcart = true
}

// CartCleared returns if the edge cart was cleared.
func (m *LineItemMutation) CartCleared() bool {
	return m.clearedcart
}

// CartID returns the cart id in the mutation.
func (m *LineItemMutation) CartID() (id uuid.UUID, exists bool) {
	if m.cart != nil {
		return *m.cart, true
	}
	return
}

// CartIDs returns the cart ids in the mutation.
// Note that ids always returns len(ids) <= 1 for unique edges, and you should use
// CartID instead. It exists only for internal usage by the builders.
func (m *LineItemMutation) CartIDs() (ids []uuid.UUID) {
	if id := m.cart; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetCart reset all changes of the "cart" edge.
func (m *LineItemMutation) ResetCart() {
	m.cart = nil
	m.clearedcart = false
}

// Op returns the operation name.
func (m *LineItemMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (LineItem).
func (m *LineItemMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during
// this mutation. Note that, in order to get all numeric
// fields that were in/decremented, call AddedFields().
func (m *LineItemMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.quantity != nil {
		fields = append(fields, lineitem.FieldQuantity)
	}
	return fields
}

// Field returns the value of a field with the given name.
// The second boolean value indicates that this field was
// not set, or was not define in the schema.
func (m *LineItemMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case lineitem.FieldQuantity:
		return m.Quantity()
	}
	return nil, false
}

// OldField returns the old value of the field from the database.
// An error is returned if the mutation operation is not UpdateOne,
// or the query to the database was failed.
func (m *LineItemMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case lineitem.FieldQuantity:
		return m.OldQuantity(ctx)
	}
	return nil, fmt.Errorf("unknown LineItem field %s", name)
}

// SetField sets the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *LineItemMutation) SetField(name string, value ent.Value) error {
	switch name {
	case lineitem.FieldQuantity:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetQuantity(v)
		return nil
	}
	return fmt.Errorf("unknown LineItem field %s", name)
}

// AddedFields returns all numeric fields that were incremented
// or decremented during this mutation.
func (m *LineItemMutation) AddedFields() []string {
	var fields []string
	if m.addquantity != nil {
		fields = append(fields, lineitem.FieldQuantity)
	}
	return fields
}

// AddedField returns the numeric value that was in/decremented
// from a field with the given name. The second value indicates
// that this field was not set, or was not define in the schema.
func (m *LineItemMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case lineitem.FieldQuantity:
		return m.AddedQuantity()
	}
	return nil, false
}

// AddField adds the value for the given name. It returns an
// error if the field is not defined in the schema, or if the
// type mismatch the field type.
func (m *LineItemMutation) AddField(name string, value ent.Value) error {
	switch name {
	case lineitem.FieldQuantity:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddQuantity(v)
		return nil
	}
	return fmt.Errorf("unknown LineItem numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared
// during this mutation.
func (m *LineItemMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicates if this field was
// cleared in this mutation.
func (m *LineItemMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value for the given name. It returns an
// error if the field is not defined in the schema.
func (m *LineItemMutation) ClearField(name string) error {
	return fmt.Errorf("unknown LineItem nullable field %s", name)
}

// ResetField resets all changes in the mutation regarding the
// given field name. It returns an error if the field is not
// defined in the schema.
func (m *LineItemMutation) ResetField(name string) error {
	switch name {
	case lineitem.FieldQuantity:
		m.ResetQuantity()
		return nil
	}
	return fmt.Errorf("unknown LineItem field %s", name)
}

// AddedEdges returns all edge names that were set/added in this
// mutation.
func (m *LineItemMutation) AddedEdges() []string {
	edges := make([]string, 0, 2)
	if m.item != nil {
		edges = append(edges, lineitem.EdgeItem)
	}
	if m.cart != nil {
		edges = append(edges, lineitem.EdgeCart)
	}
	return edges
}

// AddedIDs returns all ids (to other nodes) that were added for
// the given edge name.
func (m *LineItemMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case lineitem.EdgeItem:
		if id := m.item; id != nil {
			return []ent.Value{*id}
		}
	case lineitem.EdgeCart:
		if id := m.cart; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this
// mutation.
func (m *LineItemMutation) RemovedEdges() []string {
	edges := make([]string, 0, 2)
	return edges
}

// RemovedIDs returns all ids (to other nodes) that were removed for
// the given edge name.
func (m *LineItemMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this
// mutation.
func (m *LineItemMutation) ClearedEdges() []string {
	edges := make([]string, 0, 2)
	if m.cleareditem {
		edges = append(edges, lineitem.EdgeItem)
	}
	if m.clearedcart {
		edges = append(edges, lineitem.EdgeCart)
	}
	return edges
}

// EdgeCleared returns a boolean indicates if this edge was
// cleared in this mutation.
func (m *LineItemMutation) EdgeCleared(name string) bool {
	switch name {
	case lineitem.EdgeItem:
		return m.cleareditem
	case lineitem.EdgeCart:
		return m.clearedcart
	}
	return false
}

// ClearEdge clears the value for the given name. It returns an
// error if the edge name is not defined in the schema.
func (m *LineItemMutation) ClearEdge(name string) error {
	switch name {
	case lineitem.EdgeItem:
		m.ClearItem()
		return nil
	case lineitem.EdgeCart:
		m.ClearCart()
		return nil
	}
	return fmt.Errorf("unknown LineItem unique edge %s", name)
}

// ResetEdge resets all changes in the mutation regarding the
// given edge name. It returns an error if the edge is not
// defined in the schema.
func (m *LineItemMutation) ResetEdge(name string) error {
	switch name {
	case lineitem.EdgeItem:
		m.ResetItem()
		return nil
	case lineitem.EdgeCart:
		m.ResetCart()
		return nil
	}
	return fmt.Errorf("unknown LineItem edge %s", name)
}