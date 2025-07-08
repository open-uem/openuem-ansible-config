package ansible

const (
	Absent  = "absent"
	Present = "present"
)

type AnsiblePlaybooks struct {
	Playbooks []AnsiblePlaybook
}

type AnsiblePlaybook struct {
	Name       string `yaml:"name"`
	Hosts      string `yaml:"hosts"`
	Connection string `yaml:"connection"`
	Tasks      []any  `yaml:"tasks"`
}

type AnsibleBuiltinGroup struct {
	TaskName   string                        `yaml:"name"`
	Parameters AnsibleBuiltinGroupParameters `yaml:"ansible.builtin.group"`
}

type AnsibleBuiltinGroupParameters struct {
	Name   string `yaml:"name"`
	GID    int    `yaml:"gid,omitempty"`
	System bool   `yaml:"system,omitempty"`
	Force  bool   `yaml:"force,omitempty"`
	State  string `yaml:"state,omitempty"`
}

type AnsibleBuiltinAddUser struct {
	TaskName   string                          `yaml:"name"`
	Parameters AnsibleBuiltinAddUserParameters `yaml:"ansible.builtin.user"`
}

type AnsibleBuiltinAddUserParameters struct {
	Name                         string  `yaml:"name"`
	Append                       bool    `yaml:"append,omitempty"`
	Comment                      string  `yaml:"comment,omitempty"`
	CreateHome                   bool    `yaml:"create_home,omitempty"`
	Expires                      float64 `yaml:"expires,omitempty"`
	Force                        bool    `yaml:"force,omitempty"`
	GenerateSSHKey               bool    `yaml:"generate_ssh_key,omitempty"`
	Group                        string  `yaml:"group,omitempty"`
	Groups                       string  `yaml:"groups,omitempty"`
	Hidden                       bool    `yaml:"hidden,omitempty"`
	Home                         string  `yaml:"home,omitempty"`
	NonUnique                    bool    `yaml:"non_unique,omitempty"`
	Password                     string  `yaml:"password,omitempty"`
	PasswordExpireAccountDisable int     `yaml:"password_expire_account_disable,omitempty"`
	PasswordExpireMax            int     `yaml:"password_expire_max,omitempty"`
	PasswordExpireMin            int     `yaml:"password_expire_min,omitempty"`
	PasswordExpireWarn           int     `yaml:"password_expire_warn,omitempty"`
	PasswordLock                 bool    `yaml:"password_lock"`
	SeUser                       string  `yaml:"seuser,omitempty"`
	Shell                        string  `yaml:"shell,omitempty"`
	Skeleton                     string  `yaml:"skeleton,omitempty"`
	SSHKeyBits                   int     `yaml:"ssh_key_bits,omitempty"`
	SSHKeyComment                string  `yaml:"ssh_key_comment,omitempty"`
	SSHKeyFile                   string  `yaml:"ssh_key_file,omitempty"`
	SSHKeyPassphrase             string  `yaml:"ssh_key_passphrase,omitempty"`
	SSHKeyType                   string  `yaml:"ssh_key_type,omitempty"`
	State                        string  `yaml:"state,omitempty"`
	System                       bool    `yaml:"system,omitempty"`
	UID                          int     `yaml:"uid,omitempty"`
	UIDMax                       int     `yaml:"uid_max,omitempty"`
	UIDMin                       int     `yaml:"uid_min,omitempty"`
	Umask                        string  `yaml:"umask,omitempty"`
	UpdatePassword               string  `yaml:"update_password,omitempty"`
}

type AnsibleBuiltinRemoveUser struct {
	TaskName   string                             `yaml:"name"`
	Parameters AnsibleBuiltinRemoveUserParameters `yaml:"ansible.builtin.user"`
}

type AnsibleBuiltinRemoveUserParameters struct {
	Name  string `yaml:"name"`
	Force bool   `yaml:"force,omitempty"`
	State string `yaml:"state,omitempty"`
}

type AnsibleBuiltinShell struct {
	TaskName string                  `yaml:"name"`
	Shell    string                  `yaml:"ansible.builtin.shell"`
	Args     AnsibleBuiltinShellArgs `yaml:"args,omitempty"`
}

type AnsibleBuiltinShellArgs struct {
	Executable string `yaml:"executable,omitempty"`
	Creates    string `yaml:"creates,omitempty"`
}

type CommunityGeneralFlatpak struct {
	TaskName   string                            `yaml:"name"`
	Parameters CommunityGeneralFlatpakParameters `yaml:"community.general.flatpak"`
}

type CommunityGeneralFlatpakParameters struct {
	Executable     string `yaml:"executable,omitempty"`
	Method         string `yaml:"method,omitempty"`
	Name           string `yaml:"name,omitempty"`
	NoDependencies bool   `yaml:"no_dependencies,omitempty"`
	Remote         string `yaml:"remote,omitempty"`
	State          string `yaml:"state,omitempty"`
}

type CommunityGeneralHomeBrew struct {
	TaskName   string                             `yaml:"name"`
	Parameters CommunityGeneralHomeBrewParameters `yaml:"community.general.homebrew"`
}

type CommunityGeneralHomeBrewParameters struct {
	ForceFormula   bool   `yaml:"force_formula,omitempty"`
	InstallOptions string `yaml:"install_options,omitempty"`
	Name           string `yaml:"name,omitempty"`
	Path           string `yaml:"path,omitempty"`
	State          string `yaml:"state,omitempty"`
	UpdateHomeBrew bool   `yaml:"update_homebrew,omitempty"`
	UpgradeAll     bool   `yaml:"upgrade_all,omitempty"`
	UpgradeOptions string `yaml:"upgrade_options,omitempty"`
}

type CommunityGeneralHomeBrewCask struct {
	TaskName   string                                 `yaml:"name"`
	Parameters CommunityGeneralHomeBrewCaskParameters `yaml:"community.general.homebrew_cask"`
}

type CommunityGeneralHomeBrewCaskParameters struct {
	AccepExternalApps bool   `yaml:"accept_external_apps,omitempty"`
	Greedy            bool   `yaml:"greedy,omitempty"`
	InstallOptions    string `yaml:"install_options,omitempty"`
	Name              string `yaml:"name,omitempty"`
	Path              string `yaml:"path,omitempty"`
	State             string `yaml:"state,omitempty"`
	SudoPassword      string `yaml:"sudo_password,omitempty"`
	UpdateHomeBrew    bool   `yaml:"update_homebrew,omitempty"`
	UpgradeAll        bool   `yaml:"upgrade_all,omitempty"`
}
