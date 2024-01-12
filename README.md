# Embloy Go

Embloy's Go SDK for interacting with your Embloy integration.

## Usage

Install Embloy-Go SDK:

```go
import "github.com/embloy/embloy-go/embloy"
```

Then you can run the following command to retrieve the package:

```go
go get -u github.com/embloy/embloy-go/embloy
```

Integrate it in your service:

```go
import (
    "fmt"
    "github.com/embloy/embloy-go/embloy"
)

func your-service-endpoint() {
    client := embloy.NewEmbloyClient("your-client-token", map[string]string{
        "mode":        "job",
        "success_url": "/success",
        "cancel_url":  "/failure",
        "job_slug":    "job#1",
    })

    response, err := client.MakeRequest()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("redirect_url:", response)
}
```

---

© Carlo Bortolan, Jan Hummel

> Carlo Bortolan &nbsp;&middot;&nbsp;
> GitHub [@carlobortolan](https://github.com/carlobortolan) &nbsp;&middot;&nbsp;
> contact via [bortolanoffice@embloy.com](mailto:bortolanoffice@embloy.com)
>
> Jan Hummel &nbsp;&middot;&nbsp;
