package ansible

import (
	"errors"
)

func InstallFlatpakPackage(taskName string, name string, latest bool) (*CommunityGeneralFlatpak, error) {
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

	return &f, nil
}

func UninstallFlatpakPackage(taskName string, name string) (*CommunityGeneralFlatpak, error) {
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

	return &f, nil
}
