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

	envar.Envar(&variables, envar.Config{})

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

## Fatal Erroring

If no value can be found in the chain, then Envar will throw a fatal error to crash the application, rather than offloading the error handling logic to the parent application.
