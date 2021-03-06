// Code generated by entc, DO NOT EDIT.

package token

import (
	"github.com/andreylm/mind_systems/pkg/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Token {
	return predicate.Token(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
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
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
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
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// URL applies equality check predicate on the "url" field. It's identical to URLEQ.
func URL(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	},
	)
}

// URLShortened applies equality check predicate on the "url_shortened" field. It's identical to URLShortenedEQ.
func URLShortened(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURLShortened), v))
	},
	)
}

// URLEQ applies the EQ predicate on the "url" field.
func URLEQ(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURL), v))
	},
	)
}

// URLNEQ applies the NEQ predicate on the "url" field.
func URLNEQ(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURL), v))
	},
	)
}

// URLIn applies the In predicate on the "url" field.
func URLIn(vs ...string) predicate.Token {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Token(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURL), v...))
	},
	)
}

// URLNotIn applies the NotIn predicate on the "url" field.
func URLNotIn(vs ...string) predicate.Token {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Token(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURL), v...))
	},
	)
}

// URLGT applies the GT predicate on the "url" field.
func URLGT(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURL), v))
	},
	)
}

// URLGTE applies the GTE predicate on the "url" field.
func URLGTE(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURL), v))
	},
	)
}

// URLLT applies the LT predicate on the "url" field.
func URLLT(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURL), v))
	},
	)
}

// URLLTE applies the LTE predicate on the "url" field.
func URLLTE(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURL), v))
	},
	)
}

// URLContains applies the Contains predicate on the "url" field.
func URLContains(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURL), v))
	},
	)
}

// URLHasPrefix applies the HasPrefix predicate on the "url" field.
func URLHasPrefix(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURL), v))
	},
	)
}

// URLHasSuffix applies the HasSuffix predicate on the "url" field.
func URLHasSuffix(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURL), v))
	},
	)
}

// URLEqualFold applies the EqualFold predicate on the "url" field.
func URLEqualFold(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURL), v))
	},
	)
}

// URLContainsFold applies the ContainsFold predicate on the "url" field.
func URLContainsFold(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURL), v))
	},
	)
}

// URLShortenedEQ applies the EQ predicate on the "url_shortened" field.
func URLShortenedEQ(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldURLShortened), v))
	},
	)
}

// URLShortenedNEQ applies the NEQ predicate on the "url_shortened" field.
func URLShortenedNEQ(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldURLShortened), v))
	},
	)
}

// URLShortenedIn applies the In predicate on the "url_shortened" field.
func URLShortenedIn(vs ...string) predicate.Token {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Token(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldURLShortened), v...))
	},
	)
}

// URLShortenedNotIn applies the NotIn predicate on the "url_shortened" field.
func URLShortenedNotIn(vs ...string) predicate.Token {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Token(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldURLShortened), v...))
	},
	)
}

// URLShortenedGT applies the GT predicate on the "url_shortened" field.
func URLShortenedGT(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldURLShortened), v))
	},
	)
}

// URLShortenedGTE applies the GTE predicate on the "url_shortened" field.
func URLShortenedGTE(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldURLShortened), v))
	},
	)
}

// URLShortenedLT applies the LT predicate on the "url_shortened" field.
func URLShortenedLT(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldURLShortened), v))
	},
	)
}

// URLShortenedLTE applies the LTE predicate on the "url_shortened" field.
func URLShortenedLTE(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldURLShortened), v))
	},
	)
}

// URLShortenedContains applies the Contains predicate on the "url_shortened" field.
func URLShortenedContains(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldURLShortened), v))
	},
	)
}

// URLShortenedHasPrefix applies the HasPrefix predicate on the "url_shortened" field.
func URLShortenedHasPrefix(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldURLShortened), v))
	},
	)
}

// URLShortenedHasSuffix applies the HasSuffix predicate on the "url_shortened" field.
func URLShortenedHasSuffix(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldURLShortened), v))
	},
	)
}

// URLShortenedEqualFold applies the EqualFold predicate on the "url_shortened" field.
func URLShortenedEqualFold(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldURLShortened), v))
	},
	)
}

// URLShortenedContainsFold applies the ContainsFold predicate on the "url_shortened" field.
func URLShortenedContainsFold(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldURLShortened), v))
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Token) predicate.Token {
	return predicate.Token(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
