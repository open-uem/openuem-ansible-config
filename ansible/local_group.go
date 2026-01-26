package ansible

import "errors"

func AddLocalGroup(taskName string, name string, gid int, system bool, ignore_errors bool) (*AnsibleBuiltinGroup, error) {
	group := AnsibleBuiltinGroup{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	group.TaskName = taskName

	if name == "" {
		return nil, errors.New("group name cannot be empty")
	}

	group.Parameters = AnsibleBuiltinGroupParameters{
		Name: name,
	}

	if gid > 0 {
		group.Parameters.GID = gid
	}
	group.Parameters.System = system
	group.Parameters.State = Present

	group.IgnoreErrors = ignore_errors

	return &group, nil
}

func RemoveLocalGroup(taskName string, name string, force bool, ignore_errors bool) (*AnsibleBuiltinGroup, error) {
	group := AnsibleBuiltinGroup{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	group.TaskName = taskName

	if name == "" {
		return nil, errors.New("group name cannot be empty")
	}

	group.Parameters = AnsibleBuiltinGroupParameters{
		Name: name,
	}

	group.Parameters.Force = force
	group.Parameters.State = Absent

	group.IgnoreErrors = ignore_errors

	return &group, nil
}
