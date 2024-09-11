// cmd/serve.go
package cmd

import (
	"fmt"
	"net/http"

	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!")
		})

		handler := cors.Default().Handler(mux)
		fmt.Println("Starting server on :8080")
		http.ListenAndServe(":8080", handler)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
