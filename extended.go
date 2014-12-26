package topfee

import ()

const PROC_DIR = "/proc/"
const TASK_DIR = "/task"
const FD_DIR = "/fd"
const STATUS = "/status"

type Extended struct {
	status          *StatusInfo
	numOpenFiles    int
	numSubProcesses int
}

func numOpenFiles(pid string) int {
	return numFilesInDir(PROC_DIR + pid + TASK_DIR)
}

func numSubProcesses(pid string) int {
	return numFilesInDir(PROC_DIR + pid + FD_DIR)
}

func extended(p *Process) error {
	p.extended = new(Extended)
	p.extended.numOpenFiles = numOpenFiles(p.PidString)
	p.extended.numSubProcesses = numSubProcesses(p.PidString)
	status, err := getStatus(p.PidString)
	if err != nil {
		return err
	}
	p.extended.status = status

	return nil

}
