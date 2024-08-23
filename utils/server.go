package utils

import (
	"gin-vue-admin/global"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"time"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

type Server struct {
	Os   Os     `json:"os"`
	Cpu  Cpu    `json:"cpu"`
	Ram  Ram    `json:"ram"`
	Disk []Disk `json:"disk"`
}

type Os struct {
	GOOS         string `json:"goos"`
	NumCPU       int    `json:"num_cpu"`
	Compiler     string `json:"compiler"`
	GoVersion    string `json:"go_version"`
	NumGoroutine int    `json:"num_goroutine"`
}

type Cpu struct {
	Cpus  []float64 `json:"cpus"`
	Cores int       `json:"cores"`
}

type Ram struct {
	UsedMB      int `json:"used_mb"`
	TotalMB     int `json:"total_mb"`
	UsedPercent int `json:"used_percent"`
}

type Disk struct {
	MountPoint  string `json:"mount_point"`
	UsedMB      int    `json:"used_mb"`
	UsedGB      int    `json:"used_gb"`
	TotalMB     int    `json:"total_mb"`
	TotalGB     int    `json:"total_gb"`
	UsedPercent int    `json:"used_percent"`
}

// InitOS
// @function: InitCPU
// @description: OS信息
// @return: o Os, err error
func InitOS() (o Os) {
	o.GOOS = runtime.GOOS
	o.NumCPU = runtime.NumCPU()
	o.Compiler = runtime.Compiler
	o.GoVersion = runtime.Version()
	o.NumGoroutine = runtime.NumGoroutine()
	return o
}

// InitCPU
// @function: InitCPU
// @description: CPU信息
// @return: c Cpu, err error
func InitCPU() (c Cpu, err error) {
	if cores, err := cpu.Counts(false); err != nil {
		return c, err
	} else {
		c.Cores = cores
	}
	if cpus, err := cpu.Percent(time.Duration(200)*time.Millisecond, true); err != nil {
		return c, err
	} else {
		c.Cpus = cpus
	}
	return c, nil
}

// InitRAM
// @function: InitRAM
// @description: RAM信息
// @return: r Ram, err error
func InitRAM() (r Ram, err error) {
	if u, err := mem.VirtualMemory(); err != nil {
		return r, err
	} else {
		r.UsedMB = int(u.Used) / MB
		r.TotalMB = int(u.Total) / MB
		r.UsedPercent = int(u.UsedPercent)
	}
	return r, nil
}

// InitDisk
// @function: InitDisk
// @description: 硬盘信息
// @return: d Disk, err error
func InitDisk() (d []Disk, err error) {
	for i := range global.GVA_CONFIG.DiskList {
		mp := global.GVA_CONFIG.DiskList[i].MountPoint
		if u, err := disk.Usage(mp); err != nil {
			return d, err
		} else {
			d = append(d, Disk{
				MountPoint:  mp,
				UsedMB:      int(u.Used) / MB,
				UsedGB:      int(u.Used) / GB,
				TotalMB:     int(u.Total) / MB,
				TotalGB:     int(u.Total) / GB,
				UsedPercent: int(u.UsedPercent),
			})
		}
	}
	return d, nil
}
