package examples

import (
	"fmt"
	"github.com/mr-destructive/turso-go"
)

func main() {
	client, _ := turso.NewClient("", "YOUR-AUTH-TOKEN")
	tokens, _ := client.Tokens.List()
	fmt.Println(tokens)

	mint, _ := client.Tokens.Mint("test")
	fmt.Println(mint)

	validate, _ := client.Tokens.Validate()
	fmt.Println(validate)
}
