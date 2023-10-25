## Turso Golang Client

This is a Golang client for interacting with the Turso API. It provides methods for:

- [x] Managing API tokens
- [ ] Tracking Database Usage
- [ ] Managing Logical DBs and Instances

## Installation

```bash
go get github.com/mr-destructive/turso-go
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/mr-destructive/turso-go"
)

func main() {
    client := turso.NewClient("", "YOUR_API_TOKEN")
    tokens, err := client.Tokens.List()
    if err != nil {
        panic(err)
    }
    fmt.Println(tokens)
}
```

- Create a new client, first parameter is the Turso API base URL (for the API incase of self hosted), leave empty if using turso.
- Provide the API token by logging in to the CLI

### Authentication

The client must be created with a valid API token. Get an API token from the [Turso CLI](https://docs.turso.tech/reference/turso-cli).

- The token will be automatically included in all requests to the API.

### Organizations

- Get all the organisations for the authenticated user:

```go
client := turso.NewClient("", "YOUR_API_TOKEN")

orgs, err := client.Organizations.List()

if err != nil {
    panic(err)
}
fmt.Println(orgs)
```

- Get members with their roles in a given organisation:

```go
client := turso.NewClient("", "YOUR_API_TOKEN")

// provide the org slug
members, err := client.Organizations.Members("org_slug")

if err != nil {
    panic(err)
}
fmt.Println(orgs)
```

### Instances


## References

- [Turso Platform REST API docs](https://docs.turso.tech/reference/platform-rest-api/)
- [Turso CLI](https://docs.turso.tech/reference/turso-cli)
