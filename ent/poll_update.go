// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"poll-app/ent/poll"
	"poll-app/ent/polloption"
	"poll-app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PollUpdate is the builder for updating Poll entities.
type PollUpdate struct {
	config
	hooks    []Hook
	mutation *PollMutation
}

// Where appends a list predicates to the PollUpdate builder.
func (pu *PollUpdate) Where(ps ...predicate.Poll) *PollUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetQuestion sets the "question" field.
func (pu *PollUpdate) SetQuestion(s string) *PollUpdate {
	pu.mutation.SetQuestion(s)
	return pu
}

// AddPollOptionIDs adds the "poll_options" edge to the PollOption entity by IDs.
func (pu *PollUpdate) AddPollOptionIDs(ids ...int) *PollUpdate {
	pu.mutation.AddPollOptionIDs(ids...)
	return pu
}

// AddPollOptions adds the "poll_options" edges to the PollOption entity.
func (pu *PollUpdate) AddPollOptions(p ...*PollOption) *PollUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPollOptionIDs(ids...)
}

// Mutation returns the PollMutation object of the builder.
func (pu *PollUpdate) Mutation() *PollMutation {
	return pu.mutation
}

// ClearPollOptions clears all "poll_options" edges to the PollOption entity.
func (pu *PollUpdate) ClearPollOptions() *PollUpdate {
	pu.mutation.ClearPollOptions()
	return pu
}

// RemovePollOptionIDs removes the "poll_options" edge to PollOption entities by IDs.
func (pu *PollUpdate) RemovePollOptionIDs(ids ...int) *PollUpdate {
	pu.mutation.RemovePollOptionIDs(ids...)
	return pu
}

// RemovePollOptions removes "poll_options" edges to PollOption entities.
func (pu *PollUpdate) RemovePollOptions(p ...*PollOption) *PollUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePollOptionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PollUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PollUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PollUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PollUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PollUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(poll.Table, poll.Columns, sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Question(); ok {
		_spec.SetField(poll.FieldQuestion, field.TypeString, value)
	}
	if pu.mutation.PollOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: poll.PollOptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedPollOptionsIDs(); len(nodes) > 0 && !pu.mutation.PollOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: poll.PollOptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PollOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: poll.PollOptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{poll.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PollUpdateOne is the builder for updating a single Poll entity.
type PollUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PollMutation
}

// SetQuestion sets the "question" field.
func (puo *PollUpdateOne) SetQuestion(s string) *PollUpdateOne {
	puo.mutation.SetQuestion(s)
	return puo
}

// AddPollOptionIDs adds the "poll_options" edge to the PollOption entity by IDs.
func (puo *PollUpdateOne) AddPollOptionIDs(ids ...int) *PollUpdateOne {
	puo.mutation.AddPollOptionIDs(ids...)
	return puo
}

// AddPollOptions adds the "poll_options" edges to the PollOption entity.
func (puo *PollUpdateOne) AddPollOptions(p ...*PollOption) *PollUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPollOptionIDs(ids...)
}

// Mutation returns the PollMutation object of the builder.
func (puo *PollUpdateOne) Mutation() *PollMutation {
	return puo.mutation
}

// ClearPollOptions clears all "poll_options" edges to the PollOption entity.
func (puo *PollUpdateOne) ClearPollOptions() *PollUpdateOne {
	puo.mutation.ClearPollOptions()
	return puo
}

// RemovePollOptionIDs removes the "poll_options" edge to PollOption entities by IDs.
func (puo *PollUpdateOne) RemovePollOptionIDs(ids ...int) *PollUpdateOne {
	puo.mutation.RemovePollOptionIDs(ids...)
	return puo
}

// RemovePollOptions removes "poll_options" edges to PollOption entities.
func (puo *PollUpdateOne) RemovePollOptions(p ...*PollOption) *PollUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePollOptionIDs(ids...)
}

// Where appends a list predicates to the PollUpdate builder.
func (puo *PollUpdateOne) Where(ps ...predicate.Poll) *PollUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PollUpdateOne) Select(field string, fields ...string) *PollUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Poll entity.
func (puo *PollUpdateOne) Save(ctx context.Context) (*Poll, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PollUpdateOne) SaveX(ctx context.Context) *Poll {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PollUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PollUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PollUpdateOne) sqlSave(ctx context.Context) (_node *Poll, err error) {
	_spec := sqlgraph.NewUpdateSpec(poll.Table, poll.Columns, sqlgraph.NewFieldSpec(poll.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Poll.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, poll.FieldID)
		for _, f := range fields {
			if !poll.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != poll.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Question(); ok {
		_spec.SetField(poll.FieldQuestion, field.TypeString, value)
	}
	if puo.mutation.PollOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: poll.PollOptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedPollOptionsIDs(); len(nodes) > 0 && !puo.mutation.PollOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: poll.PollOptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PollOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   poll.PollOptionsTable,
			Columns: poll.PollOptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(polloption.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Poll{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{poll.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
