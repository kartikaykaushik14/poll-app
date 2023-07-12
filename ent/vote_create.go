// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"poll-app/ent/vote"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// VoteCreate is the builder for creating a Vote entity.
type VoteCreate struct {
	config
	mutation *VoteMutation
	hooks    []Hook
}

// SetUserId sets the "userId" field.
func (vc *VoteCreate) SetUserId(s string) *VoteCreate {
	vc.mutation.SetUserId(s)
	return vc
}

// SetPollOptionId sets the "pollOptionId" field.
func (vc *VoteCreate) SetPollOptionId(s string) *VoteCreate {
	vc.mutation.SetPollOptionId(s)
	return vc
}

// Mutation returns the VoteMutation object of the builder.
func (vc *VoteCreate) Mutation() *VoteMutation {
	return vc.mutation
}

// Save creates the Vote in the database.
func (vc *VoteCreate) Save(ctx context.Context) (*Vote, error) {
	return withHooks(ctx, vc.sqlSave, vc.mutation, vc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (vc *VoteCreate) SaveX(ctx context.Context) *Vote {
	v, err := vc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vc *VoteCreate) Exec(ctx context.Context) error {
	_, err := vc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vc *VoteCreate) ExecX(ctx context.Context) {
	if err := vc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vc *VoteCreate) check() error {
	if _, ok := vc.mutation.UserId(); !ok {
		return &ValidationError{Name: "userId", err: errors.New(`ent: missing required field "Vote.userId"`)}
	}
	if _, ok := vc.mutation.PollOptionId(); !ok {
		return &ValidationError{Name: "pollOptionId", err: errors.New(`ent: missing required field "Vote.pollOptionId"`)}
	}
	return nil
}

func (vc *VoteCreate) sqlSave(ctx context.Context) (*Vote, error) {
	if err := vc.check(); err != nil {
		return nil, err
	}
	_node, _spec := vc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	vc.mutation.id = &_node.ID
	vc.mutation.done = true
	return _node, nil
}

func (vc *VoteCreate) createSpec() (*Vote, *sqlgraph.CreateSpec) {
	var (
		_node = &Vote{config: vc.config}
		_spec = sqlgraph.NewCreateSpec(vote.Table, sqlgraph.NewFieldSpec(vote.FieldID, field.TypeInt))
	)
	if value, ok := vc.mutation.UserId(); ok {
		_spec.SetField(vote.FieldUserId, field.TypeString, value)
		_node.UserId = value
	}
	if value, ok := vc.mutation.PollOptionId(); ok {
		_spec.SetField(vote.FieldPollOptionId, field.TypeString, value)
		_node.PollOptionId = value
	}
	return _node, _spec
}

// VoteCreateBulk is the builder for creating many Vote entities in bulk.
type VoteCreateBulk struct {
	config
	builders []*VoteCreate
}

// Save creates the Vote entities in the database.
func (vcb *VoteCreateBulk) Save(ctx context.Context) ([]*Vote, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vcb.builders))
	nodes := make([]*Vote, len(vcb.builders))
	mutators := make([]Mutator, len(vcb.builders))
	for i := range vcb.builders {
		func(i int, root context.Context) {
			builder := vcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VoteMutation)
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
					_, err = mutators[i+1].Mutate(root, vcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, vcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vcb *VoteCreateBulk) SaveX(ctx context.Context) []*Vote {
	v, err := vcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vcb *VoteCreateBulk) Exec(ctx context.Context) error {
	_, err := vcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vcb *VoteCreateBulk) ExecX(ctx context.Context) {
	if err := vcb.Exec(ctx); err != nil {
		panic(err)
	}
}
