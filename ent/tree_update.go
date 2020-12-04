// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"laforge/ent/predicate"
	"laforge/ent/tree"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TreeUpdate is the builder for updating Tree entities.
type TreeUpdate struct {
	config
	hooks      []Hook
	mutation   *TreeMutation
	predicates []predicate.Tree
}

// Where adds a new predicate for the builder.
func (tu *TreeUpdate) Where(ps ...predicate.Tree) *TreeUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// AddTargetIDs adds the target edge to Tree by ids.
func (tu *TreeUpdate) AddTargetIDs(ids ...string) *TreeUpdate {
	tu.mutation.AddTargetIDs(ids...)
	return tu
}

// AddTarget adds the target edges to Tree.
func (tu *TreeUpdate) AddTarget(t ...*Tree) *TreeUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTargetIDs(ids...)
}

// AddSourceIDs adds the source edge to Tree by ids.
func (tu *TreeUpdate) AddSourceIDs(ids ...string) *TreeUpdate {
	tu.mutation.AddSourceIDs(ids...)
	return tu
}

// AddSource adds the source edges to Tree.
func (tu *TreeUpdate) AddSource(t ...*Tree) *TreeUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddSourceIDs(ids...)
}

// Mutation returns the TreeMutation object of the builder.
func (tu *TreeUpdate) Mutation() *TreeMutation {
	return tu.mutation
}

// ClearTarget clears all "target" edges to type Tree.
func (tu *TreeUpdate) ClearTarget() *TreeUpdate {
	tu.mutation.ClearTarget()
	return tu
}

// RemoveTargetIDs removes the target edge to Tree by ids.
func (tu *TreeUpdate) RemoveTargetIDs(ids ...string) *TreeUpdate {
	tu.mutation.RemoveTargetIDs(ids...)
	return tu
}

// RemoveTarget removes target edges to Tree.
func (tu *TreeUpdate) RemoveTarget(t ...*Tree) *TreeUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTargetIDs(ids...)
}

// ClearSource clears all "source" edges to type Tree.
func (tu *TreeUpdate) ClearSource() *TreeUpdate {
	tu.mutation.ClearSource()
	return tu
}

// RemoveSourceIDs removes the source edge to Tree by ids.
func (tu *TreeUpdate) RemoveSourceIDs(ids ...string) *TreeUpdate {
	tu.mutation.RemoveSourceIDs(ids...)
	return tu
}

// RemoveSource removes source edges to Tree.
func (tu *TreeUpdate) RemoveSource(t ...*Tree) *TreeUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveSourceIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TreeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TreeUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TreeUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TreeUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TreeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tree.Table,
			Columns: tree.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tree.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if tu.mutation.TargetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tree.TargetTable,
			Columns: tree.TargetPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTargetIDs(); len(nodes) > 0 && !tu.mutation.TargetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tree.TargetTable,
			Columns: tree.TargetPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tree.TargetTable,
			Columns: tree.TargetPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tree.SourceTable,
			Columns: tree.SourcePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedSourceIDs(); len(nodes) > 0 && !tu.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tree.SourceTable,
			Columns: tree.SourcePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tree.SourceTable,
			Columns: tree.SourcePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tree.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TreeUpdateOne is the builder for updating a single Tree entity.
type TreeUpdateOne struct {
	config
	hooks    []Hook
	mutation *TreeMutation
}

// AddTargetIDs adds the target edge to Tree by ids.
func (tuo *TreeUpdateOne) AddTargetIDs(ids ...string) *TreeUpdateOne {
	tuo.mutation.AddTargetIDs(ids...)
	return tuo
}

// AddTarget adds the target edges to Tree.
func (tuo *TreeUpdateOne) AddTarget(t ...*Tree) *TreeUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTargetIDs(ids...)
}

// AddSourceIDs adds the source edge to Tree by ids.
func (tuo *TreeUpdateOne) AddSourceIDs(ids ...string) *TreeUpdateOne {
	tuo.mutation.AddSourceIDs(ids...)
	return tuo
}

// AddSource adds the source edges to Tree.
func (tuo *TreeUpdateOne) AddSource(t ...*Tree) *TreeUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddSourceIDs(ids...)
}

// Mutation returns the TreeMutation object of the builder.
func (tuo *TreeUpdateOne) Mutation() *TreeMutation {
	return tuo.mutation
}

// ClearTarget clears all "target" edges to type Tree.
func (tuo *TreeUpdateOne) ClearTarget() *TreeUpdateOne {
	tuo.mutation.ClearTarget()
	return tuo
}

// RemoveTargetIDs removes the target edge to Tree by ids.
func (tuo *TreeUpdateOne) RemoveTargetIDs(ids ...string) *TreeUpdateOne {
	tuo.mutation.RemoveTargetIDs(ids...)
	return tuo
}

// RemoveTarget removes target edges to Tree.
func (tuo *TreeUpdateOne) RemoveTarget(t ...*Tree) *TreeUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTargetIDs(ids...)
}

// ClearSource clears all "source" edges to type Tree.
func (tuo *TreeUpdateOne) ClearSource() *TreeUpdateOne {
	tuo.mutation.ClearSource()
	return tuo
}

// RemoveSourceIDs removes the source edge to Tree by ids.
func (tuo *TreeUpdateOne) RemoveSourceIDs(ids ...string) *TreeUpdateOne {
	tuo.mutation.RemoveSourceIDs(ids...)
	return tuo
}

// RemoveSource removes source edges to Tree.
func (tuo *TreeUpdateOne) RemoveSource(t ...*Tree) *TreeUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveSourceIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (tuo *TreeUpdateOne) Save(ctx context.Context) (*Tree, error) {
	var (
		err  error
		node *Tree
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TreeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TreeUpdateOne) SaveX(ctx context.Context) *Tree {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TreeUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TreeUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TreeUpdateOne) sqlSave(ctx context.Context) (_node *Tree, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tree.Table,
			Columns: tree.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: tree.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Tree.ID for update")}
	}
	_spec.Node.ID.Value = id
	if tuo.mutation.TargetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tree.TargetTable,
			Columns: tree.TargetPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTargetIDs(); len(nodes) > 0 && !tuo.mutation.TargetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tree.TargetTable,
			Columns: tree.TargetPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TargetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tree.TargetTable,
			Columns: tree.TargetPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tree.SourceTable,
			Columns: tree.SourcePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedSourceIDs(); len(nodes) > 0 && !tuo.mutation.SourceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tree.SourceTable,
			Columns: tree.SourcePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.SourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tree.SourceTable,
			Columns: tree.SourcePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: tree.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tree{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tree.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}