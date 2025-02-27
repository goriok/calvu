package cversion

import(
  "calvu/internal/git"
  "fmt"
)

func FromGit() (*Version, error) {
  headTime, err := git.HeadDate()
  if err != nil{
    fmt.Println(err)
    return nil, err
  }

  v := New(*headTime, 0)
  patch := git.LatestPatch(v.Major, v.Minor)
  v.Patch = patch

  return v, nil
}

