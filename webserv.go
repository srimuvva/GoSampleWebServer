package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "flag"
    "strconv"
)

// Struct to de-serialize json ouput of http://uinames.com/api/
type Person struct {
    FirstName string `json:"name"`
    SurName  string `json:"surname"`
}

// Struct to de-serialize json ouput of
// http://api.icndb.com/jokes/random?firstname="+person.FirstName+"&surname="+person.SurName
type Resp struct {
    Type string `json:"type"`
    Value struct {
      Id int `json:"id"`
      Joke string `json:"joke"`
    } `json:"value"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://uinames.com/api/")
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    var person Person
    // Decode json to Struct
    if err = json.Unmarshal(body, &person); err != nil {
        panic(err)
    }
    joke_url := fmt.Sprintf("http://api.icndb.com/jokes/random?firstName=%s&lastName=%s",person.FirstName,person.SurName)
    resp, err = http.Get(joke_url)
    body, err = ioutil.ReadAll(resp.Body)
    var resp2 Resp
    if err = json.Unmarshal(body, &resp2); err != nil {
        panic(err)
    }
    // Writeback response string.
    fmt.Fprintf(w, resp2.Value.Joke)
}

func main() {
    // Get the port number from cmd line, default is 8080 if not given from cmd line.
    var port = flag.Int("port", 8080, "Port number to listen-on")
    flag.Parse()

    // Register callback with http library.
    http.HandleFunc("/", handler)
    fmt.Printf("Starting webserver on port %v",*port)

    // Start the server on IP ANY:<Port>
    http.ListenAndServe(":"+strconv.Itoa(*port), nil)
}
