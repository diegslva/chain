/*
Package chain implements the web API for https://chain.com/. There are methods
to connect to both the blockchain query endpoints and webhook endpoints.

A full overview of the Chain API can be found here https://chain.com/docs.

When using webhooks, make sure you read https://chain.com/docs#webhooks-setup
on how to setup your system.

Example

Executing an endpoint call to get the latest block:

    package main

    import (
        "fmt"
        "net/http"

        "github.com/qedus/chain"
    )

    func main() {
        c := chain.New(http.DefaultClient, chain.MainNet, [apiKeyID], [apiKeySecret])
        block, err := c.GetLatestBlock()
        if err != nil {
            panic(err)
        }

        fmt.Println("Block hash", block.Hash)
    }
*/
package chain
