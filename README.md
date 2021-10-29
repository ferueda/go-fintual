go-fintual
=======

[![GoDoc](https://godoc.org/github.com/ferueda/go-fintual?status.svg)](http://godoc.org/github.com/ferueda/go-fintual)

go-fintual is a Go client library for accessing the [Fintual API](https://fintual.cl/api-docs).

## Installation

To install the library, simply

`go get github.com/ferueda/go-fintual`

## Usage
```go
import "github.com/ferueda/go-fintual/fintual"
```

Construct a new Fintual client, then use the various services on the client to access different parts of the Fintual API. For example:

```go
client := fintual.NewClient(nil)
ctx := context.Background()

// list all banks
banks, err := client.Banks.ListAll(ctx, nil)
```

Some API methods have optional parameters that can be passed in order to filter results. For example:

```go
client := fintual.NewClient(nil)
ctx := context.Background()

// list all banks with the word "nova" in their name
params := &fintual.BankListParams{Query: "nova"}
banks, err := client.Banks.ListAll(ctx, params)
```

### Authentication
For authenticating the client, just call the provided Client.Authenticate method with valid credentials:

```go
client := fintual.NewClient(nil)
ctx := context.Background()

err := client.Authenticate(ctx, "email@email.com", "validPassword")
```
## Coverage

### Auth
* POST /access_token

### Asset Providers
* GET /asset_providers
* GET /asset_providers/:id
* GET /asset_providers/:id/conceptual_assets

### Banks
* GET /banks

### Conceptual Assets
* GET /conceptual_assets
* GET /conceptual_assets/:id
* GET /conceptual_assets/:id/real_assets

### Goals
* GET /goals
* GET /goals/:id

### Real Assets
* GET /real_assets/:id
* GET /real_assets/:id/days
* GET /real_assets/:id/expense_ratio

## How to Contribute

* Fork a repository
* Add/Fix something
* Check that tests are passing
* Create PR

Current contributors:

- [Felipe Rueda](https://github.com/ferueda)

## License ##

This library is distributed under the MIT License found in the [LICENSE](./LICENSE)
file.