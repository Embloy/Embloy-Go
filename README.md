# Embloy Go

Embloy's Go SDK for interacting with your Embloy integration.

## Usage

Install Embloy-Go SDK:

```
go get embloy/embloy-go
```

Integrate it in your service:

```
package main

import (
    "fmt"
    "github.com/embloy/embloy-go/embloy"
)

func main() {
    client := embloy.NewEmbloyClient("your-client-token", map[string]string{
        "mode":        "job",
        "success_url": "your-success-url",
        "cancel_url":  "your-cancel-url",
        "job_slug":    "your-job-slug",
    })

    result, err := client.MakeRequest()
    if err != nil {
        fmt.Println("Error:", err) // Handle error
        return
    }

    fmt.Println("redirect_url:", result)
}
```


---

© Carlo Bortolan, Jan Hummel

> Carlo Bortolan &nbsp;&middot;&nbsp;
> GitHub [@carlobortolan](https://github.com/carlobortolan) &nbsp;&middot;&nbsp;
> contact via [bortolanoffice@embloy.com](mailto:bortolanoffice@embloy.com)
>
> Jan Hummel &nbsp;&middot;&nbsp;
