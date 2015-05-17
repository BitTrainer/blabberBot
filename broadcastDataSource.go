package service

//BroadcastDataSource manages broadcasts in permanent storage.
type BroadcastDataSource interface {
	Insert(message string, minutesUntilExpiry int) (id int, err error)
	SelectAllActive() (broadcasts []Broadcast, err error)
}
