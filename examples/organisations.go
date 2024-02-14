package examples

import (
	"fmt"
	"github.com/mr-destructive/turso-go"
)

func main() {
	client, _ := turso.NewClient("", "YOUR-AUTH-TOKEN")
	tokens, _ := client.Tokens.List()
	fmt.Println(tokens)

	orgs, _ := client.Organizations.List()
	fmt.Println(orgs)

	members, _ := client.Organizations.Members("org_slug")
	fmt.Println(members)
}
