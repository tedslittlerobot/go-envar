Envar for Golang
================

Easily assign configuration variables from the environment or elsewhere.

**WARNING - STILL VERY MUCH IN DEVELOPMENT - API MAY CHANGE**

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
	
	e := envar.Make()

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

### Configurable Resolvers

#### MapResolver

The `MapResolver` struct allows you to specify an arbitrary map of data (perhaps retrieved from an external source!) to use as a data source.

```go
package main

import (
	"github.com/tedslittlerobot/go-envar"
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
	"log"
)

type MyVariables struct {
	value   string `envar:"my-map:FOO"`
	monkeys string `envar:"my-map:MONKEYS,default:no-monkeys"`
}

func main() {
	variables := MyVariables{}
	
	e := envar.Make()

	e.Resolvers.AddResolver("my-map", envarResolvers.MapResolver{
		Contents: map[string]string{ // This could be an external call
			"FOO": "foo",
			"BAR": "bar",
			"BAZ": "baz",
		},
	})

	e.Apply(&variables)

	log.Printf("value: %s, Monkeys: %s", variables.value, variables.monkeys)
	// Prints "value: foo, Monkeys: no-monkeys"
}
```

#### AWS Parameter Store Resolver


There is a helper function to use the current AWS context to fill the `MapResolver` with all SSM parameter store variables from a given path prefix.

There are two helper functions - one which gets you provide a path prefix, and one to retrieve the path prefix from an environment variable. 
The latter is likely used if you have prefixed the parameter store paths based on say, environment, or tenant. 

```go
package main

import (
	"context"
	"github.com/tedslittlerobot/go-envar"
	envarResolvers "github.com/tedslittlerobot/go-envar/support/resolvers"
	"log"
)

type MyVariables struct {
	value   string `envar:"from-ssm:value"`
	monkeys string `envar:"from-ssm:value/at/path,default:no-monkeys"`
}

func main() {
	variables := MyVariables{}
	e := envar.Make()

	e.AddSsmResolver(context.TODO(), "from-ssm", "/prod/", true)
	e.AddSsmResolverFromEnv(context.TODO(), "from-ssm-2", "SSM_PREFIX_ENVIRONMENT_VARIABLE", true)

	e.Apply(&variables)
}
```

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
```

Then you can register your resolvers on your Envar instance:

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
	e := envar.Make()
	
	e.Resolvers.AddResolver("peter", AlwaysReturnPeterResolver{})

	e.Apply(&variables)
}
```
