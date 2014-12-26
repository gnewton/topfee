package topfee

import (
	"fmt"
)

func getStatus(pid string) (*StatusInfo, error) {
	values, err := readColonFile(PROC_DIR + pid + STATUS)
	if err != nil {
		return nil, err
	}
	//fmt.Println(values)

	a := values["Name"]
	a = values["State"]
	a = values["Tgid"]
	a = values["Pid"]
	a = values["PPid"]
	a = values["TracerPid"]
	a = values["Uid"]
	a = values["Gid"]
	a = values["FDSize"]
	a = values["Groups"]
	a = values["VmPeak"]
	a = values["VmSize"]
	a = values["VmLck"]
	a = values["VmPin"]
	a = values["VmHWM"]
	a = values["VmRSS"]
	a = values["VmData"]
	a = values["VmStk"]
	a = values["VmExe"]
	a = values["VmLib"]
	a = values["VmPTE"]
	a = values["VmSwap"]
	a = values["Threads"]
	a = values["SigQ"]
	a = values["SigPnd"]
	a = values["ShdPnd"]
	a = values["SigBlk"]
	a = values["SigIgn"]
	a = values["SigCgt"]
	a = values["CapInh"]
	a = values["CapPrm"]
	a = values["CapEff"]
	a = values["CapBnd"]
	a = values["Seccomp"]
	a = values["Cpus_allowed"]
	a = values["Cpus_allowed_list"]
	a = values["Mems_allowed"]
	a = values["Mems_allowed_list"]
	a = values["voluntary_ctxt_switches"]
	a = values["nonvoluntary_ctxt_switches"]
	fmt.Println("mm " + a)
	return nil, nil
}
