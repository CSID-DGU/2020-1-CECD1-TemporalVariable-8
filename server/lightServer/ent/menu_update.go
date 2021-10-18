// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"lightServer/ent/category"
	"lightServer/ent/file"
	"lightServer/ent/menu"
	"lightServer/ent/predicate"
	"lightServer/ent/restaurant"

	"github.com/Rhymond/go-money"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// MenuUpdate is the builder for updating Menu entities.
type MenuUpdate struct {
	config
	hooks      []Hook
	mutation   *MenuMutation
	predicates []predicate.Menu
}

// Where adds a new predicate for the builder.
func (mu *MenuUpdate) Where(ps ...predicate.Menu) *MenuUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetName sets the name field.
func (mu *MenuUpdate) SetName(s string) *MenuUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetDescription sets the description field.
func (mu *MenuUpdate) SetDescription(s string) *MenuUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetIsOption sets the isOption field.
func (mu *MenuUpdate) SetIsOption(b bool) *MenuUpdate {
	mu.mutation.SetIsOption(b)
	return mu
}

// SetPrice sets the price field.
func (mu *MenuUpdate) SetPrice(m *money.Money) *MenuUpdate {
	mu.mutation.SetPrice(m)
	return mu
}

// ClearPrice clears the value of price.
func (mu *MenuUpdate) ClearPrice() *MenuUpdate {
	mu.mutation.ClearPrice()
	return mu
}

// SetOwnerID sets the owner edge to Restaurant by id.
func (mu *MenuUpdate) SetOwnerID(id int) *MenuUpdate {
	mu.mutation.SetOwnerID(id)
	return mu
}

// SetOwner sets the owner edge to Restaurant.
func (mu *MenuUpdate) SetOwner(r *Restaurant) *MenuUpdate {
	return mu.SetOwnerID(r.ID)
}

// AddCategoryIDs adds the category edge to Category by ids.
func (mu *MenuUpdate) AddCategoryIDs(ids ...int) *MenuUpdate {
	mu.mutation.AddCategoryIDs(ids...)
	return mu
}

// AddCategory adds the category edges to Category.
func (mu *MenuUpdate) AddCategory(c ...*Category) *MenuUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mu.AddCategoryIDs(ids...)
}

// SetImagesID sets the images edge to File by id.
func (mu *MenuUpdate) SetImagesID(id int) *MenuUpdate {
	mu.mutation.SetImagesID(id)
	return mu
}

// SetNillableImagesID sets the images edge to File by id if the given value is not nil.
func (mu *MenuUpdate) SetNillableImagesID(id *int) *MenuUpdate {
	if id != nil {
		mu = mu.SetImagesID(*id)
	}
	return mu
}

// SetImages sets the images edge to File.
func (mu *MenuUpdate) SetImages(f *File) *MenuUpdate {
	return mu.SetImagesID(f.ID)
}

// AddOptionIDs adds the options edge to Menu by ids.
func (mu *MenuUpdate) AddOptionIDs(ids ...int) *MenuUpdate {
	mu.mutation.AddOptionIDs(ids...)
	return mu
}

// AddOptions adds the options edges to Menu.
func (mu *MenuUpdate) AddOptions(m ...*Menu) *MenuUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.AddOptionIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (mu *MenuUpdate) Mutation() *MenuMutation {
	return mu.mutation
}

// ClearOwner clears the "owner" edge to type Restaurant.
func (mu *MenuUpdate) ClearOwner() *MenuUpdate {
	mu.mutation.ClearOwner()
	return mu
}

// ClearCategory clears all "category" edges to type Category.
func (mu *MenuUpdate) ClearCategory() *MenuUpdate {
	mu.mutation.ClearCategory()
	return mu
}

// RemoveCategoryIDs removes the category edge to Category by ids.
func (mu *MenuUpdate) RemoveCategoryIDs(ids ...int) *MenuUpdate {
	mu.mutation.RemoveCategoryIDs(ids...)
	return mu
}

// RemoveCategory removes category edges to Category.
func (mu *MenuUpdate) RemoveCategory(c ...*Category) *MenuUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mu.RemoveCategoryIDs(ids...)
}

// ClearImages clears the "images" edge to type File.
func (mu *MenuUpdate) ClearImages() *MenuUpdate {
	mu.mutation.ClearImages()
	return mu
}

// ClearOptions clears all "options" edges to type Menu.
func (mu *MenuUpdate) ClearOptions() *MenuUpdate {
	mu.mutation.ClearOptions()
	return mu
}

// RemoveOptionIDs removes the options edge to Menu by ids.
func (mu *MenuUpdate) RemoveOptionIDs(ids ...int) *MenuUpdate {
	mu.mutation.RemoveOptionIDs(ids...)
	return mu
}

// RemoveOptions removes options edges to Menu.
func (mu *MenuUpdate) RemoveOptions(m ...*Menu) *MenuUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return mu.RemoveOptionIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MenuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MenuUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MenuUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MenuUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MenuUpdate) check() error {
	if v, ok := mu.mutation.Name(); ok {
		if err := menu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := mu.mutation.OwnerID(); mu.mutation.OwnerCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"owner\"")
	}
	return nil
}

func (mu *MenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: menu.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldName,
		})
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldDescription,
		})
	}
	if value, ok := mu.mutation.IsOption(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldIsOption,
		})
	}
	if value, ok := mu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: menu.FieldPrice,
		})
	}
	if mu.mutation.PriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: menu.FieldPrice,
		})
	}
	if mu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   menu.OwnerTable,
			Columns: []string{menu.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: restaurant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   menu.OwnerTable,
			Columns: []string{menu.OwnerColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.CategoryTable,
			Columns: menu.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !mu.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.CategoryTable,
			Columns: menu.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.CategoryTable,
			Columns: menu.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   menu.ImagesTable,
			Columns: []string{menu.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   menu.ImagesTable,
			Columns: []string{menu.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.OptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menu.OptionsTable,
			Columns: menu.OptionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedOptionsIDs(); len(nodes) > 0 && !mu.mutation.OptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menu.OptionsTable,
			Columns: menu.OptionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.OptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menu.OptionsTable,
			Columns: menu.OptionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MenuUpdateOne is the builder for updating a single Menu entity.
type MenuUpdateOne struct {
	config
	hooks    []Hook
	mutation *MenuMutation
}

// SetName sets the name field.
func (muo *MenuUpdateOne) SetName(s string) *MenuUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetDescription sets the description field.
func (muo *MenuUpdateOne) SetDescription(s string) *MenuUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetIsOption sets the isOption field.
func (muo *MenuUpdateOne) SetIsOption(b bool) *MenuUpdateOne {
	muo.mutation.SetIsOption(b)
	return muo
}

// SetPrice sets the price field.
func (muo *MenuUpdateOne) SetPrice(m *money.Money) *MenuUpdateOne {
	muo.mutation.SetPrice(m)
	return muo
}

// ClearPrice clears the value of price.
func (muo *MenuUpdateOne) ClearPrice() *MenuUpdateOne {
	muo.mutation.ClearPrice()
	return muo
}

// SetOwnerID sets the owner edge to Restaurant by id.
func (muo *MenuUpdateOne) SetOwnerID(id int) *MenuUpdateOne {
	muo.mutation.SetOwnerID(id)
	return muo
}

// SetOwner sets the owner edge to Restaurant.
func (muo *MenuUpdateOne) SetOwner(r *Restaurant) *MenuUpdateOne {
	return muo.SetOwnerID(r.ID)
}

// AddCategoryIDs adds the category edge to Category by ids.
func (muo *MenuUpdateOne) AddCategoryIDs(ids ...int) *MenuUpdateOne {
	muo.mutation.AddCategoryIDs(ids...)
	return muo
}

// AddCategory adds the category edges to Category.
func (muo *MenuUpdateOne) AddCategory(c ...*Category) *MenuUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return muo.AddCategoryIDs(ids...)
}

// SetImagesID sets the images edge to File by id.
func (muo *MenuUpdateOne) SetImagesID(id int) *MenuUpdateOne {
	muo.mutation.SetImagesID(id)
	return muo
}

// SetNillableImagesID sets the images edge to File by id if the given value is not nil.
func (muo *MenuUpdateOne) SetNillableImagesID(id *int) *MenuUpdateOne {
	if id != nil {
		muo = muo.SetImagesID(*id)
	}
	return muo
}

// SetImages sets the images edge to File.
func (muo *MenuUpdateOne) SetImages(f *File) *MenuUpdateOne {
	return muo.SetImagesID(f.ID)
}

// AddOptionIDs adds the options edge to Menu by ids.
func (muo *MenuUpdateOne) AddOptionIDs(ids ...int) *MenuUpdateOne {
	muo.mutation.AddOptionIDs(ids...)
	return muo
}

// AddOptions adds the options edges to Menu.
func (muo *MenuUpdateOne) AddOptions(m ...*Menu) *MenuUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.AddOptionIDs(ids...)
}

// Mutation returns the MenuMutation object of the builder.
func (muo *MenuUpdateOne) Mutation() *MenuMutation {
	return muo.mutation
}

// ClearOwner clears the "owner" edge to type Restaurant.
func (muo *MenuUpdateOne) ClearOwner() *MenuUpdateOne {
	muo.mutation.ClearOwner()
	return muo
}

// ClearCategory clears all "category" edges to type Category.
func (muo *MenuUpdateOne) ClearCategory() *MenuUpdateOne {
	muo.mutation.ClearCategory()
	return muo
}

// RemoveCategoryIDs removes the category edge to Category by ids.
func (muo *MenuUpdateOne) RemoveCategoryIDs(ids ...int) *MenuUpdateOne {
	muo.mutation.RemoveCategoryIDs(ids...)
	return muo
}

// RemoveCategory removes category edges to Category.
func (muo *MenuUpdateOne) RemoveCategory(c ...*Category) *MenuUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return muo.RemoveCategoryIDs(ids...)
}

// ClearImages clears the "images" edge to type File.
func (muo *MenuUpdateOne) ClearImages() *MenuUpdateOne {
	muo.mutation.ClearImages()
	return muo
}

// ClearOptions clears all "options" edges to type Menu.
func (muo *MenuUpdateOne) ClearOptions() *MenuUpdateOne {
	muo.mutation.ClearOptions()
	return muo
}

// RemoveOptionIDs removes the options edge to Menu by ids.
func (muo *MenuUpdateOne) RemoveOptionIDs(ids ...int) *MenuUpdateOne {
	muo.mutation.RemoveOptionIDs(ids...)
	return muo
}

// RemoveOptions removes options edges to Menu.
func (muo *MenuUpdateOne) RemoveOptions(m ...*Menu) *MenuUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return muo.RemoveOptionIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MenuUpdateOne) Save(ctx context.Context) (*Menu, error) {
	var (
		err  error
		node *Menu
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MenuUpdateOne) SaveX(ctx context.Context) *Menu {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MenuUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MenuUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MenuUpdateOne) check() error {
	if v, ok := muo.mutation.Name(); ok {
		if err := menu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := muo.mutation.OwnerID(); muo.mutation.OwnerCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"owner\"")
	}
	return nil
}

func (muo *MenuUpdateOne) sqlSave(ctx context.Context) (_node *Menu, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   menu.Table,
			Columns: menu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: menu.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Menu.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldName,
		})
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: menu.FieldDescription,
		})
	}
	if value, ok := muo.mutation.IsOption(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: menu.FieldIsOption,
		})
	}
	if value, ok := muo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: menu.FieldPrice,
		})
	}
	if muo.mutation.PriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: menu.FieldPrice,
		})
	}
	if muo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   menu.OwnerTable,
			Columns: []string{menu.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: restaurant.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   menu.OwnerTable,
			Columns: []string{menu.OwnerColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.CategoryTable,
			Columns: menu.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedCategoryIDs(); len(nodes) > 0 && !muo.mutation.CategoryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.CategoryTable,
			Columns: menu.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.CategoryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   menu.CategoryTable,
			Columns: menu.CategoryPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: category.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.ImagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   menu.ImagesTable,
			Columns: []string{menu.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ImagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   menu.ImagesTable,
			Columns: []string{menu.ImagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.OptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menu.OptionsTable,
			Columns: menu.OptionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedOptionsIDs(); len(nodes) > 0 && !muo.mutation.OptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menu.OptionsTable,
			Columns: menu.OptionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.OptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   menu.OptionsTable,
			Columns: menu.OptionsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: menu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Menu{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{menu.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
