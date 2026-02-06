# Go client for Forward Email

<div align="center">

[![Forward Email SDK Tests](https://github.com/forwardemail/forwardemail-api-go/actions/workflows/tests.yml/badge.svg)](https://github.com/forwardemail/forwardemail-api-go/actions/workflows/tests.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/forwardemail/forwardemail-api-go)](https://goreportcard.com/report/github.com/forwardemail/forwardemail-api-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/forwardemail/forwardemail-api-go.svg)](https://pkg.go.dev/github.com/forwardemail/forwardemail-api-go)
[![Go Version](https://img.shields.io/github/go-mod/go-version/forwardemail/forwardemail-api-go)](go.mod)
[![License](https://img.shields.io/github/license/forwardemail/forwardemail-api-go)](LICENSE)

</div>

- [Forward Email API Documentation](https://forwardemail.net/en/email-api)
- [Forward Email Terraform Provider](https://github.com/forwardemail/terraform-provider-forwardemail)

### How to install

```shell
$ go get github.com/forwardemail/forwardemail-api-go
```

### Basic usage

```go
import "github.com/forwardemail/forwardemail-api-go/forwardemail"

client, err := forwardemail.NewClient(forwardemail.ClientOptions{
    APIKey: key,
})
if err != nil {
    log.Fatal(err)
}

account, err := client.GetAccount()
```
