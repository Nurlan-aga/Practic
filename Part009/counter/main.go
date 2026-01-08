package counter

import (
	"fmt"
	cfg "testPackages/config"
)

func init() {
	Interval = cfg.Api.Counter.Interval
	fmt.Println("counter package initialized")
}