package models

import (
	"log"
	"sync"

	"github.com/beego/beego/v2/client/orm"
)

type Config struct {
	SelectedSemester   string
	SelectedSchedule   int
	HiddenPastSemester bool
}

type ConfigWithLock struct {
	Config
	sync.RWMutex
}

var GlobalConfig ConfigWithLock

func init() {
	var err error
	err = GlobalConfig.ReadFromDB()
	if err == orm.ErrNoRows {
		GlobalConfig.Config = Config{
			SelectedSemester: "",
			SelectedSchedule: 0,
		}
		err = GlobalConfig.SaveToDB()
	}
	if err != nil {
		panic(err)
	}
}

type ConfigInDb struct {
	Id int `orm:"pk;default(1)"`
	Config
}

func (l *ConfigWithLock) ReadFromDB() error {
	c := ConfigInDb{
		Id: 1,
		Config: Config{
			SelectedSemester: "",
			SelectedSchedule: 0,
		},
	}
	err := o.Read(&c, "id")
	if err != nil {
		log.Printf("GetConfig Err: %d, %v\n", err)
		return err
	}
	GlobalConfig.Lock()
	GlobalConfig.Config = c.Config
	GlobalConfig.Unlock()
	return nil
}

func (l *ConfigWithLock) SaveToDB() error {
	GlobalConfig.RLock()
	wrap := ConfigInDb{
		Id:     1,
		Config: GlobalConfig.Config,
	}
	GlobalConfig.RUnlock()
	_, err := o.InsertOrUpdate(&wrap)
	if err != nil {
		log.Printf("SaveConfig Err: %d, %v\n", err)
		return err
	}
	return err
}
