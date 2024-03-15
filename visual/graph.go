package visual

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	database "github.com/neymarsabin/batterarch/models"
)

type GraphXY struct {
	XData []int64
	YData []opts.LineData
}

func HtmlGraph(GraphXY *GraphXY) {
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Battery Vs Timestamp",
			Subtitle: "Dumb",
		}),
	)

	line.SetXAxis(GraphXY.XData).
		AddSeries("Battery Level Line Graph", GraphXY.YData).
		SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))

	f, err := os.Create("graph.html")
	if err != nil {
		fmt.Println("Error while creating the graph : ", err)
	}

	line.Render(f)
	exec.Command("xdg-open", "graph.html").Start()
}

func BatteryLevelGraphXY(records []database.BatteryDetails) *GraphXY {
	xData := make([]int64, 0)
	yData := make([]opts.LineData, 0)
	for _, record := range records {
		xData = append(xData, record.RecordedAtUnix)
		yData = append(yData, opts.LineData{Value: record.BatteryLevel})
	}

	return &GraphXY{
		XData: xData,
		YData: yData,
	}
}

func BatteryCycleGraphXY(records []database.BatteryDetails) *GraphXY {
	xData := make([]int64, 0)
	yData := make([]opts.LineData, 0)
	for _, record := range records {
		xData = append(xData, record.RecordedAtUnix)
		yData = append(yData, opts.LineData{Value: record.CycleCount})
	}

	return &GraphXY{
		XData: xData,
		YData: yData,
	}
}
