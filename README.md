# [Embloy Go](https://pkg.go.dev/github.com/embloy/embloy-go) &middot; [![GitHub license](https://img.shields.io/badge/license-AGPL3.0-blue.svg)](https://github.com/Embloy/Embloy-Go/blob/main/LICENSE) [![Go Reference](https://pkg.go.dev/badge/github.com/embloy/embloy-go.svg)](https://pkg.go.dev/github.com/embloy/embloy-go) [![Issues](https://img.shields.io/github/issues/Embloy/Embloy-Go)](https://github.com/Embloy/Embloy-Go/issues) [![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/Embloy/Embloy-Go/pulls)

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
	sessionData := embloy.SessionData{
		Mode:       "job",
		SuccessURL: "your-success-url",
		CancelURL:  "your-cancel-url",
		JobSlug:    "your-job-slug",
	}

	client := embloy.NewEmbloyClient("your-client-token", sessionData)

    response, err := client.MakeRequest()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("redirect_url:", response)
}
```

## Publish a new release

```bash
git tag v0.1.x
git push origin v0.1.x
```

---

© Carlo Bortolan, Jan Hummel

> Carlo Bortolan &nbsp;&middot;&nbsp;
> GitHub [@carlobortolan](https://github.com/carlobortolan) &nbsp;&middot;&nbsp;
> contact via [bortolanoffice@embloy.com](mailto:bortolanoffice@embloy.com)
>
> Jan Hummel &nbsp;&middot;&nbsp;
> GitHub [@github4touchdouble](https://github.com/github4touchdouble) &nbsp;&middot;&nbsp;
> contact via [hummeloffice@embloy.com](mailto:hummeloffice@embloy.com)
