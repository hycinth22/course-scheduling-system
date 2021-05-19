package models

import (
	"log"

	"github.com/beego/beego/v2/client/orm"
)

type Config struct {
	SelectedSemester string
	SelectedSchedule int
}

type ConfigInDb struct {
	Id int `orm:"pk;default(1)"`
	Config
}

func GetConfig() (Config, error) {
	c := ConfigInDb{
		Id: 1,
		Config: Config{
			SelectedSemester: "",
			SelectedSchedule: 0,
		},
	}

	err := o.Read(&c, "id")
	if err == orm.ErrNoRows {
		err = SaveConfig(Config{
			SelectedSemester: "",
			SelectedSchedule: 0,
		})
	}
	if err != nil {
		log.Printf("GetConfig Err: %d, %v\n", err)
		return Config{}, err
	}
	return c.Config, nil
}

func SaveConfig(c Config) error {
	wrap := ConfigInDb{
		Id:     1,
		Config: c,
	}

	_, err := o.InsertOrUpdate(&wrap)
	if err != nil {
		log.Printf("SaveConfig Err: %d, %v\n", err)
		return err
	}
	return err
}
