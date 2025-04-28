# goeduidiam

[![Go Reference](https://pkg.go.dev/badge/github.com/masv3971/goeduidiam.svg)](https://pkg.go.dev/github.com/masv3971/goeduidiam)

## Installation 
```
go get github.com/masv3971/goeduidiam
 ```

 ## Example
 ```go
 package main

import (
    "github.com/masv3971/goeduidiam"
)

func main() {
    ctx := context.TODO() // Make a useful context if you like
    
    iam := goeduidiam.New(goeduidiam.Config{
        URL: "example-eduid.sunet.se",
    })

    // resp is usual blank
    user, resp, err := iam.Users.Get(ctx, &goeduidiam.RequestCFG{
        ScimID: "testID",
    })
    if err != nil {
        // handle error
    }

    searchGroup, _, err := iam.Groups.Search(ctx, &goeduidiam.RequestCFG{
        Search: goeduidiam.SearchRequest{
            Schemas:    []string{},  
	        Filter:     "",     
	        StartIndex: 1, 
	        Count       1,   
	        Attributes  []string{},
    })
    if err != nil {
        // handle error
    }
}  
```

## Make release
1. Update RELEASE file with new tag and commit comment. 
2. ```$ make ```