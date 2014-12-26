package topfee

import (
	"fmt"
	"log"
	"os/exec"
	. "strconv"
	. "strings"

	"time"
)

const TIME_INTERVAL = "15:04:05"

type LineParseError struct {
	lineName  string
	fieldName string
	line      string
	err       error
}

func (e LineParseError) Error() string {
	return fmt.Sprintf("LineParseError %v |  %v | %v | %v", e.lineName, e.fieldName, e.line, e.err)
}

func topAsString() (string, error) {
	out, err := exec.Command("top", "-b", "-n", "1").Output()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return string(out), nil
}

func parseTopString(t string) (*TopInfo, error) {
	ti := new(TopInfo)

	lines := Split(t, "\n")

	err := parseUptime(lines[0], ti)
	if err != nil {
		return nil, err
	}

	err = parseTasks(lines[1], ti)
	if err != nil {
		return nil, err
	}

	err = parseCpu(lines[2], ti)
	if err != nil {
		return nil, err
	}

	err = parseMem(lines[3], ti)
	if err != nil {
		return nil, err
	}

	err = parseSwap(lines[4], ti)
	if err != nil {
		return nil, err
	}

	err = parseProcesses(lines[5:], ti)
	if err != nil {
		return nil, err
	}

	return ti, nil

}

func parseUptime(l string, t *TopInfo) error {
	lineId := "uptime"
	parts := flatten(Split(l, " "))
	if len(parts) < 11 {
		return LineParseError{
			lineId,
			"all",
			l,
			nil,
		}
	}
	t.Time = parts[2]

	shift, err = parseUp(l, parts, t)
	if err != nil {
		return err
	}

	if len(parts) < 11+shift {
		return LineParseError{
			"Wrong length",
			"",
			l,
			nil,
		}
	}
	var err error
	t.NumUsers, err = Atoi(parts[5+shift])
	if err != nil {
		return LineParseError{
			lineId,
			"numUsers",
			l,
			err,
		}
	}

	t.Load1, err = fl(parts[9+shift][0 : len(parts[9+shift])-1])
	if err != nil {
		return LineParseError{
			lineId,
			"load1",
			l,
			err,
		}
	}

	//t.Load5, err = fl(parts[12][0 : len(parts[12])-1])
	t.Load5, err = fl(parts[10+shift][0 : len(parts[10+shift])-1])
	if err != nil {
		return LineParseError{
			lineId,
			"load5",
			l,
			err,
		}
	}

	t.Load15, err = fl(parts[11+shift])
	if err != nil {
		return LineParseError{
			lineId,
			"load15",
			l,
			err,
		}
	}
	return nil
}

const MINUTES = "min"
const HOUR = "min"
const DAY = "day"

func parseUp(line string, parts []string, t *TopInfo) (int, error) {
	if parts[5] == MINUTES {

	}
	if parts[5] == DAY {

	}
	return 1, nil
}

func parseTasks(l string, t *TopInfo) error {
	parts := flatten(Split(l, " "))
	//printParts(parts)
	var err error
	t.Tasks, err = Atoi(parts[1])
	ferr(err)
	t.TasksRunning, err = Atoi(parts[3])
	ferr(err)
	t.TasksSleeping, err = Atoi(parts[5])
	ferr(err)
	t.TasksStopped, err = Atoi(parts[7])
	ferr(err)
	t.TasksZombie, err = Atoi(parts[9])
	ferr(err)

	return nil
}

func parseCpu(l string, t *TopInfo) error {
	lineId := "cpu"
	parts := flatten(Split(l, " "))
	if len(parts) != 17 {
		return LineParseError{
			lineId,
			"Wrong # fields",
			Itoa(len(parts)) + ": " + l,
			nil,
		}
	}
	//printParts(parts)
	var err error
	t.CpuUser, err = fl(parts[1])
	if err != nil {
		return LineParseError{
			lineId,
			"cpuUser",
			"line",
			err,
		}
	}

	t.CpuSys, err = fl(parts[3])
	if err != nil {
		return LineParseError{
			lineId,
			"cpuSys",
			"line",
			err,
		}
	}

	t.CpuNice, err = fl(parts[5])
	if err != nil {
		return LineParseError{
			lineId,
			"cpuNice",
			"line",
			err,
		}
	}

	t.CpuIdle, err = fl(parts[7])
	if err != nil {
		return LineParseError{
			lineId,
			"cpuIdle",
			"line",
			err,
		}
	}

	t.CpuWaiting, err = fl(parts[9])
	if err != nil {
		return LineParseError{
			lineId,
			"cpuWaiting",
			"line",
			err,
		}
	}

	t.CpuHi, err = fl(parts[11])
	if err != nil {
		return LineParseError{
			lineId,
			"cpuHi",
			"line",
			err,
		}
	}

	t.CpuSi, err = fl(parts[13])
	if err != nil {
		return LineParseError{
			lineId,
			"cpuSi",
			"line",
			err,
		}
	}

	t.CpuSt, err = fl(parts[15])
	if err != nil {
		return LineParseError{
			lineId,
			"cpuSt",
			"line",
			err,
		}
	}
	return nil

}

func parseMem(l string, t *TopInfo) error {
	lineId := "memory"
	parts := flatten(Split(l, " "))
	printParts(parts)

	var err error
	t.MemTotal, err = i64(parts[2])
	if err != nil {
		return LineParseError{
			lineId,
			"MemTotal",
			l,
			err,
		}
	}

	t.MemUsed, err = i64(parts[4])
	if err != nil {
		return LineParseError{
			lineId,
			"MemUsed",
			l,
			err,
		}
	}

	t.MemFree, err = i64(parts[6])
	if err != nil {
		return LineParseError{
			lineId,
			"MemFree",
			l,
			err,
		}
	}

	t.MemBuffers, err = i64(parts[8])
	if err != nil {
		return LineParseError{
			lineId,
			"MemBuffers",
			l,
			err,
		}
	}
	return nil
}

func parseSwap(l string, t *TopInfo) error {
	lineId := "swap"
	parts := flatten(Split(l, " "))
	printParts(parts)
	var err error
	t.SwapTotal, err = i64(parts[2])
	if err != nil {
		return LineParseError{
			lineId,
			"SwapTotal",
			l,
			err,
		}
	}

	t.SwapUsed, err = i64(parts[4])
	if err != nil {
		return LineParseError{
			lineId,
			"SwapUsed",
			l,
			err,
		}
	}

	t.SwapFree, err = i64(parts[6])
	if err != nil {
		return LineParseError{
			lineId,
			"SwapFree",
			l,
			err,
		}
	}

	t.SwapCached, err = i64(parts[8])
	if err != nil {
		return LineParseError{
			lineId,
			"SwapCached",
			l,
			err,
		}
	}
	return nil
}

func parseProcesses(processes []string, t *TopInfo) error {
	t.Processes = make([]*Process, 0, 0)
	for i, processLine := range processes {
		if i < 2 {
			continue
		}
		fmt.Println(i)
		//parseProcess(processLine, t)
		if processLine == "" {
			continue
		}
		process, err := parseProcess(processLine)
		if err != nil {
			return err
		}
		//fmt.Printf("ZZZ:: %+v\n", process)
		t.Processes = append(t.Processes, process)
	}
	return nil
}

func parseProcess(l string) (*Process, error) {
	lineId := "process"
	process := new(Process)
	fmt.Println("[" + l + "]")
	fmt.Println("++")
	parts := Split(l, " ")
	//printParts(parts)
	//fmt.Println("XX")
	parts = flatten(Split(l, " "))
	printParts(parts)
	var err error
	process.PidString = parts[0]
	process.Pid, err = Atoi(process.PidString)
	if err != nil {
		return nil, LineParseError{
			lineId,
			"Pid",
			l,
			err,
		}
	}

	process.User = parts[1]
	process.Priority = parts[2]
	if err != nil {
		return nil, LineParseError{
			lineId,
			"Priority",
			l,
			err,
		}
	}

	process.Nice, err = Atoi(parts[3])
	if err != nil {
		return nil, LineParseError{
			lineId,
			"Nice",
			l,
			err,
		}
	}

	process.VirtualString = parts[4]
	process.Virtual, err = kilobytes(parts[4])
	if err != nil {
		return nil, LineParseError{
			lineId,
			"Virtual",
			l,
			err,
		}
	}

	process.ResidentString = parts[5]
	process.Resident, err = kilobytes(parts[5])
	if err != nil {
		return nil, LineParseError{
			lineId,
			"Resident",
			l,
			err,
		}
	}

	process.SharedString = parts[6]
	process.Shared, err = kilobytes(parts[6])
	if err != nil {
		return nil, LineParseError{
			lineId,
			"Shared",
			l,
			err,
		}
	}

	process.State = parts[7]
	if err != nil {
		return nil, LineParseError{
			lineId,
			"State",
			l,
			err,
		}
	}

	process.Cpu, err = fl(parts[8])
	if err != nil {
		return nil, LineParseError{
			lineId,
			"Cpu",
			l,
			err,
		}
	}

	process.Mem, err = fl(parts[9])
	if err != nil {
		return nil, LineParseError{
			lineId,
			"Mem",
			l,
			err,
		}
	}

	process.Name = parts[11]
	process.TimeString = parts[10]
	process.Time, err = makeTime(process.TimeString)
	if err != nil {
		return nil, err
	}
	return process, nil
}

func makeTime(t string) (time.Duration, error) {
	// assumes pattern TIME_INTERVAL = "15:04:05"

	parts := Split(t, ":")

	if len(parts) != 2 {
		return time.Minute * 0, LineParseError{
			"process",
			"Time",
			t,
			nil,
		}
	}

	tmp, err := i64(parts[0])
	if err != nil {
		return time.Minute * 0, err
	}
	hours := time.Duration(tmp)

	var minutes, seconds time.Duration
	if Contains(parts[1], ".") {
		parts = Split(parts[1], ".")

		tmp, err = i64(parts[0])
		if err != nil {
			return time.Minute * 0, err
		}
		minutes = time.Duration(tmp)

		tmp, err = i64(parts[1])
		if err != nil {
			return time.Minute * 0, err
		}
		seconds = time.Duration(tmp)
	} else {
		tmp, err := i64(parts[1])
		if err != nil {
			return time.Minute * 0, err
		}
		minutes = time.Duration(tmp)
	}

	return time.Hour*hours + time.Minute*minutes + time.Second*seconds, nil
}

func toBytes(s string) int {

	return 10
}

//func line0(s string,

func printParts(p []string) {
	fmt.Println("------------------------------------------------------------------")
	for i, part := range p {
		fmt.Println(Itoa(i) + " [" + part + "]")
	}
}

func fl(s string) (float32, error) {
	v, err := ParseFloat(s, 32)
	return float32(v), err
}

func fl64(s string) (float64, error) {
	v, err := ParseFloat(s, 64)
	return float64(v), err
}

func ferr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func perr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func flatten(s []string) []string {
	f := make([]string, 0, 0)
	for _, v := range s {
		if v == "" {
			continue
		} else {
			f = append(f, v)
		}
	}
	return f
}

const MEGA = 1024 * 1024
const GIGA = MEGA * 1024

func kilobytes(s string) (uint64, error) {
	lastChar := s[len(s)-1]

	if lastChar == 'm' || lastChar == 'g' {
		base, err := fl64(s[0 : len(s)-1])
		if err != nil {
			log.Println(err)
			return 0, err
		}
		switch {
		case lastChar == 'm':
			return uint64(base * MEGA), nil
		case lastChar == 'g':
			return uint64(base * GIGA), nil
		}

	}

	return i64(s)
}

func i64(s string) (uint64, error) {
	if s == "" {
		return 0, nil
	} else {
		val, err := ParseUint(s, 10, 64)
		if err != nil {
			fmt.Println("s=" + s)
			fmt.Println(err)
			return 0, err
		}
		return val, nil
	}

}
