package ansible

import (
	"errors"
	"fmt"
	"strconv"
)

func AddLinuxLocalUser(taskName string,
	append bool, comment string, create_home bool, expires float64, force bool,
	generate_ssh_key bool, group string, groups string, home string, name string,
	non_unique bool, password string, password_expire_account_disable int,
	password_expire_max int, password_expire_min int, password_expire_warn int,
	password_lock bool, shell string, skeleton string, ssh_key_bits int,
	ssh_key_comment string, ssh_key_file string, ssh_key_passphrase string,
	ssh_key_type string, system bool, umask string,
	uid int, uid_max int, uid_min int) (*AnsibleBuiltinUser, error) {

	user := AnsibleBuiltinUser{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	user.TaskName = taskName

	if name == "" {
		return nil, errors.New("username cannot be empty")
	}

	user.Parameters = AnsibleBuiltinUserParameters{
		Name:           name,
		Comment:        comment,
		CreateHome:     strconv.FormatBool(create_home),
		Append:         append,
		System:         system,
		Force:          force,
		GenerateSSHKey: generate_ssh_key,
		Groups:         groups,
		NonUnique:      non_unique,
		PasswordLock:   password_lock,
		State:          "present",
		UpdatePassword: "on_create",
	}

	if expires >= 0 {
		user.Parameters.Expires = expires
	}

	if group != "" {
		user.Parameters.Group = group
	}

	if home != "" {
		user.Parameters.Home = home
	}

	if password != "" {
		user.Parameters.Password = fmt.Sprintf("{{ '%s' | password_hash('sha512') }}", password)
	}

	if password_expire_account_disable >= 0 {
		user.Parameters.PasswordExpireAccountDisable = password_expire_account_disable
	}

	if password_expire_max >= 0 {
		user.Parameters.PasswordExpireMax = password_expire_max
	}

	if password_expire_min >= 0 {
		user.Parameters.PasswordExpireMin = password_expire_min
	}

	if password_expire_warn >= 0 {
		user.Parameters.PasswordExpireWarn = password_expire_warn
	}

	if shell != "" {
		user.Parameters.Shell = shell
	}

	if skeleton != "" {
		if !create_home {
			return nil, errors.New("create_home must set to use skeleton")
		}
		user.Parameters.Skeleton = skeleton
	}

	if ssh_key_bits >= 0 {
		user.Parameters.SSHKeyBits = ssh_key_bits
	}

	if ssh_key_comment != "" {
		user.Parameters.SSHKeyComment = ssh_key_comment
	}

	if ssh_key_file != "" {
		user.Parameters.SSHKeyFile = ssh_key_file
	}

	if ssh_key_passphrase != "" {
		user.Parameters.SSHKeyPassphrase = ssh_key_passphrase
	}

	if ssh_key_type != "" {
		user.Parameters.SSHKeyType = ssh_key_type
	}

	if umask != "" {
		user.Parameters.Umask = umask
	}

	if uid >= 0 {
		user.Parameters.UID = uid
	}

	if uid_max >= 0 {
		user.Parameters.UIDMax = uid_max
	}

	if uid_min >= 0 {
		user.Parameters.UIDMin = uid_min
	}

	return &user, nil
}
