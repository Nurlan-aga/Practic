package config

import (
	"fmt"
	"os"
	"strconv"
	"testPackages/defaults"
)

var Api *Config

type Counter struct {
	Interval int
}
type Log struct {
	Driver string
}

type Config struct {
	Counter *Counter
	Log     *Log
}

func init() {
	driver := os.Getenv("LOG_DRIVER")
	if driver == "" {
		fmt.Println("using default driver for logger ", defaults.DEFAULT_LOG_DRIVER)
		driver = defaults.DEFAULT_LOG_DRIVER
	}

	c, err := strconv.Atoi(os.Getenv("INTERVAL"))
	if err != nil {
		fmt.Println(err)
		fmt.Println("using default value 1")
		c = defaults.DEFAULT_INTERVAL
	}

	fmt.Println("config initialized successful")

	Api = &Config{Counter: &Counter{Interval: c}, Log: &Log{Driver: driver}}
}