package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
)

type Data map[string]interface{}

func ReadSecret(client *api.Client, path string) Data {
	secret, err := client.Logical().Read(path)
	if err != nil {
		fmt.Println("Failed to read secret")
		fmt.Println(err)
		os.Exit(2)
	}

	result := make(Data)
	if secret != nil {
		for key, value := range secret.Data {
			result[key] = value
		}
	}
	return result
}

func WriteSecret(client *api.Client, path string, data Data) {
	_, err := client.Logical().Write(path, data)
	if err != nil {
		fmt.Println("Failed to read secret")
		fmt.Println(err)
		os.Exit(2)
	}
}

func main() {
	deleteFlag := flag.Bool("delete", false, "Delete the key instead of updating it")
	flag.Parse()

	args := flag.Args()

	if len(args) < 2 {
		fmt.Println(fmt.Sprintf("usage %s path key=value | %s path --delete key", os.Args[0], os.Args[0]))
		os.Exit(1)
	}
	path := args[0]

	config := api.DefaultConfig()

	client, err := api.NewClient(config)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	secret := ReadSecret(client, path)
	if *deleteFlag {
		delete(secret, args[1])
	} else {
		parts := strings.SplitN(args[1], "=", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format, expected key=value")
			os.Exit(1)
		}
		secret[parts[0]] = parts[1]
	}
	WriteSecret(client, path, secret)
}
