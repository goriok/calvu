// cli tool for semantic versioning
package main

import (
	"github.com/sirupsen/logrus"
	"svcli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
