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

package main

import (
  "os"
  //"fmt"

  "github.com/qnib/go-testhttpd/cmd"
  "github.com/urfave/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "go-testhttpd"
    app.Version = "0.0.1"
    app.Flags = []cli.Flag {
      cli.StringFlag{
        Name: "status-sequence",
        Value: "200,201,202",
        Usage: "Sequence of HTTP status codes to return",
      },
    }
    app.Action = func(c *cli.Context) error {
      return cmd.TestServer(c)
    }
    app.Run(os.Args)
}
