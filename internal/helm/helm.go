package helm

import (
  "fmt"
  "os"
  "gopkg.in/yaml.v3"
  "calvu/internal/models"
)
const ChartFile = "Chart.yaml"

func SetVersion(version string) (*models.Chart, error){
	data, err := os.ReadFile(ChartFile)
	if err != nil {
		return nil, err
	}

	var chart models.Chart
	err = yaml.Unmarshal(data, &chart)
	if err != nil {
		return nil, err
	}

	fmt.Println("Current version:", chart.Version)
	chart.Version = version
	updatedData, err := yaml.Marshal(&chart)
	if err != nil {
		return nil, err
	}

	err = os.WriteFile(ChartFile, updatedData, 0644)
	if err != nil {
		return nil, err
	}

	fmt.Println("Updated version:", chart.Version)
  return &chart, nil
}

