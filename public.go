package topfee

import (
	"sort"
)

type SortType int

const (
	NoSort       SortType = iota
	SortCpu               = iota
	SortMem               = iota
	SortName              = iota
	SortPid               = iota
	SortResident          = iota
	SortShared            = iota
	SortTime              = iota
	SortVirtual           = iota
	// extended
	SortNumOpenFiles    = iota
	SortNumSubProcesses = iota
	SortNumThreads      = iota
)

func Top() (*TopInfo, error) {
	out, err := topAsString()
	if err != nil {
		return nil, err
	}
	topString := string(out)
	return parseTopString(topString)
}

func TopSorted(sortBy SortType) (*TopInfo, error) {
	if sortBy >= SortNumOpenFiles {

	}
	topInfo, err := Top()
	if err != nil {
		return nil, err
	}
	var thisSort sort.Interface
	switch sortBy {
	case SortCpu:
		thisSort = ByCpu(topInfo.Processes)
	case SortMem:
		thisSort = ByMem(topInfo.Processes)
	case SortName:
		thisSort = ByName(topInfo.Processes)
	case SortPid:
		thisSort = ByPid(topInfo.Processes)
	case SortResident:
		thisSort = ByResident(topInfo.Processes)
	case SortShared:
		thisSort = ByShared(topInfo.Processes)
	case SortTime:
		thisSort = ByTime(topInfo.Processes)
	case SortVirtual:
		thisSort = ByVirtual(topInfo.Processes)
	}
	sort.Sort(thisSort)

	return topInfo, nil

}

func TopExtended() (*TopInfo, error) {
	topInfo, err := Top()
	if err != nil {
		return nil, err
	}
	for _, p := range topInfo.Processes {
		extended(p)
	}

	return topInfo, nil
}
