// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/rohanshrestha09/go-graph-ent/ent/blog"
	"github.com/rohanshrestha09/go-graph-ent/ent/predicate"
)

// BlogUpdate is the builder for updating Blog entities.
type BlogUpdate struct {
	config
	hooks    []Hook
	mutation *BlogMutation
}

// Where appends a list predicates to the BlogUpdate builder.
func (bu *BlogUpdate) Where(ps ...predicate.Blog) *BlogUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BlogUpdate) SetUpdatedAt(t time.Time) *BlogUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// Mutation returns the BlogMutation object of the builder.
func (bu *BlogUpdate) Mutation() *BlogMutation {
	return bu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlogUpdate) Save(ctx context.Context) (int, error) {
	bu.defaults()
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlogUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlogUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlogUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BlogUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := blog.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

func (bu *BlogUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(blog.Table, blog.Columns, sqlgraph.NewFieldSpec(blog.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(blog.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BlogUpdateOne is the builder for updating a single Blog entity.
type BlogUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlogMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BlogUpdateOne) SetUpdatedAt(t time.Time) *BlogUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// Mutation returns the BlogMutation object of the builder.
func (buo *BlogUpdateOne) Mutation() *BlogMutation {
	return buo.mutation
}

// Where appends a list predicates to the BlogUpdate builder.
func (buo *BlogUpdateOne) Where(ps ...predicate.Blog) *BlogUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlogUpdateOne) Select(field string, fields ...string) *BlogUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blog entity.
func (buo *BlogUpdateOne) Save(ctx context.Context) (*Blog, error) {
	buo.defaults()
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlogUpdateOne) SaveX(ctx context.Context) *Blog {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlogUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlogUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BlogUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := blog.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

func (buo *BlogUpdateOne) sqlSave(ctx context.Context) (_node *Blog, err error) {
	_spec := sqlgraph.NewUpdateSpec(blog.Table, blog.Columns, sqlgraph.NewFieldSpec(blog.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blog.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blog.FieldID)
		for _, f := range fields {
			if !blog.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(blog.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &Blog{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blog.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
