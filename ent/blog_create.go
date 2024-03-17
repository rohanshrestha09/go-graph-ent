// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/rohanshrestha09/go-graph-ent/common/enums"
	"github.com/rohanshrestha09/go-graph-ent/ent/blog"
	"github.com/rohanshrestha09/go-graph-ent/ent/user"
)

// BlogCreate is the builder for creating a Blog entity.
type BlogCreate struct {
	config
	mutation *BlogMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bc *BlogCreate) SetCreatedAt(t time.Time) *BlogCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BlogCreate) SetNillableCreatedAt(t *time.Time) *BlogCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BlogCreate) SetUpdatedAt(t time.Time) *BlogCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BlogCreate) SetNillableUpdatedAt(t *time.Time) *BlogCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetTitle sets the "title" field.
func (bc *BlogCreate) SetTitle(s string) *BlogCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetSlug sets the "slug" field.
func (bc *BlogCreate) SetSlug(s string) *BlogCreate {
	bc.mutation.SetSlug(s)
	return bc
}

// SetContent sets the "content" field.
func (bc *BlogCreate) SetContent(s string) *BlogCreate {
	bc.mutation.SetContent(s)
	return bc
}

// SetImage sets the "image" field.
func (bc *BlogCreate) SetImage(s string) *BlogCreate {
	bc.mutation.SetImage(s)
	return bc
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (bc *BlogCreate) SetNillableImage(s *string) *BlogCreate {
	if s != nil {
		bc.SetImage(*s)
	}
	return bc
}

// SetStatus sets the "status" field.
func (bc *BlogCreate) SetStatus(es enums.BlogStatus) *BlogCreate {
	bc.mutation.SetStatus(es)
	return bc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (bc *BlogCreate) SetNillableStatus(es *enums.BlogStatus) *BlogCreate {
	if es != nil {
		bc.SetStatus(*es)
	}
	return bc
}

// SetUserID sets the "user_id" field.
func (bc *BlogCreate) SetUserID(u uuid.UUID) *BlogCreate {
	bc.mutation.SetUserID(u)
	return bc
}

// SetID sets the "id" field.
func (bc *BlogCreate) SetID(i int) *BlogCreate {
	bc.mutation.SetID(i)
	return bc
}

// SetUser sets the "user" edge to the User entity.
func (bc *BlogCreate) SetUser(u *User) *BlogCreate {
	return bc.SetUserID(u.ID)
}

// Mutation returns the BlogMutation object of the builder.
func (bc *BlogCreate) Mutation() *BlogMutation {
	return bc.mutation
}

// Save creates the Blog in the database.
func (bc *BlogCreate) Save(ctx context.Context) (*Blog, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlogCreate) SaveX(ctx context.Context) *Blog {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlogCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlogCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BlogCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := blog.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := blog.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
	if _, ok := bc.mutation.Status(); !ok {
		v := blog.DefaultStatus
		bc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlogCreate) check() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Blog.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Blog.updated_at"`)}
	}
	if _, ok := bc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Blog.title"`)}
	}
	if _, ok := bc.mutation.Slug(); !ok {
		return &ValidationError{Name: "slug", err: errors.New(`ent: missing required field "Blog.slug"`)}
	}
	if _, ok := bc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Blog.content"`)}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Blog.status"`)}
	}
	if v, ok := bc.mutation.Status(); ok {
		if err := blog.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Blog.status": %w`, err)}
		}
	}
	if _, ok := bc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Blog.user_id"`)}
	}
	if _, ok := bc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Blog.user"`)}
	}
	return nil
}

func (bc *BlogCreate) sqlSave(ctx context.Context) (*Blog, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BlogCreate) createSpec() (*Blog, *sqlgraph.CreateSpec) {
	var (
		_node = &Blog{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(blog.Table, sqlgraph.NewFieldSpec(blog.FieldID, field.TypeInt))
	)
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(blog.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(blog.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(blog.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.Slug(); ok {
		_spec.SetField(blog.FieldSlug, field.TypeString, value)
		_node.Slug = value
	}
	if value, ok := bc.mutation.Content(); ok {
		_spec.SetField(blog.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := bc.mutation.Image(); ok {
		_spec.SetField(blog.FieldImage, field.TypeString, value)
		_node.Image = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.SetField(blog.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if nodes := bc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   blog.UserTable,
			Columns: []string{blog.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BlogCreateBulk is the builder for creating many Blog entities in bulk.
type BlogCreateBulk struct {
	config
	err      error
	builders []*BlogCreate
}

// Save creates the Blog entities in the database.
func (bcb *BlogCreateBulk) Save(ctx context.Context) ([]*Blog, error) {
	if bcb.err != nil {
		return nil, bcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blog, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlogCreateBulk) SaveX(ctx context.Context) []*Blog {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlogCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlogCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
