// Code generated by ent, DO NOT EDIT.

package deviceinfo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/stablecog/go-apps/database/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldID, id))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldType, v))
}

// Os applies equality check predicate on the "os" field. It's identical to OsEQ.
func Os(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldOs, v))
}

// Browser applies equality check predicate on the "browser" field. It's identical to BrowserEQ.
func Browser(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldBrowser, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContainsFold(FieldType, v))
}

// OsEQ applies the EQ predicate on the "os" field.
func OsEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldOs, v))
}

// OsNEQ applies the NEQ predicate on the "os" field.
func OsNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldOs, v))
}

// OsIn applies the In predicate on the "os" field.
func OsIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldOs, vs...))
}

// OsNotIn applies the NotIn predicate on the "os" field.
func OsNotIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldOs, vs...))
}

// OsGT applies the GT predicate on the "os" field.
func OsGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldOs, v))
}

// OsGTE applies the GTE predicate on the "os" field.
func OsGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldOs, v))
}

// OsLT applies the LT predicate on the "os" field.
func OsLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldOs, v))
}

// OsLTE applies the LTE predicate on the "os" field.
func OsLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldOs, v))
}

// OsContains applies the Contains predicate on the "os" field.
func OsContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContains(FieldOs, v))
}

// OsHasPrefix applies the HasPrefix predicate on the "os" field.
func OsHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasPrefix(FieldOs, v))
}

// OsHasSuffix applies the HasSuffix predicate on the "os" field.
func OsHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasSuffix(FieldOs, v))
}

// OsEqualFold applies the EqualFold predicate on the "os" field.
func OsEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEqualFold(FieldOs, v))
}

// OsContainsFold applies the ContainsFold predicate on the "os" field.
func OsContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContainsFold(FieldOs, v))
}

// BrowserEQ applies the EQ predicate on the "browser" field.
func BrowserEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldBrowser, v))
}

// BrowserNEQ applies the NEQ predicate on the "browser" field.
func BrowserNEQ(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldBrowser, v))
}

// BrowserIn applies the In predicate on the "browser" field.
func BrowserIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldBrowser, vs...))
}

// BrowserNotIn applies the NotIn predicate on the "browser" field.
func BrowserNotIn(vs ...string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldBrowser, vs...))
}

// BrowserGT applies the GT predicate on the "browser" field.
func BrowserGT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldBrowser, v))
}

// BrowserGTE applies the GTE predicate on the "browser" field.
func BrowserGTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldBrowser, v))
}

// BrowserLT applies the LT predicate on the "browser" field.
func BrowserLT(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldBrowser, v))
}

// BrowserLTE applies the LTE predicate on the "browser" field.
func BrowserLTE(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldBrowser, v))
}

// BrowserContains applies the Contains predicate on the "browser" field.
func BrowserContains(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContains(FieldBrowser, v))
}

// BrowserHasPrefix applies the HasPrefix predicate on the "browser" field.
func BrowserHasPrefix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasPrefix(FieldBrowser, v))
}

// BrowserHasSuffix applies the HasSuffix predicate on the "browser" field.
func BrowserHasSuffix(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldHasSuffix(FieldBrowser, v))
}

// BrowserEqualFold applies the EqualFold predicate on the "browser" field.
func BrowserEqualFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEqualFold(FieldBrowser, v))
}

// BrowserContainsFold applies the ContainsFold predicate on the "browser" field.
func BrowserContainsFold(v string) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldContainsFold(FieldBrowser, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.DeviceInfo {
	return predicate.DeviceInfo(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasGenerations applies the HasEdge predicate on the "generations" edge.
func HasGenerations() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GenerationsTable, GenerationsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenerationsWith applies the HasEdge predicate on the "generations" edge with a given conditions (other predicates).
func HasGenerationsWith(preds ...predicate.Generation) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenerationsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GenerationsTable, GenerationsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUpscales applies the HasEdge predicate on the "upscales" edge.
func HasUpscales() predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UpscalesTable, UpscalesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUpscalesWith applies the HasEdge predicate on the "upscales" edge with a given conditions (other predicates).
func HasUpscalesWith(preds ...predicate.Upscale) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UpscalesInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, UpscalesTable, UpscalesColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
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
func Not(p predicate.DeviceInfo) predicate.DeviceInfo {
	return predicate.DeviceInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}
