package ansible

import (
	"errors"
)

func ExecuteScript(taskName string, shell string, executable string, creates string, agent string) (*AnsibleBuiltinShell, error) {
	builtinShell := AnsibleBuiltinShell{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	builtinShell.TaskName = taskName
	builtinShell.Args.Creates = creates

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
