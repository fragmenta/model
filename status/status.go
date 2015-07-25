package status


import (
     "github.com/fragmenta/query"
     "github.com/fragmenta/view/helpers"
)

type ModelStatus struct {
	Status int64
}

// Status values
// If these need to be substantially modified for a particular model, 
// it may be better to move this into the model package concerned and modify as required
const (
    Draft = 0
    Final = 10
    Suspended = 11
    Unavailable = 12
    Published = 100
    Featured = 101
)



// Return an array of statuses for a status select
func (m *ModelStatus) StatusOptions() []helpers.Option {
	options := make([]helpers.Option, 0)
	
    options = append(options,helpers.Option{Draft,"Draft"})
    options = append(options,helpers.Option{Final,"Final"})
    options = append(options,helpers.Option{Suspended,"Suspended"})
    options = append(options,helpers.Option{Published,"Published"})
  
    // For now leave featured off default statuses
   // options = append(options,helpers.Option{Featured,"Featured"})
    
    return options
}

// Return the string representation of the model status
func (m *ModelStatus) StatusDisplay() string {
    for _,o := range m.StatusOptions() {
        if o.Id == m.Status {
            return o.Name
        }
    }
    return ""
}

// Model status

// Is the status Draft?
func (m *ModelStatus) IsDraft() bool {
	return m.Status == Draft
}

// Is the status Final?
func (m *ModelStatus) IsFinal() bool {
	return m.Status == Final
}

// Is the status Suspended?
func (m *ModelStatus) IsSuspended() bool {
	return m.Status == Suspended
}

// Is the status Unavailable?
func (m *ModelStatus) IsUnavailable() bool {
	return m.Status == Unavailable
}

// Is the status Published?
// Anything over Published status is Published
func (m *ModelStatus) IsPublished() bool {
	return m.Status >= Published // NB >=
}

// Is the status Featured?
func (m *ModelStatus) IsFeatured() bool {
	return m.Status == Featured
}



// CHAINABLE FINDER FUNCTIONS
// Apply with query.Apply(status.WherePublished) etc
// Or define on your own models instead...


// Modify the given query to select status draft
func WhereDraft(q *query.Query) *query.Query {
    return q.Where("status = ?", Draft)
}

// Modify the given query to select status Final 
func WhereFinal(q *query.Query) *query.Query {
    return q.Where("status = ?", Final)
}

// Modify the given query to select status Suspended 
func WhereSuspended(q *query.Query) *query.Query {
    return q.Where("status = ?", Suspended)
}

// Modify the given query to select status Featured 
func WhereFeatured(q *query.Query) *query.Query {
    return q.Where("status = ?", Featured)
}

// Modify the given query to select status Published 
func WherePublished(q *query.Query) *query.Query {
    return q.Where("status >= ?", Published)
}

// Modify the given query to select records with null status
func Null(q *query.Query) *query.Query {
    return q.Where("status IS NULL")
}

// Modify the given query to select records which do not have null status
func NotNull(q *query.Query) *query.Query {
    return q.Where("status IS NOT NULL")
}

// Modify the given query to order records by status
func Order(q *query.Query) *query.Query {
    return q.Order("status desc")
}
