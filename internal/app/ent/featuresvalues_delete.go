// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"products-service/internal/app/ent/featuresvalues"
	"products-service/internal/app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeaturesValuesDelete is the builder for deleting a FeaturesValues entity.
type FeaturesValuesDelete struct {
	config
	hooks    []Hook
	mutation *FeaturesValuesMutation
}

// Where appends a list predicates to the FeaturesValuesDelete builder.
func (fvd *FeaturesValuesDelete) Where(ps ...predicate.FeaturesValues) *FeaturesValuesDelete {
	fvd.mutation.Where(ps...)
	return fvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fvd *FeaturesValuesDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fvd.sqlExec, fvd.mutation, fvd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fvd *FeaturesValuesDelete) ExecX(ctx context.Context) int {
	n, err := fvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fvd *FeaturesValuesDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(featuresvalues.Table, sqlgraph.NewFieldSpec(featuresvalues.FieldID, field.TypeInt))
	if ps := fvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fvd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fvd.mutation.done = true
	return affected, err
}

// FeaturesValuesDeleteOne is the builder for deleting a single FeaturesValues entity.
type FeaturesValuesDeleteOne struct {
	fvd *FeaturesValuesDelete
}

// Where appends a list predicates to the FeaturesValuesDelete builder.
func (fvdo *FeaturesValuesDeleteOne) Where(ps ...predicate.FeaturesValues) *FeaturesValuesDeleteOne {
	fvdo.fvd.mutation.Where(ps...)
	return fvdo
}

// Exec executes the deletion query.
func (fvdo *FeaturesValuesDeleteOne) Exec(ctx context.Context) error {
	n, err := fvdo.fvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{featuresvalues.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fvdo *FeaturesValuesDeleteOne) ExecX(ctx context.Context) {
	if err := fvdo.Exec(ctx); err != nil {
		panic(err)
	}
}
