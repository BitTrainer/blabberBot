package blabberBot

import "errors"

//BroadcastManager ...
type BroadcastManager struct {
	DataSource BroadcastDataSource
}

//Save ...
func (manager BroadcastManager) Save(message string, minutesUntilExpiry int) (id int, err error) {
	if manager.DataSource == nil {
		err = errors.New("Data Source not set before calling save. Please use the DataSource property to set this.")
		return id, err
	}
	id, err = manager.DataSource.Insert(message, minutesUntilExpiry)
	return id, err
}

//GetCurrentBroadcasts ...
func (manager BroadcastManager) GetCurrentBroadcasts() (broadcasts []Broadcast, err error) {
	if manager.DataSource == nil {
		err = errors.New("Data Source not set before calling save. Please use the DataSource property to set this.")
		return broadcasts, err
	}
	return manager.DataSource.SelectAllActive()
}
