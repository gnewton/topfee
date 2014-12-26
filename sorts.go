package topfee

import (
//"sort"
)

type ByPid []*Process

func (a ByPid) Len() int           { return len(a) }
func (a ByPid) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPid) Less(i, j int) bool { return a[i].Pid < a[j].Pid }

type ByName []*Process

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

type ByVirtual []*Process

func (a ByVirtual) Len() int           { return len(a) }
func (a ByVirtual) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVirtual) Less(i, j int) bool { return a[i].Virtual > a[j].Virtual }

type ByResident []*Process

func (a ByResident) Len() int           { return len(a) }
func (a ByResident) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByResident) Less(i, j int) bool { return a[i].Resident > a[j].Resident }

type ByShared []*Process

func (a ByShared) Len() int           { return len(a) }
func (a ByShared) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByShared) Less(i, j int) bool { return a[i].Shared > a[j].Shared }

type ByCpu []*Process

func (a ByCpu) Len() int           { return len(a) }
func (a ByCpu) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCpu) Less(i, j int) bool { return a[i].Cpu > a[j].Cpu }

type ByMem []*Process

func (a ByMem) Len() int           { return len(a) }
func (a ByMem) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByMem) Less(i, j int) bool { return a[i].Mem > a[j].Mem }

var byMem *ByMem

type ByTime []*Process

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].Time > a[j].Time }

var sortMap = map[SortType][]*Process{
//SortMem: ByMem,
}

func init() {

}
