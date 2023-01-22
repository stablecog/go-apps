// Code generated by ent, DO NOT EDIT.

package upscaleoutput

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldLTE(FieldID, id))
}

// ImageURL applies equality check predicate on the "image_url" field. It's identical to ImageURLEQ.
func ImageURL(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEQ(FieldImageURL, v))
}

// UpscaleID applies equality check predicate on the "upscale_id" field. It's identical to UpscaleIDEQ.
func UpscaleID(v uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEQ(FieldUpscaleID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEQ(FieldUpdatedAt, v))
}

// ImageURLEQ applies the EQ predicate on the "image_url" field.
func ImageURLEQ(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEQ(FieldImageURL, v))
}

// ImageURLNEQ applies the NEQ predicate on the "image_url" field.
func ImageURLNEQ(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldNEQ(FieldImageURL, v))
}

// ImageURLIn applies the In predicate on the "image_url" field.
func ImageURLIn(vs ...string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldIn(FieldImageURL, vs...))
}

// ImageURLNotIn applies the NotIn predicate on the "image_url" field.
func ImageURLNotIn(vs ...string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldNotIn(FieldImageURL, vs...))
}

// ImageURLGT applies the GT predicate on the "image_url" field.
func ImageURLGT(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldGT(FieldImageURL, v))
}

// ImageURLGTE applies the GTE predicate on the "image_url" field.
func ImageURLGTE(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldGTE(FieldImageURL, v))
}

// ImageURLLT applies the LT predicate on the "image_url" field.
func ImageURLLT(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldLT(FieldImageURL, v))
}

// ImageURLLTE applies the LTE predicate on the "image_url" field.
func ImageURLLTE(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldLTE(FieldImageURL, v))
}

// ImageURLContains applies the Contains predicate on the "image_url" field.
func ImageURLContains(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldContains(FieldImageURL, v))
}

// ImageURLHasPrefix applies the HasPrefix predicate on the "image_url" field.
func ImageURLHasPrefix(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldHasPrefix(FieldImageURL, v))
}

// ImageURLHasSuffix applies the HasSuffix predicate on the "image_url" field.
func ImageURLHasSuffix(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldHasSuffix(FieldImageURL, v))
}

// ImageURLEqualFold applies the EqualFold predicate on the "image_url" field.
func ImageURLEqualFold(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEqualFold(FieldImageURL, v))
}

// ImageURLContainsFold applies the ContainsFold predicate on the "image_url" field.
func ImageURLContainsFold(v string) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldContainsFold(FieldImageURL, v))
}

// UpscaleIDEQ applies the EQ predicate on the "upscale_id" field.
func UpscaleIDEQ(v uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEQ(FieldUpscaleID, v))
}

// UpscaleIDNEQ applies the NEQ predicate on the "upscale_id" field.
func UpscaleIDNEQ(v uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldNEQ(FieldUpscaleID, v))
}

// UpscaleIDIn applies the In predicate on the "upscale_id" field.
func UpscaleIDIn(vs ...uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldIn(FieldUpscaleID, vs...))
}

// UpscaleIDNotIn applies the NotIn predicate on the "upscale_id" field.
func UpscaleIDNotIn(vs ...uuid.UUID) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldNotIn(FieldUpscaleID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasUpscales applies the HasEdge predicate on the "upscales" edge.
func HasUpscales() predicate.UpscaleOutput {
	return predicate.UpscaleOutput(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UpscalesTable, UpscalesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpscalesWith applies the HasEdge predicate on the "upscales" edge with a given conditions (other predicates).
func HasUpscalesWith(preds ...predicate.Upscale) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UpscalesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UpscalesTable, UpscalesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UpscaleOutput) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UpscaleOutput) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UpscaleOutput) predicate.UpscaleOutput {
	return predicate.UpscaleOutput(func(s *sql.Selector) {
		p(s.Not())
	})
}
