package load

import (
	"github.com/shirou/gopsutil/load"

	"iSystem/formats"
	"iSystem/providers/cpu"
)

var (
	Load string
)

func init() {
	loadInfo, err := load.LoadAvg()
	if err != nil {
		panic(err)
	}

	Load = loadFormat(loadInfo)
}

func loadFormat(loadInfo *load.LoadAvgStat) string {
	loadAvg := loadInfo.Load1 / float64(cpu.CPUs*int(cpu.Cores))
	format := "%.2f, %.2f, %.2f"
	if loadAvg > 1.0 {
		return formats.Dangerf(format, loadInfo.Load1, loadInfo.Load5, loadInfo.Load15)
	} else if loadAvg > 0.7 {
		return formats.Warningf(format, loadInfo.Load1, loadInfo.Load5, loadInfo.Load15)
	} else {
		return formats.Successf(format, loadInfo.Load1, loadInfo.Load5, loadInfo.Load15)
	}
}