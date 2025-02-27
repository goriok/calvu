package version

import (
  "time"
  "fmt"
)

func FromTime(t time.Time) string{
	y := t.Year() - 2000

	return fmt.Sprintf("%02d.%d.%02d%02d%02d", 
    y, 
    t.Month(), 
    t.Hour(), 
    t.Minute(), 
    t.Second(),
  )
}

