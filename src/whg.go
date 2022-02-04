package main

import (
    "context"
    "net/http"
    "time"
	"fmt"
    "io"
    "crypto/sha256"
    "encoding/hex"
    "bytes"

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
    var s = []byte(`{
        "ClaimStatus": [
          "Opened"
        ],
        "WHGClaimStatus": {
          "Restoration": [
            {
              "name": "firstContact",
              "completed": true
            },
            {
              "name": "jobCompletion",
              "completed": false
            }
          ],
          "Mitigation": [
            {
              "name": "firstContact",
              "completed": true
            },
            {
              "name": "mitigationWorkComplete",
              "completed": false
            }
          ]
        }
      }`)
    h := sha256.New()
    h.Write([]byte(s))
    hash := hex.EncodeToString(h.Sum(nil))

    fmt.Println(s, hash)
    //req, _ := http.NewRequest(http.MethodGet, "https://dev-api.westhillglobal.com/project/0086FBJ/contractors-matched", nil)
    req, _ := http.NewRequest(http.MethodPost, "https://apidev.westhillglobal.com/admin/allclaims", bytes.NewBuffer(s))
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