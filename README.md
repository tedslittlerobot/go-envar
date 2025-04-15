Envar for Golang
================

Easily assign configuration variables from the environment or elsewhere.

```go
package main

import (
	"github.com/tedslittlerobot/go-envar"
	"log"
)

type MyVariables struct {
	SqlHost     string `envar:"env:MYSQL_HOST,default:localhost"`
	SqlPort     int    `envar:"env:MYSQL_PORT,default:3306"`
	SqlDatabase string `envar:"env:MYSQL_DB_NAME"`
}

func main() {
	variables := MyVariables{}
	
	e := envar.MakeWithDefaults()

	e.Apply(&variables)

	log.Printf("Host: %s, Port: %d, Database: %s", variables.SqlHost, variables.SqlPort, variables.SqlDatabase)
}
```

## What problem does this solve

There are many popular environment variable resolvers for Golang. However, the most they do is:

- Get a value from a single Environment Variable
- Allow the providing of a default value

This is more than sufficient for most use cases. There are some cases where a project dictates that your configuration variables may be coming from more complex places.

## Fallback Chain

```go
package main

type MyVariables struct {
	SqlDatabase string `envar:"env:MYSQL_DB_NAME,env:DATABASE,env:MYSQL_LOCAL_DB_NAME"`
}
```

Here, you can see that many environment variables can be attempted in order. The first non-empty value will be allocated to the struct property.

## Errors and Handling

All errors in envar are sent as `panic` calls, rather than bubbling them up as error responses. This decision has been made as typically these errors are all non-recoverable, and should be encountered in the setup part of an application (ie. before the bulk of the application code), and there should be very little by way of valid recovering from such errors.

Regardless, you can always `recover` from a panic if error handling is required.

## Resolvers

### Basic Resolvers

The resolvers you have seen in the examples above are `EnvironmentVariableResolver` (named `env`) `RawValueResolver` (named `default`) respectively.

- `env:MY_VARIABLE` The `EnvironmentVariableResolver` will use a call to `os.Getenv("MY_VARIABLE")` to get its value.
- `default:foobar` The `RawValueResolver` will always return the value you provide after the colon. This has been nicknamed to `default` as that would typically be how this would be used.

These resolvers will be registered automatically for you.

### Custom Resolvers

Envar has a driver-based resolver solution, so you can add your own resolvers very easily. You simply need to implement the `ResolverInterface` interface like so:

```go
package main

import "github.com/tedslittlerobot/go-envar"

type AlwaysReturnPeterResolver struct{}

// Resolve is responsible for calling Resolve on the given token with a value. It should call resolve regardless of 
func (resolver AlwaysReturnPeterResolver) Resolve(token *envar.SourceToken) {
	// token.Key will be set to the value after the colon
	token.Resolve("Peter")
}

// PreLoad enables you to batch retrieve environment variables in one go. Useful if you are getting variables via a network request of some kind
func (resolver AlwaysReturnPeterResolver) PreLoad(tokens []*envar.SourceToken) {}
```

Then you can register your resolvers on the Config struct when you call Envar:

```go
package main

import (
	"github.com/tedslittlerobot/go-envar"
)

type MyVariables struct {
	Peter       string `envar:"peter:"`
	SqlHost     string `envar:"env:MYSQL_HOST,default:localhost"`
	SqlPort     int    `envar:"env:MYSQL_PORT,default:3306"`
	SqlDatabase string `envar:"env:MYSQL_DB_NAME"`
}

func main() {
	variables := MyVariables{}

	envar.Apply(&variables, envar.Config{
		Resolvers: map[string]envar.ResolverInterface{
			"peter": AlwaysReturnPeterResolver{},
		},
		// You can specify false to not use the default resolvers (env and default). 
		//If this is true, this will merge your provided resolvers on top of the internal map.
		true, 
    })
}
```
