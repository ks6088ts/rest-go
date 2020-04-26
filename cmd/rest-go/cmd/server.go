/*
Copyright © 2020 ks6088ts <ks6088ts@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package cmd ...
package cmd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type serverOptions struct {
	port int
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test: curl localhost:8080/ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func newServerOptions() *serverOptions {
	return &serverOptions{}
}

func newCmdServer() *cobra.Command {
	o := newServerOptions()
	cmd := &cobra.Command{
		Use:   "server",
		Short: "run REST server",
		Long:  `run REST server`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("port =", o.port)
			r := setupRouter()

			// Listen and Server in 0.0.0.0:port
			r.Run(fmt.Sprintf(":%d", o.port))
		},
	}

	cmd.PersistentFlags().IntVarP(&o.port, "port", "p", 8080, "type port")

	return cmd
}

func init() {
	cmd := newCmdServer()
	rootCmd.AddCommand(cmd)
}
