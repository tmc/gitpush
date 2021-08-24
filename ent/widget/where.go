// Code generated by entc, DO NOT EDIT.

package widget

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/tmc/moderncrud/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriority), v))
	})
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNote), v))
	})
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.Widget {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Widget(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNote), v...))
	})
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.Widget {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Widget(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNote), v...))
	})
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNote), v))
	})
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNote), v))
	})
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNote), v))
	})
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNote), v))
	})
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNote), v))
	})
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNote), v))
	})
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNote), v))
	})
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNote), v))
	})
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNote), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Widget {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Widget(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Widget {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Widget(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Widget {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Widget(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Widget {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Widget(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPriority), v))
	})
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPriority), v))
	})
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...int) predicate.Widget {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Widget(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPriority), v...))
	})
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...int) predicate.Widget {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Widget(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPriority), v...))
	})
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPriority), v))
	})
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPriority), v))
	})
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPriority), v))
	})
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v int) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPriority), v))
	})
}

// HasType applies the HasEdge predicate on the "type" edge.
func HasType() predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TypeTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TypeTable, TypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTypeWith applies the HasEdge predicate on the "type" edge with a given conditions (other predicates).
func HasTypeWith(preds ...predicate.WidgetType) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TypeInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TypeTable, TypeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Widget) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Widget) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
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
func Not(p predicate.Widget) predicate.Widget {
	return predicate.Widget(func(s *sql.Selector) {
		p(s.Not())
	})
}
