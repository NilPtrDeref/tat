package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "tat",
	Short: "Image to ascii art converter.",
	Long:  `Tat, short for tet-a-tet is an image to ascii art converter.`,
}

func Execute() {
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}
