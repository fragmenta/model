// Package model provides a class including Id, CreatedAt and UpdatedAt, and some utility functions, optionally include in your models
package model

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Model base class
type Model struct {
	// FIXME - are TableName and KeyName required?
	TableName string
	KeyName   string
	Id        int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Init sets up the model fields
func (m *Model) Init() {
	m.TableName = ""
	m.KeyName = "id"
	m.Id = 0
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
}

// UrlPrefix returns the url prefix for this model (normally the table name)
func (m *Model) UrlPrefix() string {
	return m.TableName
}

func (m *Model) UrlCreate() string {
	return fmt.Sprintf("/%s/%d/create", m.TableName, m.Id)
}

func (m *Model) UrlUpdate() string {
	return fmt.Sprintf("/%s/%d/update", m.TableName, m.Id)
}

func (m *Model) UrlDestroy() string {
	return fmt.Sprintf("/%s/%d/destroy", m.TableName, m.Id)
}

// This should use ToSlug to create a slug from name as well...
func (m *Model) UrlShow() string {
	return fmt.Sprintf("/%s/%d", m.TableName, m.Id)
}

func (m *Model) UrlIndex() string {
	return fmt.Sprintf("/%s", m.TableName, m.Id)
}

// Convert a file name to something suitable for use on the web as part of a url
func ToSlug(name string) string {
	// Lowercase
	slug := strings.ToLower(name)

	// Replace _ with - for consistent style
	slug = strings.Replace(slug, "_", "-", -1)
	slug = strings.Replace(slug, " ", "-", -1)

	// In case of regexp failure, replace at least /
	slug = strings.Replace(slug, "/", "-", -1)

	// Run regexp - remove all letters except a-z0-9-
	re, err := regexp.Compile("[^a-z0-9-]*")
	if err != nil {
		fmt.Println("ToSlug regexp failed")
	} else {
		slug = re.ReplaceAllString(slug, "")
	}

	return slug
}

// Return the table name for this object
func (m *Model) Table() string {
	return m.TableName
}

// Use id for primary key by default - used by query
func (m *Model) PrimaryKey() string {
	return m.KeyName
}

// What name should we return for select options?
func (m *Model) SelectName() string {
	return fmt.Sprintf("%s-%d", m.TableName, m.Id) // Usually override with name or a summary
}

// What value should we return for select options?
func (m *Model) SelectValue() string {
	return fmt.Sprintf("%d", m.Id)
}

// PrimaryKeyValue returns the unique id
func (m *Model) PrimaryKeyValue() int64 {
	return m.Id
}

// OwnedBy returns true if the user id passed in owns this model
func (m *Model) OwnedBy(uid int64) bool {
	// In models composed with base model, you may want to check a user_id field or join table
	// In this base model, we return false by default
	return false
}

// CacheKey generates a cache key for this model object, dependent on id and UpdatedAt
// should we generate a hash of this to ensure we fit in small key size?
func (m *Model) CacheKey() string {
	// This should really be some form of hash based on this data...
	return fmt.Sprintf("%s/%d/%s", m.TableName, m.Id, m.UpdatedAt)
}

// String returns a string representation of the model
func (m *Model) String() string {
	return fmt.Sprintf("%s/%d", m.TableName, m.Id)
}
