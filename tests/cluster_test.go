package tests

import (
	cluster "Soccer-Penalty-Kick-ML-Threading/clustering"
	"fmt"
	"testing"
)

/**
 * TestCluster()
 * Tests the various methods and algorithms in the cluster package
 * Uses scatterplot to view results
 */
func TestCluster(t *testing.T) {

	// Generate sample data
	data := [][]float64{{6, 8, 0, 2, 1}, {0, 2, 12, 8, 9}, {4, 25, 11, 6, 5}}
	fmt.Println(data)

	// Test PCA
	result, _ := cluster.PCA(data, 2)
	fmt.Println("===PCA RESULT===")
	fmt.Println(result)

	// Print results and compare

	// EXAMPLE OF SCATTERPLOT USE FROM DOCUMENTATION
	// rnd := rand.New(rand.NewSource(1))

	// // randomPoints returns some random x, y points
	// // with some interesting kind of trend.
	// randomPoints := func(n int) plotter.XYs {
	// 	pts := make(plotter.XYs, n)
	// 	for i := range pts {
	// 		if i == 0 {
	// 			pts[i].X = rnd.Float64()
	// 		} else {
	// 			pts[i].X = pts[i-1].X + rnd.Float64()
	// 		}
	// 		pts[i].Y = pts[i].X + 10*rnd.Float64()
	// 	}
	// 	return pts
	// }

	// n := 15
	// scatterData := randomPoints(n)
	// lineData := randomPoints(n)
	// linePointsData := randomPoints(n)

	// p := plot.New()
	// p.Title.Text = "Points Example"
	// p.X.Label.Text = "X"
	// p.Y.Label.Text = "Y"
	// p.Add(plotter.NewGrid())

	// s, err := plotter.NewScatter(scatterData)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	// s.GlyphStyle.Radius = vg.Points(3)

	// l, err := plotter.NewLine(lineData)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// l.LineStyle.Width = vg.Points(1)
	// l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	// l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// lpLine.Color = color.RGBA{G: 255, A: 255}
	// lpPoints.Shape = draw.CircleGlyph{}
	// lpPoints.Color = color.RGBA{R: 255, A: 255}

	// p.Add(s, l, lpLine, lpPoints)
	// p.Legend.Add("scatter", s)
	// p.Legend.Add("line", l)
	// p.Legend.Add("line points", lpLine, lpPoints)

	// err = p.Save(350, 350, "./img/scatter.png")
	// if err != nil {
	// 	log.Panic(err)
	// }
}
