// Code generated by entc, DO NOT EDIT.

package ent

import (
	"bitsports/ent/category"
	"bitsports/ent/product"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CategoryCreate is the builder for creating a Category entity.
type CategoryCreate struct {
	config
	mutation *CategoryMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CategoryCreate) SetName(s string) *CategoryCreate {
	cc.mutation.SetName(s)
	return cc
}

// AddProductIDs adds the "products" edge to the Product entity by IDs.
func (cc *CategoryCreate) AddProductIDs(ids ...int) *CategoryCreate {
	cc.mutation.AddProductIDs(ids...)
	return cc
}

// AddProducts adds the "products" edges to the Product entity.
func (cc *CategoryCreate) AddProducts(p ...*Product) *CategoryCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProductIDs(ids...)
}

// Mutation returns the CategoryMutation object of the builder.
func (cc *CategoryCreate) Mutation() *CategoryMutation {
	return cc.mutation
}

// Save creates the Category in the database.
func (cc *CategoryCreate) Save(ctx context.Context) (*Category, error) {
	var (
		err  error
		node *Category
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CategoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CategoryCreate) SaveX(ctx context.Context) *Category {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CategoryCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CategoryCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CategoryCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Category.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := category.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Category.name": %w`, err)}
		}
	}
	return nil
}

func (cc *CategoryCreate) sqlSave(ctx context.Context) (*Category, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CategoryCreate) createSpec() (*Category, *sqlgraph.CreateSpec) {
	var (
		_node = &Category{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: category.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: category.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: category.FieldName,
		})
		_node.Name = value
	}
	if nodes := cc.mutation.ProductsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   category.ProductsTable,
			Columns: category.ProductsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
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

// CategoryCreateBulk is the builder for creating many Category entities in bulk.
type CategoryCreateBulk struct {
	config
	builders []*CategoryCreate
}

// Save creates the Category entities in the database.
func (ccb *CategoryCreateBulk) Save(ctx context.Context) ([]*Category, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Category, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CategoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CategoryCreateBulk) SaveX(ctx context.Context) []*Category {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CategoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CategoryCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
