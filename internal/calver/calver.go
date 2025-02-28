package calver

import (
  "time"
  "calvu/internal/models"
)

func FromTime(t time.Time) *models.Version {
  return &models.Version {
    Major: major(t),
    Minor: minor(t),
  }
}

func Bump(v models.Version) models.Version {
  p := v.Patch + 1
  return models.Version {
    Major: v.Major,
    Minor: v.Minor,
    Patch: p,
  }
}

func major(t time.Time) int {
  m := t.Year() - 2000
  return m
}

func minor(t time.Time) int {
  return int(t.Month())
}
