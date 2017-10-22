package main

import (
	"github.com/slawek87/GOBrokers/realEstate"
	"github.com/slawek87/GOBrokers/flatHouse"
)

// Put here all methods and functions which you need before run main().
func init()  {
	realEstate.InitMigrations()
	flatHouse.InitMigrations()
}

func main()  {
    println("Starting!")
}
