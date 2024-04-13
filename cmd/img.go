package cmd

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/spf13/cobra"
)

var img = &cobra.Command{
	Use:   "img",
	Short: "Convert a given input image to ascii art.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filepath := args[0]
		file, err := os.Open(filepath)
		if err != nil {
			return err
		}
		defer file.Close()

		graphic, format, err := image.Decode(file)
		if err != nil {
			return err
		}

		fmt.Println(format)
		fmt.Println(graphic.Bounds())
		return nil
	},
}

func init() {
	root.AddCommand(img)
}
