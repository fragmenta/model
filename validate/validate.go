// Package validate provides methods for validating params passing to and from the database
// It does not deal with params which come in from the web
package validate

import (
	"errors"
	"fmt"
	"time"
)



// Various methods for casting params back to real values from interface{} 
// (as they come in from the db drivers)

// Validate a param exists and return float value
func Float(param interface{}) float64 {
	var v float64
	if param != nil {
		v = param.(float64)
	}
	return v
}

func Boolean(param interface{}) bool {
	var v bool
	if param != nil {
		v = param.(bool)
	}
	return v
}


// Validate a param exists and return int value
func Int(param interface{}) int64 {
	var v int64
	if param != nil {
		v = param.(int64)
	}
	return v
}

// Validate a param exists and return string value
func String(param interface{}) string {
	var v string
	if param != nil {
		v = param.(string)
	}
	return v
}

// Validate a param exists and return time value
// This could result in invalid times being assigned... hmm
func Time(param interface{}) time.Time {
	var v time.Time
	if param != nil {
		v = param.(time.Time)
	}
	return v
}

func Length(param string, min int, max int) error {
	length := len(param)
	if min != -1 && length < min {
		return errors.New(fmt.Sprintf("Length of string %s %d, expected > %d", param, length, min))
	}
	if max != -1 && length > max {
		return errors.New(fmt.Sprintf("Length of string %s %d, expected < %d", param, length, max))
	}
	return nil
}

// Filter the param names given, and return filtered params.
// Actions might use CleanParams to restrict params depending on auth status for example.
// Usage: params = CleanParams(params, []string{"id", "title"})
func CleanParams(params map[string]string, allowed []string) (map[string]string,error) {

    // Delete redirect param - perhaps do this unless explicitly included in allowed?
    delete(params, "redirect")
   
	// We have some valid params, so filter params so they only include those ones
	// NB order doesn't matter here
    var err error 
    INVALID := "FRAGMENTA_INVALID_PARAM"
    
	validated := map[string]string{}
	for k, v := range params {
        validated[k] = INVALID
        
		for _, validKey := range allowed {
			if k == validKey {
				validated[k] = v
			} 
		}
        
        if (validated[k] == INVALID) {
          err = fmt.Errorf("Invalid param found:%s\n Params:%v",k,params);   
          break;
        }
	}


	return validated,err
}
