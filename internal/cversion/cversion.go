package cversion

import (
  "time"
  "fmt"
)

type Version struct {
  Major int
  Minor int
  Patch int
}

func New(t time.Time, p int) *Version {
  return &Version {
    Major: major(t),
    Minor: minor(t),
    Patch: p,
  }
}

func (v *Version) Value() string {
  return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func (v *Version) Bump(){
    v.Patch += 1
}

func major(t time.Time) int {
  m := t.Year() - 2000
  return m
}

func minor(t time.Time) int {
  return int(t.Month())
}
