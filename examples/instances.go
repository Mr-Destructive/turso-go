package examples

import (
	"fmt"
	"github.com/mr-destructive/turso-go"
)

func main() {
	client, _ := turso.NewClient("", "YOUR-AUTH-TOKEN")
	tokens, _ := client.Tokens.List()
	fmt.Println(tokens)

	instances, _ := client.Organizations.Instances("org_slug", "db_name")
	fmt.Println(instances)

    instance, _ := client.Organizations.Instance("org_slug", "db_name", "instance_name")
    fmt.Println(instance)
}
