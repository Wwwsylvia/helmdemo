package main

import (
	"fmt"

	"helm.sh/helm/v3/pkg/pusher"
	"helm.sh/helm/v3/pkg/registry"
)

var host = "myregistry.azurecr.io"
var username = "username"
var password = "password"
var ref = "helmtest/hello-world/hello-world-0.1.0.tgz"
var repo = "helmdemo"

func main() {
	client, err := registry.NewClient()
	checkError(err)

	err = client.Login(host, registry.LoginOptBasicAuth(username, password))
	checkError(err)

	p, err := pusher.NewOCIPusher(pusher.WithRegistryClient(client))
	checkError(err)

	url := fmt.Sprintf("%s/%s", host, repo)
	err = p.Push(ref, url)
	checkError(err)

	fmt.Println("Push succeeded")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
