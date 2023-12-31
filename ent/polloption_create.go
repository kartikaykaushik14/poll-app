// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"poll-app/ent/poll"
	"poll-app/ent/polloption"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PollOptionCreate is the builder for creating a PollOption entity.
type PollOptionCreate struct {
	config
	mutation *PollOptionMutation
	hooks    []Hook
}

// SetOption sets the "option" field.
func (poc *PollOptionCreate) SetOption(s string) *PollOptionCreate {
	poc.mutation.SetOption(s)
	return poc
}

// AddPollIDs adds the "poll" edge to the Poll entity by IDs.
func (poc *PollOptionCreate) AddPollIDs(ids ...int) *PollOptionCreate {
	poc.mutation.AddPollIDs(ids...)
	return poc
}

// AddPoll adds the "poll" edges to the Poll entity.
func (poc *PollOptionCreate) AddPoll(p ...*Poll) *PollOptionCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return poc.AddPollIDs(ids...)
}

// Mutation returns the PollOptionMutation object of the builder.
func (poc *PollOptionCreate) Mutation() *PollOptionMutation {
	return poc.mutation
}

// Save creates the PollOption in the database.
func (poc *PollOptionCreate) Save(ctx context.Context) (*PollOption, error) {
	return withHooks(ctx, poc.sqlSave, poc.mutation, poc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (poc *PollOptionCreate) SaveX(ctx context.Context) *PollOption {
	v, err := poc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (poc *PollOptionCreate) Exec(ctx context.Context) error {
	_, err := poc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (poc *PollOptionCreate) ExecX(ctx context.Context) {
	if err := poc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (poc *PollOptionCreate) check() error {
	if _, ok := poc.mutation.Option(); !ok {
		return &ValidationError{Name: "option", err: errors.New(`ent: missing required field "PollOption.option"`)}
	}
	return nil
}

func (poc *PollOptionCreate) sqlSave(ctx context.Context) (*PollOption, error) {
	if err := poc.check(); err != nil {
		return nil, err
	}
	_node, _spec := poc.createSpec()
	if err := sqlgraph.CreateNode(ctx, poc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	poc.mutation.id = &_node.ID
	poc.mutation.done = true
	return _node, nil
}

func (poc *PollOptionCreate) createSpec() (*PollOption, *sqlgraph.CreateSpec) {
	var (
		_node = &PollOption{config: poc.config}
		_spec = sqlgraph.NewCreateSpec(polloption.Table, sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt))
	)
	if value, ok := poc.mutation.Option(); ok {
		_spec.SetField(polloption.FieldOption, field.TypeString, value)
		_node.Option = value
	}
	if nodes := poc.mutation.PollIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   polloption.PollTable,
			Columns: polloption.PollPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PollOptionCreateBulk is the builder for creating many PollOption entities in bulk.
type PollOptionCreateBulk struct {
	config
	builders []*PollOptionCreate
}

// Save creates the PollOption entities in the database.
func (pocb *PollOptionCreateBulk) Save(ctx context.Context) ([]*PollOption, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pocb.builders))
	nodes := make([]*PollOption, len(pocb.builders))
	mutators := make([]Mutator, len(pocb.builders))
	for i := range pocb.builders {
		func(i int, root context.Context) {
			builder := pocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PollOptionMutation)
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
					_, err = mutators[i+1].Mutate(root, pocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pocb *PollOptionCreateBulk) SaveX(ctx context.Context) []*PollOption {
	v, err := pocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pocb *PollOptionCreateBulk) Exec(ctx context.Context) error {
	_, err := pocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pocb *PollOptionCreateBulk) ExecX(ctx context.Context) {
	if err := pocb.Exec(ctx); err != nil {
		panic(err)
	}
}
