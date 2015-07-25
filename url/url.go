// FIXME - remove this package and all refs - now put into the base model
package url

import (
	"fmt"
	"regexp"
	"strings"
)

type Resource interface {
	// The url prefix used for the resource (typically the same as the Table())
	UrlPrefix() string

	// The primary key value as a string
	PrimaryKeyValue() int64
}

// Generic REST url handlers, which are passed a base model

func Index(m Resource) string {
	return fmt.Sprintf("/%s", m.UrlPrefix())
}

func Create(m Resource) string {
	return fmt.Sprintf("/%s/create", m.UrlPrefix())
}

func Show(m Resource) string {
	return fmt.Sprintf("/%s/%d", m.UrlPrefix(), m.PrimaryKeyValue())
}

func Update(m Resource) string {
	return fmt.Sprintf("/%s/%d/update", m.UrlPrefix(), m.PrimaryKeyValue())
}

func Destroy(m Resource) string {
	return fmt.Sprintf("/%s/%d/destroy", m.UrlPrefix(), m.PrimaryKeyValue())
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
