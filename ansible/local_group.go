package ansible

import "errors"

func AddLocalGroup(name string, gid int, system bool) (*AnsiblePlaybook, error) {
	playbook := AnsiblePlaybook{}

	group := AnsibleBuiltinGroup{}

	if name == "" {
		return nil, errors.New("group name cannot be empty")
	}
	group.Name = name
	if gid > 0 {
		group.GID = gid
	}
	group.System = system
	group.State = Present

	playbook.Tasks = append(playbook.Tasks, group)

	return &playbook, nil
}

func RemoveLocalGroup(name string, force bool) (*AnsiblePlaybook, error) {
	playbook := AnsiblePlaybook{}

	group := AnsibleBuiltinGroup{}

	if name == "" {
		return nil, errors.New("group name cannot be empty")
	}
	group.Name = name
	group.Force = force
	group.State = Absent

	playbook.Tasks = append(playbook.Tasks, group)

	return &playbook, nil
}
