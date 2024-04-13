package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var img = &cobra.Command{
	Use:   "img",
	Short: "Convert a given input image to ascii art.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("img called")
	},
}

func init() {
	root.AddCommand(img)
}
