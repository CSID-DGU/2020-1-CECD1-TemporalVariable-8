// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lightServer/ent/order"
	"lightServer/ent/orderfield"
	"lightServer/ent/restaurant"
	"lightServer/ent/user"
	"time"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetOrderAt sets the order_at field.
func (oc *OrderCreate) SetOrderAt(t time.Time) *OrderCreate {
	oc.mutation.SetOrderAt(t)
	return oc
}

// SetNillableOrderAt sets the order_at field if the given value is not nil.
func (oc *OrderCreate) SetNillableOrderAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetOrderAt(*t)
	}
	return oc
}

// SetCookingAt sets the cooking_at field.
func (oc *OrderCreate) SetCookingAt(t time.Time) *OrderCreate {
	oc.mutation.SetCookingAt(t)
	return oc
}

// SetNillableCookingAt sets the cooking_at field if the given value is not nil.
func (oc *OrderCreate) SetNillableCookingAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCookingAt(*t)
	}
	return oc
}

// SetDeliverAt sets the deliver_at field.
func (oc *OrderCreate) SetDeliverAt(t time.Time) *OrderCreate {
	oc.mutation.SetDeliverAt(t)
	return oc
}

// SetNillableDeliverAt sets the deliver_at field if the given value is not nil.
func (oc *OrderCreate) SetNillableDeliverAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetDeliverAt(*t)
	}
	return oc
}

// SetCompleteAt sets the complete_at field.
func (oc *OrderCreate) SetCompleteAt(t time.Time) *OrderCreate {
	oc.mutation.SetCompleteAt(t)
	return oc
}

// SetNillableCompleteAt sets the complete_at field if the given value is not nil.
func (oc *OrderCreate) SetNillableCompleteAt(t *time.Time) *OrderCreate {
	if t != nil {
		oc.SetCompleteAt(*t)
	}
	return oc
}

// SetWhoID sets the who edge to User by id.
func (oc *OrderCreate) SetWhoID(id int) *OrderCreate {
	oc.mutation.SetWhoID(id)
	return oc
}

// SetWho sets the who edge to User.
func (oc *OrderCreate) SetWho(u *User) *OrderCreate {
	return oc.SetWhoID(u.ID)
}

// SetWhereID sets the where edge to Restaurant by id.
func (oc *OrderCreate) SetWhereID(id int) *OrderCreate {
	oc.mutation.SetWhereID(id)
	return oc
}

// SetWhere sets the where edge to Restaurant.
func (oc *OrderCreate) SetWhere(r *Restaurant) *OrderCreate {
	return oc.SetWhereID(r.ID)
}

// AddItemIDs adds the items edge to OrderField by ids.
func (oc *OrderCreate) AddItemIDs(ids ...int) *OrderCreate {
	oc.mutation.AddItemIDs(ids...)
	return oc
}

// AddItems adds the items edges to OrderField.
func (oc *OrderCreate) AddItems(o ...*OrderField) *OrderCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddItemIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	var (
		err  error
		node *Order
	)
	oc.defaults()
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			node, err = oc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (oc *OrderCreate) defaults() {
	if _, ok := oc.mutation.OrderAt(); !ok {
		v := order.DefaultOrderAt()
		oc.mutation.SetOrderAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrderCreate) check() error {
	if _, ok := oc.mutation.OrderAt(); !ok {
		return &ValidationError{Name: "order_at", err: errors.New("ent: missing required field \"order_at\"")}
	}
	if _, ok := oc.mutation.WhoID(); !ok {
		return &ValidationError{Name: "who", err: errors.New("ent: missing required edge \"who\"")}
	}
	if _, ok := oc.mutation.WhereID(); !ok {
		return &ValidationError{Name: "where", err: errors.New("ent: missing required edge \"where\"")}
	}
	return nil
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		_node = &Order{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: order.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.OrderAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldOrderAt,
		})
		_node.OrderAt = value
	}
	if value, ok := oc.mutation.CookingAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldCookingAt,
		})
		_node.CookingAt = &value
	}
	if value, ok := oc.mutation.DeliverAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldDeliverAt,
		})
		_node.DeliverAt = &value
	}
	if value, ok := oc.mutation.CompleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldCompleteAt,
		})
		_node.CompleteAt = &value
	}
	if nodes := oc.mutation.WhoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.WhoTable,
			Columns: []string{order.WhoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.WhereIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   order.WhereTable,
			Columns: []string{order.WhereColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: restaurant.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   order.ItemsTable,
			Columns: []string{order.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: orderfield.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OrderCreateBulk is the builder for creating a bulk of Order entities.
type OrderCreateBulk struct {
	config
	builders []*OrderCreate
}

// Save creates the Order entities in the database.
func (ocb *OrderCreateBulk) Save(ctx context.Context) ([]*Order, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Order, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ocb *OrderCreateBulk) SaveX(ctx context.Context) []*Order {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
