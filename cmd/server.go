package cmd

import (
	"fmt"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/jeanmarcboite/librarytruc/internal/controllers"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.New()
		logger, _ := zap.NewProduction()
		// Add a ginzap middleware, which:
		//   - Logs all requests, like a combined access and error log.
		//   - Logs to stdout.
		//   - RFC3339 with UTC time format.
		r.Use(ginzap.Ginzap(logger, time.RFC3339, true))

		// Logs all panic to error log
		//   - stack means whether output the stack info.
		r.Use(ginzap.RecoveryWithZap(logger, true))

		// Example ping request.
		r.GET("/ping", func(c *gin.Context) {
			c.String(200, "pong "+fmt.Sprint(time.Now().Unix()))
		})

		// Example when panic happen.
		r.GET("/panic", func(c *gin.Context) {
			panic("An unexpected error happen!")
		})

		r.LoadHTMLGlob("internal/templates/*.html")
		r.GET("/", controllers.Home)
		r.GET("/book/:id", controllers.LookupID)
		r.Run() // listen and serve on 0.0.0.0:8080
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
