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

		bounds := graphic.Bounds()
		resized := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{bounds.Dx() / 40, bounds.Dy() / 40}})
		draw.NearestNeighbor.Scale(resized, resized.Rect, graphic, graphic.Bounds(), draw.Over, nil)

		bounds = resized.Bounds()
		style := lipgloss.NewStyle()
		for y := 0; y < bounds.Dy(); y++ {
			for x := 0; x < bounds.Dx(); x++ {
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
	root.AddCommand(img)
}
