package ansible

import "errors"

func ExecuteScript(taskName string, shell string, executable string) (*AnsibleBuiltinShell, error) {
	builtinShell := AnsibleBuiltinShell{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	builtinShell.TaskName = taskName

	if shell == "" {
		return nil, errors.New("shell script cannot be empty")
	}
	builtinShell.Shell = shell

	if executable != "" {
		builtinShell.Args = AnsibleBuiltinShellArgs{
			Executable: executable,
		}
	}

	return &builtinShell, nil
}
