
package main

import (
       "encoding/json"
       "io/ioutil"
       "log"
       "net/http"
)

const (
      apiURL = "https://api.github.com"
      userEndpoint = "/users/"
)

type User struct {
     Login string `json:"login"`
     Name string `json:"name"`
     Email string `json:"email"`
     Bio string `json:"bio"`
}

func getUsers(name string) User {
     resp, err := http.Get(apiURL + userEndpoint + name)
     if err != nil {
     	log.Fatalf("Error retrieving data: %s\n", err)
     }
     defer resp.Body.Close()

     body, err := ioutil.ReadAll(resp.Body)
     if err != nil {
     	log.Fatalf("Error reading data: %s\n", err)
     }

     var user User
     json.Unmarshal(body, &user)
     return user
}

