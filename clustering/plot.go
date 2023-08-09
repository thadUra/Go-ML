package cluster

import (
	"errors"
	"image/color"
	"math/rand"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

/**
 * ScatterPlot2DimenData()
 * Creates a scatterplot with provided data, labels, and plot parameters
 * Params consist of output directory, title, and x/y labels
 */
func ScatterPlot2DimenData(data [][]float64, label []string, params []string) error {

	// Check plot params dimensions
	if len(params) != 4 {
		return errors.New("ScatterPlot2DimenData(): wrong plot parameter size")
	}

	// Get count of unique labels
	iterator := make(map[string]int)
	count := make(map[string]int)
	index := make(map[string]int)
	unique := 0
	for i := 0; i < len(label); i++ {
		if count[label[i]] == 0 {
			index[label[i]] = unique
			iterator[label[i]] = 0
			unique++
		}
		count[label[i]]++
	}

	// Build data points for plotter
	points := make([]plotter.XYs, unique)
	done := false
	iter := 0
	for !done {
		// Append data point
		var point plotter.XY
		point.X = data[iter][0]
		point.Y = data[iter][1]
		points[index[label[iter]]] = append(points[index[label[iter]]], point)
		// fmt.Printf("	Appended (%f, %f) to %s\n", point.X, point.Y, label[iter])
		iterator[label[iter]]++
		iter++
		// Check if done
		check := true
		for key := range iterator {
			if iterator[key] != count[key] {
				check = false
			}
		}
		done = check
	}

	// Initalize plot instance
	plot := plot.New()
	plot.Title.Text = params[1]
	plot.X.Label.Text = params[2]
	plot.Y.Label.Text = params[3]
	plot.Add(plotter.NewGrid())

	// Add data to plot instance
	for i := 0; i < len(points); i++ {
		s, err := plotter.NewScatter(points[i])
		if err != nil {
			return err
		}
		red := uint8(rand.Float64() * 255)
		green := uint8(rand.Float64() * 255)
		blue := uint8(rand.Float64() * 255)
		s.GlyphStyle.Color = color.RGBA{R: red, G: green, B: blue, A: 255}
		s.GlyphStyle.Radius = vg.Points(2)
		s.GlyphStyle.Shape = draw.CircleGlyph{}
		plot.Add(s)
		for key, val := range index {
			if val == i {
				plot.Legend.Add(key, s)
				break
			}
		}
	}

	// Save image of scatterplot
	err := plot.Save(300, 300, params[0])
	if err != nil {
		return err
	}
	return nil
}
