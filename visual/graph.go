package visual

import (
	"fmt"

	"github.com/guptarohit/asciigraph"
	database "github.com/neymarsabin/batterarch/models"
)

func GenerateGraph(records []database.BatteryDetails) {
	data := []float64{}
	for _, record := range records {
		floatValue := float64(record.BatteryLevel)
		data = append(data, floatValue)
	}
	graph := asciigraph.Plot(data, asciigraph.Offset(10), asciigraph.Height(10), asciigraph.Width(50))
	fmt.Println(graph)
}
