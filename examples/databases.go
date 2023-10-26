package examples

import (
	"fmt"
	"github.com/mr-destructive/turso-go"
)

func main() {
	client, _ := turso.NewClient("", "YOUR-AUTH-TOKEN")
	tokens, _ := client.Tokens.List()
	fmt.Println(tokens)

	dbs, _ := client.Organizations.Databases("org_slug")
	fmt.Println(dbs)

	db, _ := client.Organizations.Database("org_slug", "db_name")
	fmt.Println(db)

    usage, _ := client.Organizations.DBUsage("org_slug", "db_name")
    fmt.Println(usage)
}
