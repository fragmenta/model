// role.Role adds a status field and associated functions to a Role
package role

import (
         "github.com/fragmenta/query"
         "github.com/fragmenta/view/helpers"
)

// Constants for this package - usage role.Reader
const (
    Anon = 0
    Reader = 10
    Customer = 20
    Editor = 30
    Author = 40
    Owner = 50
    Admin = 100
)


// The role Role - usage embed *role.ModelRole to add roles
// or just role your own if you have specific requirements
type ModelRole struct {
	Role int64
}

func (m *ModelRole) RoleValue() int64 {
	return m.Role
}

func (m *ModelRole) Anon() bool {
	return m.Role == Anon
}

func (m *ModelRole) Reader() bool {
	return m.Role == Reader
}

func (m *ModelRole) Customer() bool {
	return m.Role == Customer
}

func (m *ModelRole) Editor() bool {
	return m.Role == Editor
}

func (m *ModelRole) Author() bool {
	return m.Role == Author
}

func (m *ModelRole) Owner() bool {
	return m.Role == Owner
}

func (m *ModelRole) Admin() bool {
	return m.Role == Admin
}

// Return an array of Role values for this model (embedders may override this and roledisplay to extend)
func (m *ModelRole) RoleOptions() []helpers.Option {
	options := make([]helpers.Option, 0)
	
    options = append(options,helpers.Option{Reader,"Reader"})
    options = append(options,helpers.Option{Customer,"Customer"})
    options = append(options,helpers.Option{Editor,"Editor"})
    options = append(options,helpers.Option{Author,"Author"})
    options = append(options,helpers.Option{Owner,"Owner"})
    options = append(options,helpers.Option{Admin,"Administrator"})
    
    return options
}





// Return the string representation of the Role status 
func (m *ModelRole) RoleDisplay() string {
        for _,o := range m.RoleOptions() {
            if o.Id == m.Role {
                return o.Name
            }
        }
        return ""
}


// Modify the given query to select role anon
func WhereAnon(q *query.Query) *query.Query {
    return q.Where("role = ?", Anon)
}

// Modify the given query to select role Reader
func WhereReader(q *query.Query) *query.Query {
    return q.Where("role = ?", Reader)
}

// Modify the given query to select role Customer
func WhereCustomer(q *query.Query) *query.Query {
    return q.Where("role = ?", Customer)
}

// Modify the given query to select role Editor
func WhereEditor(q *query.Query) *query.Query {
    return q.Where("role = ?", Editor)
}

// Modify the given query to select role Author
func WhereAuthor(q *query.Query) *query.Query {
    return q.Where("role = ?", Author)
}

// Modify the given query to select role Admin
func WhereAdmin(q *query.Query) *query.Query {
    return q.Where("role = ?", Admin)
}

// Modify the given query to select records with null role
func Null(q *query.Query) *query.Query {
    return q.Where("role IS NULL")
}

// Modify the given query to select records which do not have null role
func NotNull(q *query.Query) *query.Query {
    return q.Where("role IS NOT NULL")
}

// Modify the given query to order records by role
func Order(q *query.Query) *query.Query {
    return q.Order("role desc")
}
