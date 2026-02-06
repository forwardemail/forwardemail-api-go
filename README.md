# Go client for Forward Email

- [Forward Email API Documentation](https://forwardemail.net/en/email-api)
- [Forward Email Terraform Provider](https://github.com/forwardemail/terraform-provider-forwardemail)

### How to install

```shell
$ go get github.com/forwardemail/forwardemail-api-go
```

### Basic usage

```go
import "github.com/forwardemail/forwardemail-api-go/forwardemail"

client := forwardemail.NewClient(forwardemail.ClientOptions{
    ApiKey: key,
})

account, err := client.GetAccount()
```
