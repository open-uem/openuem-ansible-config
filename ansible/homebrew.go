package ansible

import (
	"errors"
)

func InstallHomeBrewFormula(taskName string, name string, installOptions string, updateHomeBrew bool) (*CommunityGeneralHomeBrew, error) {
	f := CommunityGeneralHomeBrew{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName

	f.Parameters = CommunityGeneralHomeBrewParameters{}
	if name == "" {

		return nil, errors.New("package/formula name cannot be empty")
	}

	f.Parameters.Name = name

	if installOptions != "" {
		f.Parameters.InstallOptions = installOptions
	}

	f.Parameters.UpdateHomeBrew = updateHomeBrew
	f.Parameters.State = "present"

	return &f, nil
}

func UpgradeHomeBrewFormula(taskName string, name string, updateHomeBrew bool, upgradeAll bool, upgradeOptions string) (*CommunityGeneralHomeBrew, error) {
	f := CommunityGeneralHomeBrew{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName

	f.Parameters = CommunityGeneralHomeBrewParameters{}
	if name == "" {
		if !upgradeAll {
			return nil, errors.New("package/formula name cannot be empty")
		}
	} else {
		f.Parameters.Name = name
	}

	f.Parameters.UpdateHomeBrew = updateHomeBrew
	f.Parameters.UpgradeAll = upgradeAll

	if upgradeOptions != "" {
		f.Parameters.UpgradeOptions = upgradeOptions
	}

	f.Parameters.State = "upgraded"

	return &f, nil
}

func UninstallHomeBrewFormula(taskName string, name string) (*CommunityGeneralHomeBrew, error) {
	f := CommunityGeneralHomeBrew{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName

	f.Parameters = CommunityGeneralHomeBrewParameters{}
	if name == "" {
		return nil, errors.New("package/formula name cannot be empty")
	}
	f.Parameters.Name = name
	f.Parameters.State = "absent"

	return &f, nil
}

func InstallHomeBrewCask(taskName string, name string, installOptions string, updateHomeBrew bool) (*CommunityGeneralHomeBrewCask, error) {
	f := CommunityGeneralHomeBrewCask{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName

	f.Parameters = CommunityGeneralHomeBrewCaskParameters{}
	if name == "" {

		return nil, errors.New("cask name cannot be empty")
	}
	f.Parameters.Name = name

	if installOptions != "" {
		f.Parameters.InstallOptions = installOptions
	}

	f.Parameters.UpdateHomeBrew = updateHomeBrew
	f.Parameters.State = "present"

	return &f, nil
}

func UpgradeHomeBrewCask(taskName string, name string, greedy bool, updateHomeBrew bool, upgradeAll bool) (*CommunityGeneralHomeBrewCask, error) {
	f := CommunityGeneralHomeBrewCask{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName

	f.Parameters = CommunityGeneralHomeBrewCaskParameters{}
	if name == "" {
		if !upgradeAll {
			return nil, errors.New("cask name cannot be empty")
		}
	} else {
		f.Parameters.Name = name
	}

	f.Parameters.Greedy = greedy
	f.Parameters.UpdateHomeBrew = updateHomeBrew
	f.Parameters.UpgradeAll = upgradeAll

	f.Parameters.State = "upgraded"

	return &f, nil
}

func UninstallHomeBrewCask(taskName string, name string) (*CommunityGeneralHomeBrewCask, error) {
	f := CommunityGeneralHomeBrewCask{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName

	f.Parameters = CommunityGeneralHomeBrewCaskParameters{}
	if name == "" {
		return nil, errors.New("cask name cannot be empty")
	}
	f.Parameters.Name = name
	f.Parameters.State = "absent"

	return &f, nil
}
