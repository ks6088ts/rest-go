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

	"github.com/ks6088ts/rest-go/pkg/controller"
	"github.com/ks6088ts/rest-go/pkg/e"
	"github.com/ks6088ts/rest-go/pkg/repository"
	"github.com/ks6088ts/rest-go/pkg/router"
	"github.com/ks6088ts/rest-go/pkg/service"
	"github.com/ks6088ts/rest-go/pkg/setting"
	"github.com/spf13/cobra"
)

type serverOptions struct {
	port     int
	database string
}

func newServerOptions() *serverOptions {
	return &serverOptions{}
}

func newCmdServer() *cobra.Command {
	o := newServerOptions()
	cmd := &cobra.Command{
		Use:   "server",
		Short: "run REST server",
		Long: `run REST server:
To run mysql server locally:
  $ docker-compose up -d mysql
To post a data to a server:
$ curl http://localhost:8080/products -X POST -H "Content-Type: application/json" -d  '{"Name": "sample", "Description": "this is it"}'
		`,
		Run: func(cmd *cobra.Command, args []string) {

			// Setup setting
			dbsettings, err := setting.NewDatabase(o.database)
			if err != nil {
				fmt.Println(err)
				e.PrintError(e.ErrorLoadSettings)
				os.Exit(e.ErrorLoadSettings)
			}
			fmt.Printf("[Setting loaded] file: %s, dbms: %s, connect: %s\n", o.database, dbsettings.DbmsString(), dbsettings.ConnectString())

			// Setup repository
			session, err := repository.NewSession(dbsettings.DbmsString(), dbsettings.ConnectString())
			if err != nil {
				fmt.Println(err)
				e.PrintError(e.ErrorDatabaseConnection)
				os.Exit(e.ErrorDatabaseConnection)
			}
			defer func() {
				if err := session.Close(); err != nil {
					fmt.Println(err)
					e.PrintError(e.ErrorDatabaseConnection)
					os.Exit(e.ErrorDatabaseConnection)
				}
			}()
			fmt.Println("[Database connected]")

			// Setup service
			service, err := service.NewService(session)
			if err != nil {
				fmt.Println(err)
				e.PrintError(e.ErrorCreateService)
				os.Exit(e.ErrorCreateService)
			}

			// Setup controller
			controller, err := controller.NewController(service)
			if err != nil {
				fmt.Println(err)
				e.PrintError(e.ErrorCreateController)
				os.Exit(e.ErrorCreateController)
			}

			// Setup router
			r, err := router.NewRouter(o.port, controller)
			if err != nil {
				fmt.Println(err)
				e.PrintError(e.ErrorCreateController)
				os.Exit(e.ErrorCreateController)
			}

			r.Run()
		},
	}

	cmd.PersistentFlags().IntVarP(&o.port, "port", "p", 8080, "type port")
	cmd.PersistentFlags().StringVarP(&o.database, "database", "d", "./settings.sample.yml", "type settings file path")

	return cmd
}

func init() {
	cmd := newCmdServer()
	rootCmd.AddCommand(cmd)
}
