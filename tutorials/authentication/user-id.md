# User ID Authentication

To create a new cloudns-go client with user ID authentication, you must
first collect the user ID and password from the ClouDNS control panel.

We will then pass the result of the `cloudns.AuthUserID()` function to the
`cloudns.New()` function.

```go
client, err := cloudns.New(
    cloudns.AuthUserID(12345, "JohnsPassword"),
)

if error != nil {
    fmt.Println("Error creating client:", err)
}
```

An if-statement similar to the one above can be used to check if a client was
successfully created.