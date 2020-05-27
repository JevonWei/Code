package main

// 或者host信息
import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

func main() {
	hostInfo, err := host.Info()
	fmt.Printf("%#v\n", hostInfo)
	fmt.Println(err)

	// 采集CPU信息
	cpuInfos, _ := cpu.Info()
	for _, cpuInfo := range cpuInfos {
		fmt.Printf("%#v\n", cpuInfo)
	}

	// 采集CPU的使用率
	// 显示CPU总的使用率
	cpuPercents, _ := cpu.Percent(time.Second, false)
	fmt.Println(cpuPercents)
	// 显示CPU每个core的使用率
	cpuPercents, _ = cpu.Percent(time.Second, true)
	fmt.Println(cpuPercents)

	// 采集内存信息
	memInfo, _ := mem.VirtualMemory()
	fmt.Println(memInfo)

	// 采集网络信息
	interInfos, _ := net.Interfaces()
	for _, interinfo := range interInfos {
		fmt.Println(interinfo)
	}

	fmt.Println("net IO")
	fmt.Println(net.IOCounters(false))
	fmt.Println(net.IOCounters(true))

	// 负载
	loadInfo, _ := load.Avg()
	fmt.Println(loadInfo)

	// disk信息
	// true仅显示硬件磁盘信息
	diskInfos, _ := disk.Partitions(true)
	for _, diskInfo := range diskInfos {
		fmt.Println(diskInfos)
		// 显示磁盘使用率
		diskUsage, _ := disk.Usage(diskInfo.Device)
		fmt.Println(diskUsage)
	}

	diskCounters, _ := disk.IOCounters()
	fmt.Println(diskCounters)
}
