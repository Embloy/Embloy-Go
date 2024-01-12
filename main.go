package main

import (
    "fmt"
    "github.com/embloy/embloy-go/embloy"
)

func main() {
    client := embloy.NewEmbloyClient("eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOjEsImV4cCI6MTcwNzY4NDMzMiwidHlwIjoicHJlbWl1bSIsImlhdCI6MTcwNTAyNDA5NywiaXNzIjoibWFua2RlIn0.Xjt10QO0aMQTk-9Ac55koute_Fhlj2ExAlk8CpMnVME", map[string]string{
        "mode":        "job",
        "success_url": "your-success-url",
        "cancel_url":  "your-cancel-url",
        "job_slug":    "your-job-slug",
    })

    result, err := client.MakeRequest()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Result:", result)
}
