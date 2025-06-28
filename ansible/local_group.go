package ansible

import "errors"

func AddLocalGroup(playbookName string, name string, gid int, system bool) (*AnsiblePlaybook, error) {
	playbook := AnsiblePlaybook{}
	if playbookName == "" {
		return nil, errors.New("playbook name cannot be empty")
	}
	playbook.Name = playbookName

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

func RemoveLocalGroup(playbookName string, name string, force bool) (*AnsiblePlaybook, error) {
	playbook := AnsiblePlaybook{}
	if playbookName == "" {
		return nil, errors.New("playbook name cannot be empty")
	}
	playbook.Name = playbookName

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
