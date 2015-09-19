# Package model
An optional base model class and utilities. The package includes some basic fields and methods acting on those fields. The fields included are:

```Go 
    type Model struct {
    	TableName string
    	KeyName   string
    	Id        int64
    	CreatedAt time.Time
    	UpdatedAt time.Time
    }
```

### Usage 

It can be included in models with:

```Go 
type MyModel struct {
    model.Model
    ...
```

Utility subpackages are:

* file - utilities for handling files
* validate - utilities for validating field values
