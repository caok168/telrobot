package util

import "github.com/jinzhu/gorm"

func CheckNotFoundError(err error) error {
	if err == gorm.ErrRecordNotFound {
		return nil
	}

	return err
}
