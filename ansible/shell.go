package ansible

import (
	"errors"
	"path/filepath"
)

func ExecuteScript(taskName string, shell string, executable string, run string, agent string) (*AnsibleBuiltinShell, error) {
	builtinShell := AnsibleBuiltinShell{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	builtinShell.TaskName = taskName

	if run != "once" && run != "always" {
		return nil, errors.New("wrong run option")
	}

	if run == "once" {
		if agent == "linux" {
			builtinShell.Args.Creates = filepath.Join("/opt/openuem-agent/bin/ansible", taskName, ".txt")
		}
	}

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
