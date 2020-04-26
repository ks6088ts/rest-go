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
	"os"

	"github.com/ks6088ts/rest-go/pkg/repository"
	"github.com/ks6088ts/rest-go/pkg/router"
	"github.com/spf13/cobra"
)

type serverOptions struct {
	port    int
	dbms    string
	connect string
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

			// Setup connection
			session, err := repository.NewSession(o.dbms, o.connect)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer func() {
				if err := session.Close(); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}()
			fmt.Printf("[Database connected] dbms: %s, connect: %s\n", o.dbms, o.connect)

			// Setup router
			r := router.NewRouter(o.port)
			r.Run()
		},
	}

	cmd.PersistentFlags().IntVarP(&o.port, "port", "p", 8080, "type port")
	cmd.PersistentFlags().StringVarP(&o.dbms, "dbms", "d", "mysql", "type dbms")
	cmd.PersistentFlags().StringVarP(&o.connect, "connect", "c", "root:password@tcp(localhost:3306)/db?parseTime=true", "type connect")

	return cmd
}

func init() {
	cmd := newCmdServer()
	rootCmd.AddCommand(cmd)
}
