package ansible

import (
	"errors"
)

func InstallHomeBrewFormula(taskName string, name string, installOptions string, updateHomeBrew bool, ignore_errors bool) (*CommunityGeneralHomeBrew, error) {
	f := CommunityGeneralHomeBrew{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName
	f.Become = "yes"
	f.BecomeUser = "some_user"

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

	f.IgnoreErrors = ignore_errors

	return &f, nil
}

func UpgradeHomeBrewFormula(taskName string, name string, updateHomeBrew bool, upgradeAll bool, upgradeOptions string, ignore_errors bool) (*CommunityGeneralHomeBrew, error) {
	f := CommunityGeneralHomeBrew{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName
	f.Become = "yes"
	f.BecomeUser = "some_user"

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

	f.IgnoreErrors = ignore_errors

	return &f, nil
}

func UninstallHomeBrewFormula(taskName string, name string, ignore_errors bool) (*CommunityGeneralHomeBrew, error) {
	f := CommunityGeneralHomeBrew{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName
	f.Become = "yes"
	f.BecomeUser = "some_user"

	f.Parameters = CommunityGeneralHomeBrewParameters{}
	if name == "" {
		return nil, errors.New("package/formula name cannot be empty")
	}
	f.Parameters.Name = name
	f.Parameters.State = "absent"

	f.IgnoreErrors = ignore_errors

	return &f, nil
}

func InstallHomeBrewCask(taskName string, name string, installOptions string, updateHomeBrew bool, ignore_errors bool) (*CommunityGeneralHomeBrewCask, error) {
	f := CommunityGeneralHomeBrewCask{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName
	f.Become = "yes"
	f.BecomeUser = "some_user"

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

	f.IgnoreErrors = ignore_errors

	return &f, nil
}

func UpgradeHomeBrewCask(taskName string, name string, greedy bool, updateHomeBrew bool, upgradeAll bool, ignore_errors bool) (*CommunityGeneralHomeBrewCask, error) {
	f := CommunityGeneralHomeBrewCask{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName
	f.Become = "yes"
	f.BecomeUser = "some_user"

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

	f.IgnoreErrors = ignore_errors

	return &f, nil
}

func UninstallHomeBrewCask(taskName string, name string, ignore_errors bool) (*CommunityGeneralHomeBrewCask, error) {
	f := CommunityGeneralHomeBrewCask{}
	if taskName == "" {
		return nil, errors.New("task name cannot be empty")
	}
	f.TaskName = taskName
	f.Become = "yes"
	f.BecomeUser = "some_user"

	f.Parameters = CommunityGeneralHomeBrewCaskParameters{}
	if name == "" {
		return nil, errors.New("cask name cannot be empty")
	}
	f.Parameters.Name = name
	f.Parameters.State = "absent"

	f.IgnoreErrors = ignore_errors

	return &f, nil
}
