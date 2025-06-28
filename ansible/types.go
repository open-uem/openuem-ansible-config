package ansible

const (
	Absent  = "absent"
	Present = "present"
)

type AnsiblePlaybook struct {
	Name       string `yaml:"name"`
	Hosts      string `yaml:"hosts"`
	Connection string `yaml:"connection"`
	Tasks      []any  `yaml:"tasks"`
}

type AnsibleBuiltinGroup struct {
	Name   string `yaml:"name"`
	GID    int    `yaml:"gid,omitempty"`
	System bool   `yaml:"system,omitempty"`
	Force  bool   `yaml:"force,omitempty"`
	State  string `yaml:"force,omitempty"`
}
