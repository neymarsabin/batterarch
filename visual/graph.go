package visual

import (
	"fmt"
	"os"
	"strconv"

	"github.com/guptarohit/asciigraph"
	database "github.com/neymarsabin/batterarch/models"
)

func GenerateGraph(records []database.BatteryDetails) {
	data := []float64{}
	for _, record := range records {
		floatValue, err := strconv.ParseFloat(record.BatteryLevel, 64)
		if err != nil {
			fmt.Println("Error while parsing float: ", err)
			os.Exit(1)
		}
		data = append(data, floatValue)
	}
	graph := asciigraph.Plot(data, asciigraph.Offset(10), asciigraph.Height(10), asciigraph.Width(50))
	fmt.Println(graph)
}
