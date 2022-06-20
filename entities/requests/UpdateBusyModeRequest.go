package request

import (
	"errors"
	"wafflehacks/entities/busymode"
)

type BusyMode string

func (bm BusyMode) CheckForExists() error {
	switch bm {
	case busymode.ActiveMode:
		return nil
	case busymode.BusyMode:
		return nil
	case busymode.InvisibleMode:
		return nil
	case busymode.OfflineMode:
		return nil
	default:
		return errors.New("не известный тип занятости")
	}
}
