# mapdb

A simple database using `map` go type, serialized to file.

## Instalation

```
$ go get github.com/erichnascimento/mapdb
```

## Example
```go
package main

import "fmt"
import "github.com/erichnascimento/mapdb"

func main() {

  // open db, loading from file if exists
  db := mapdb.OpenDB("/tmp/mydb.db", nil)

  defer db.Close()

  // Set a key
  db.Set("username", "erichnascimento")
  db.Set("password", "1234")

  // Get a key value
  fmt.Println(db.Get("username"))
  fmt.Println(db.Get("password"))

  // Update a key
  db.Set("username", "other-username")
  fmt.Println(db.Get("username"))

  // Delete a key
  db.Del("password")
  fmt.Println(db.Get("password"))

  // Set a json string
  db.Set("preferences", `{"foo": "bar"}`)
  fmt.Println(db.Get("preferences"))

  // Print all keys
  fmt.Println(db.Keys())

  // Save to file
  db.Save()
}
```

If you run this example, you will see:
```
$ go run example/simple.go 
erichnascimento
1234
other-username
<nil>
{"foo": "bar"}
[preferences username]
```

# License
MIT