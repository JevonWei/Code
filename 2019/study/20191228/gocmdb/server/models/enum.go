package models

const (
	StatusUnlock = 0
	StatusLock   = 1
)

const (
	AlarmTypeOffline = iota
	AlarmTypeCPU
	AlarmTypeRAM
)

const (
	AlarmStatusNew = iota
	AlarmStatusDoing
	AlarmStatusComplete
)
