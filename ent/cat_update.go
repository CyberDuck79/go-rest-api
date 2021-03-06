// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"rest_api/ent/cat"
	"rest_api/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CatUpdate is the builder for updating Cat entities.
type CatUpdate struct {
	config
	hooks    []Hook
	mutation *CatMutation
}

// Where appends a list predicates to the CatUpdate builder.
func (cu *CatUpdate) Where(ps ...predicate.Cat) *CatUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetAge sets the "age" field.
func (cu *CatUpdate) SetAge(i int) *CatUpdate {
	cu.mutation.ResetAge()
	cu.mutation.SetAge(i)
	return cu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (cu *CatUpdate) SetNillableAge(i *int) *CatUpdate {
	if i != nil {
		cu.SetAge(*i)
	}
	return cu
}

// AddAge adds i to the "age" field.
func (cu *CatUpdate) AddAge(i int) *CatUpdate {
	cu.mutation.AddAge(i)
	return cu
}

// ClearAge clears the value of the "age" field.
func (cu *CatUpdate) ClearAge() *CatUpdate {
	cu.mutation.ClearAge()
	return cu
}

// SetName sets the "name" field.
func (cu *CatUpdate) SetName(s string) *CatUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CatUpdate) SetCreatedAt(t time.Time) *CatUpdate {
	cu.mutation.SetCreatedAt(t)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CatUpdate) SetNillableCreatedAt(t *time.Time) *CatUpdate {
	if t != nil {
		cu.SetCreatedAt(*t)
	}
	return cu
}

// SetColor sets the "color" field.
func (cu *CatUpdate) SetColor(c cat.Color) *CatUpdate {
	cu.mutation.SetColor(c)
	return cu
}

// SetFeral sets the "feral" field.
func (cu *CatUpdate) SetFeral(b bool) *CatUpdate {
	cu.mutation.SetFeral(b)
	return cu
}

// SetNillableFeral sets the "feral" field if the given value is not nil.
func (cu *CatUpdate) SetNillableFeral(b *bool) *CatUpdate {
	if b != nil {
		cu.SetFeral(*b)
	}
	return cu
}

// Mutation returns the CatMutation object of the builder.
func (cu *CatUpdate) Mutation() *CatMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CatUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CatUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CatUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CatUpdate) check() error {
	if v, ok := cu.mutation.Age(); ok {
		if err := cat.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if v, ok := cu.mutation.Color(); ok {
		if err := cat.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf("ent: validator failed for field \"color\": %w", err)}
		}
	}
	return nil
}

func (cu *CatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cat.Table,
			Columns: cat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cat.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cat.FieldAge,
		})
	}
	if value, ok := cu.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cat.FieldAge,
		})
	}
	if cu.mutation.AgeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: cat.FieldAge,
		})
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cat.FieldName,
		})
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cat.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.Color(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: cat.FieldColor,
		})
	}
	if value, ok := cu.mutation.Feral(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: cat.FieldFeral,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CatUpdateOne is the builder for updating a single Cat entity.
type CatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CatMutation
}

// SetAge sets the "age" field.
func (cuo *CatUpdateOne) SetAge(i int) *CatUpdateOne {
	cuo.mutation.ResetAge()
	cuo.mutation.SetAge(i)
	return cuo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (cuo *CatUpdateOne) SetNillableAge(i *int) *CatUpdateOne {
	if i != nil {
		cuo.SetAge(*i)
	}
	return cuo
}

// AddAge adds i to the "age" field.
func (cuo *CatUpdateOne) AddAge(i int) *CatUpdateOne {
	cuo.mutation.AddAge(i)
	return cuo
}

// ClearAge clears the value of the "age" field.
func (cuo *CatUpdateOne) ClearAge() *CatUpdateOne {
	cuo.mutation.ClearAge()
	return cuo
}

// SetName sets the "name" field.
func (cuo *CatUpdateOne) SetName(s string) *CatUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CatUpdateOne) SetCreatedAt(t time.Time) *CatUpdateOne {
	cuo.mutation.SetCreatedAt(t)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CatUpdateOne) SetNillableCreatedAt(t *time.Time) *CatUpdateOne {
	if t != nil {
		cuo.SetCreatedAt(*t)
	}
	return cuo
}

// SetColor sets the "color" field.
func (cuo *CatUpdateOne) SetColor(c cat.Color) *CatUpdateOne {
	cuo.mutation.SetColor(c)
	return cuo
}

// SetFeral sets the "feral" field.
func (cuo *CatUpdateOne) SetFeral(b bool) *CatUpdateOne {
	cuo.mutation.SetFeral(b)
	return cuo
}

// SetNillableFeral sets the "feral" field if the given value is not nil.
func (cuo *CatUpdateOne) SetNillableFeral(b *bool) *CatUpdateOne {
	if b != nil {
		cuo.SetFeral(*b)
	}
	return cuo
}

// Mutation returns the CatMutation object of the builder.
func (cuo *CatUpdateOne) Mutation() *CatMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CatUpdateOne) Select(field string, fields ...string) *CatUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Cat entity.
func (cuo *CatUpdateOne) Save(ctx context.Context) (*Cat, error) {
	var (
		err  error
		node *Cat
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CatUpdateOne) SaveX(ctx context.Context) *Cat {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CatUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CatUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CatUpdateOne) check() error {
	if v, ok := cuo.mutation.Age(); ok {
		if err := cat.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf("ent: validator failed for field \"age\": %w", err)}
		}
	}
	if v, ok := cuo.mutation.Color(); ok {
		if err := cat.ColorValidator(v); err != nil {
			return &ValidationError{Name: "color", err: fmt.Errorf("ent: validator failed for field \"color\": %w", err)}
		}
	}
	return nil
}

func (cuo *CatUpdateOne) sqlSave(ctx context.Context) (_node *Cat, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cat.Table,
			Columns: cat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cat.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Cat.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cat.FieldID)
		for _, f := range fields {
			if !cat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Age(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cat.FieldAge,
		})
	}
	if value, ok := cuo.mutation.AddedAge(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: cat.FieldAge,
		})
	}
	if cuo.mutation.AgeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: cat.FieldAge,
		})
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: cat.FieldName,
		})
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: cat.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.Color(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: cat.FieldColor,
		})
	}
	if value, ok := cuo.mutation.Feral(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: cat.FieldFeral,
		})
	}
	_node = &Cat{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
