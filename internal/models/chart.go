package models

type Chart struct {
	APIVersion   string   `yaml:"apiVersion"`
	Name         string   `yaml:"name"`
	Version      string   `yaml:"version"`
	KubeVersion  *string  `yaml:"kubeVersion,omitempty"`
	Description  *string  `yaml:"description,omitempty"`
	Type         *string  `yaml:"type,omitempty"`
	Keywords     []string `yaml:"keywords,omitempty"`
	Home         *string  `yaml:"home,omitempty"`
	Sources      []string `yaml:"sources,omitempty"`
	Dependencies []struct {
		Name       string   `yaml:"name"`
		Version    string   `yaml:"version"`
		Repository *string  `yaml:"repository,omitempty"`
		Condition  *string  `yaml:"condition,omitempty"`
		Tags       []string `yaml:"tags,omitempty"`
		Alias      *string  `yaml:"alias,omitempty"`
	} `yaml:"dependencies,omitempty"`
	Maintainers []struct {
		Name  string  `yaml:"name"`
		Email *string `yaml:"email,omitempty"`
		URL   *string `yaml:"url,omitempty"`
	} `yaml:"maintainers,omitempty"`
	Icon        *string           `yaml:"icon,omitempty"`
	AppVersion  *string           `yaml:"appVersion,omitempty"`
	Deprecated  bool              `yaml:"deprecated,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
}
