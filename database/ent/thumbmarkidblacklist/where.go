// Code generated by ent, DO NOT EDIT.

package thumbmarkidblacklist

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/stablecog/sc-go/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldLTE(FieldID, id))
}

// ThumbmarkID applies equality check predicate on the "thumbmark_id" field. It's identical to ThumbmarkIDEQ.
func ThumbmarkID(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldEQ(FieldThumbmarkID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldEQ(FieldUpdatedAt, v))
}

// ThumbmarkIDEQ applies the EQ predicate on the "thumbmark_id" field.
func ThumbmarkIDEQ(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldEQ(FieldThumbmarkID, v))
}

// ThumbmarkIDNEQ applies the NEQ predicate on the "thumbmark_id" field.
func ThumbmarkIDNEQ(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldNEQ(FieldThumbmarkID, v))
}

// ThumbmarkIDIn applies the In predicate on the "thumbmark_id" field.
func ThumbmarkIDIn(vs ...string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldIn(FieldThumbmarkID, vs...))
}

// ThumbmarkIDNotIn applies the NotIn predicate on the "thumbmark_id" field.
func ThumbmarkIDNotIn(vs ...string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldNotIn(FieldThumbmarkID, vs...))
}

// ThumbmarkIDGT applies the GT predicate on the "thumbmark_id" field.
func ThumbmarkIDGT(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldGT(FieldThumbmarkID, v))
}

// ThumbmarkIDGTE applies the GTE predicate on the "thumbmark_id" field.
func ThumbmarkIDGTE(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldGTE(FieldThumbmarkID, v))
}

// ThumbmarkIDLT applies the LT predicate on the "thumbmark_id" field.
func ThumbmarkIDLT(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldLT(FieldThumbmarkID, v))
}

// ThumbmarkIDLTE applies the LTE predicate on the "thumbmark_id" field.
func ThumbmarkIDLTE(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldLTE(FieldThumbmarkID, v))
}

// ThumbmarkIDContains applies the Contains predicate on the "thumbmark_id" field.
func ThumbmarkIDContains(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldContains(FieldThumbmarkID, v))
}

// ThumbmarkIDHasPrefix applies the HasPrefix predicate on the "thumbmark_id" field.
func ThumbmarkIDHasPrefix(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldHasPrefix(FieldThumbmarkID, v))
}

// ThumbmarkIDHasSuffix applies the HasSuffix predicate on the "thumbmark_id" field.
func ThumbmarkIDHasSuffix(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldHasSuffix(FieldThumbmarkID, v))
}

// ThumbmarkIDEqualFold applies the EqualFold predicate on the "thumbmark_id" field.
func ThumbmarkIDEqualFold(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldEqualFold(FieldThumbmarkID, v))
}

// ThumbmarkIDContainsFold applies the ContainsFold predicate on the "thumbmark_id" field.
func ThumbmarkIDContainsFold(v string) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldContainsFold(FieldThumbmarkID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ThumbmarkIdBlackList) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ThumbmarkIdBlackList) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(func(s *sql.Selector) {
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
func Not(p predicate.ThumbmarkIdBlackList) predicate.ThumbmarkIdBlackList {
	return predicate.ThumbmarkIdBlackList(func(s *sql.Selector) {
		p(s.Not())
	})
}