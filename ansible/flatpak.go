package ansible

import (
	"errors"
)

func InstallFlatpakPackage(taskName string, name string, latest bool, ignore_errors bool) (*CommunityGeneralFlatpak, error) {
	f := CommunityGeneralFlatpak{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName

	f.Parameters = CommunityGeneralFlatpakParameters{}
	if name == "" {
		return nil, errors.New("package name cannot be empty")
	}

	f.Parameters.Name = name
	f.Parameters.State = "present"
	if latest {
		f.Parameters.State = "latest"
	}

	f.IgnoreErrors = ignore_errors

	return &f, nil
}

func UninstallFlatpakPackage(taskName string, name string, ignore_errors bool) (*CommunityGeneralFlatpak, error) {
	f := CommunityGeneralFlatpak{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName

	f.Parameters = CommunityGeneralFlatpakParameters{}
	if name == "" {
		return nil, errors.New("package name cannot be empty")
	}

	f.Parameters.Name = name
	f.Parameters.State = "absent"

	f.IgnoreErrors = ignore_errors

	return &f, nil
}
