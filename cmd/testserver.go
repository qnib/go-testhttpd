// Copyright © 2017 Christian Kniep <christian@qnib.org>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
    "log"
    "fmt"
    "net/http"
    "strings"

    "github.com/urfave/cli"
)

// HTTPServer holds the codes and the channel
type HTTPServer struct {
    StatusSequence []string
    StatusChan chan string
}

// NewHTTPServer returns HTTPServer
func NewHTTPServer(sc []string) HTTPServer {
    return HTTPServer{
        StatusSequence: sc,
        StatusChan: make(chan string),
    }
}

func (hs *HTTPServer) fillStatusCodeChan() string {
    for {
        for _, c := range hs.StatusSequence {
            hs.StatusChan <- c
        }
    }
}

func (hs *HTTPServer) httpHandler(w http.ResponseWriter, r *http.Request) {
    sc := <-hs.StatusChan
    log.Println(sc)
    fmt.Fprintf(w, "Hi there, I love %s!", sc)

}

// Run starts the webserver
func (hs *HTTPServer) Run() {
    go hs.fillStatusCodeChan()
    http.HandleFunc("/", hs.httpHandler)
    http.ListenAndServe(":8080", nil)

}


// TestServer serves the status-sequence of HTTP codes
func TestServer(c *cli.Context) error {
    hserver := NewHTTPServer(strings.Split(c.String("status-sequence"), ","))
    hserver.Run()
    return nil
}
