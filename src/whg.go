package main

import (
    "context"
    "net/http"
    "time"
	"fmt"
    "io"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/aws/signer/v4"
)

func main() {
 
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))

    if err != nil {
        fmt.Printf("failed to load configuration, %v", err)
    }

    credentials, err := cfg.Credentials.Retrieve(context.TODO())
	
    if err != nil {
        fmt.Printf("failed to load credentials, %v", err)
    }

    // The signer requires a payload hash. This hash is for an empty payload.
    hash := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
    req, _ := http.NewRequest(http.MethodGet, "https://dev-api.westhillglobal.com/project/0086FBJ/contractors-matched", nil)
    signer := v4.NewSigner()
    err = signer.SignHTTP(context.TODO(), credentials, req, hash, "execute-api", cfg.Region, time.Now())

    if err != nil {
        // Handle error.
    }

    // Use `req`
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("failed to call remote service: (%v)\n", err)
		return
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("service returned a status not 200: (%d)\n", res)
		return
	} else {
        fmt.Printf("service returned a status 200: (%d)\n", res.Body)
        bodyBytes, err := io.ReadAll(res.Body)
        if err != nil {
            fmt.Printf("failed to call remote service: (%v)\n", err)
        }
        bodyString := string(bodyBytes)
        fmt.Printf(bodyString)
		return
    }
}