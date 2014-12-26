package topfee

import (
	"time"
)

type TopInfo struct {
	//TimeStamp time.Time
	//Uptime    time.Duration
	Time   string
	Uptime string

	NumUsers int

	Load1  float32
	Load5  float32
	Load15 float32

	Tasks         int
	TasksRunning  int
	TasksSleeping int
	TasksStopped  int
	TasksZombie   int

	CpuUser    float32
	CpuSys     float32
	CpuNice    float32
	CpuIdle    float32
	CpuWaiting float32
	CpuHi      float32
	CpuSi      float32
	CpuSt      float32

	MemTotal   uint64
	MemUsed    uint64
	MemFree    uint64
	MemBuffers uint64

	SwapTotal  uint64
	SwapUsed   uint64
	SwapFree   uint64
	SwapCached uint64

	Processes []*Process
}

type Process struct {
	Pid            int
	PidString      string
	Name           string
	User           string
	Priority       string
	Nice           int
	Virtual        uint64
	VirtualString  string
	Resident       uint64
	ResidentString string
	Shared         uint64
	SharedString   string
	State          string
	Cpu            float32
	Mem            float32
	Time           time.Duration
	TimeString     string

	extended *Extended
}
