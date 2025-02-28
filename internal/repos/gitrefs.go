package repos

import (
  "fmt"
  "calvu/internal/git"
  "calvu/internal/calver"
  "calvu/internal/models"
)

type Git struct {}

var GitRefs = Git{}

func (g *Git) CurrentVersion() (*models.Version, error) {
  headTime, err := git.HeadDate()
  if err != nil{
    fmt.Println(err)
    return nil, err
  }

  v := calver.FromTime(*headTime)
  v.Patch = git.LatestPatch(v.Major, v.Minor)

  return v, nil
}
