// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"lightServer/ent/predicate"
	"lightServer/ent/restaurant"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// RestaurantDelete is the builder for deleting a Restaurant entity.
type RestaurantDelete struct {
	config
	hooks      []Hook
	mutation   *RestaurantMutation
	predicates []predicate.Restaurant
}

// Where adds a new predicate to the delete builder.
func (rd *RestaurantDelete) Where(ps ...predicate.Restaurant) *RestaurantDelete {
	rd.predicates = append(rd.predicates, ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *RestaurantDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rd.hooks) == 0 {
		affected, err = rd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RestaurantMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rd.mutation = mutation
			affected, err = rd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rd.hooks) - 1; i >= 0; i-- {
			mut = rd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *RestaurantDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *RestaurantDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: restaurant.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: restaurant.FieldID,
			},
		},
	}
	if ps := rd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
}

// RestaurantDeleteOne is the builder for deleting a single Restaurant entity.
type RestaurantDeleteOne struct {
	rd *RestaurantDelete
}

// Exec executes the deletion query.
func (rdo *RestaurantDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{restaurant.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *RestaurantDeleteOne) ExecX(ctx context.Context) {
	rdo.rd.ExecX(ctx)
}
