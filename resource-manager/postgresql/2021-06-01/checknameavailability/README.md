
## `github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/checknameavailability` Documentation

The `checknameavailability` SDK allows for interaction with the Azure Resource Manager Service `postgresql` (API Version `2021-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/postgresql/2021-06-01/checknameavailability"
```


### Client Initialization

```go
client := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CheckNameAvailabilityClient.Execute`

```go
ctx := context.TODO()
id := checknameavailability.NewSubscriptionID()

payload := checknameavailability.NameAvailabilityRequest{
	// ...
}


read, err := client.Execute(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```