package main

import (
	"github.com/jinzhu/gorm"
)


type runlog struct {
  gorm.Model
}

func newRunLog() *runlog {
  //db.CreateTable(&runlog{})
  rl := new(runlog);
  db.Create(rl)
  return rl;
}

func (rl *runlog) Id() (uint) {
  return rl.Model.ID;
}
