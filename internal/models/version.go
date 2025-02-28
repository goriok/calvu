package models

import (
  "fmt"
)

type Version struct {
  Major int
  Minor int
  Patch int
}

func (v *Version) String() string {
  return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
