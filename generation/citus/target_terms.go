package citus

import (
	"fmt"

	"github.com/jackc/pgx"

	"github.com/unchartedsoftware/prism/tile"
)

type TargetTerms struct {
	tile.TargetTerms
}

func (t *TargetTerms) AddQuery(query *Query) *Query {
	//Want to keep only documents that have the specified terms.
	//Use the already existing Has construct.
	hasQuery := &Has{}
	hasQuery.Field = t.TermsField
	terms := make([]interface{}, len(t.Terms))
	for i, term := range t.Terms {
		terms[i] = term
	}
	hasQuery.Values = terms

	clause, _ := hasQuery.Get(query)
	query.Where(clause)
	return query
}

func (t *TargetTerms) AddAggs(query *Query) *Query {
	//Count by term, only considering the specified terms.
	//TODO: Find a better way to make this work. The caller NEEDS to use the returned value.
	//Assume the backing field is an array. Need to unpack that array and group by the terms.
	query.Select(fmt.Sprintf("unnest(%s) AS term", t.TermsField))

	//Need to nest the existing query as a table and group by the terms.
	//TODO: Figure out how to handle error properly.
	termQuery, _ := NewQuery()

	termQuery.From(fmt.Sprintf("(%s) terms", query.GetQuery(true)))
	termQuery.GroupBy("term")
	termQuery.Select("term")
	termQuery.Select("COUNT(*) as term_count")
	termQuery.OrderBy("term_count desc")

	//Generate the filter for the terms.
	clause := ""
	for _, value := range t.Terms {
		valueParam := query.AddParameter(value)
		clause = clause + fmt.Sprintf(", %s", valueParam)
	}
	termQuery.Where(fmt.Sprintf("term IN [%s]", clause[2:]))

	return termQuery
}

// GetTerms parses the result of the terms query into a map of term -> count.
func (t *TargetTerms) GetTerms(rows *pgx.Rows) (map[string]uint32, error) {
	// build map of topics and counts
	counts := make(map[string]uint32)
	for rows.Next() {
		var term string
		var term_count uint32
		err := rows.Scan(&term, &term_count)
		if err != nil {
			return nil, fmt.Errorf("Error parsing top terms: %v", err)
		}
		counts[term] = term_count
	}
	return counts, nil
}