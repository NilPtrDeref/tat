package cmd

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"golang.org/x/image/draw"

	"github.com/charmbracelet/lipgloss"
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

		graphic, _, err := image.Decode(file)
		if err != nil {
			return err
		}

		width, err := cmd.Flags().GetInt("width")
		if err != nil {
			return err
		}
		height, err := cmd.Flags().GetInt("height")
		if err != nil {
			return err
		}

		resized := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
		draw.NearestNeighbor.Scale(resized, resized.Rect, graphic, graphic.Bounds(), draw.Over, nil)

		style := lipgloss.NewStyle()
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				r, g, b, a := resized.At(x, y).RGBA()
				color := fmt.Sprintf("#%x%x%x%x", uint8(r), uint8(g), uint8(b), uint16(a))
				space := style.Background(lipgloss.Color(color)).Render("  ")
				fmt.Print(space)
			}
			fmt.Print("\n")
		}

		return nil
	},
}

func init() {
	img.Flags().IntP("width", "x", 50, "Width of the resulting ascii art.")
	img.Flags().IntP("height", "y", 50, "Height of the resulting ascii art.")
	root.AddCommand(img)
}
