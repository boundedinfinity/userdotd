package config

import (
	"path"
)

const (
	fishLoginAvailableComponents = "fish/login.d/available"
	fishLoginEnabledComponents   = "fish/login.d/enabled"
	bashLoginAvailableComponents = "bash/login.d/available"
	bashLoginEnabledComponents   = "bash/login.d/enabled"
)

func GetContentPath(elem ...string) string {
	// dir, err := os.Getwd()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	t := []string{"content"}

	for _, e := range elem {
		t = append(t, e)
	}

	return path.Join(t...)
}

func GetBashPath(elem ...string) string {
	t := []string{GetBashDir()}
	t = append(t, elem...)
	return path.Join(t...)
}

func GetFishPath(elem ...string) string {
	t := []string{GetFishDir()}
	t = append(t, elem...)
	return path.Join(t...)
}

func GetUserPath(elem ...string) string {
	t := []string{GetUserDir()}
	t = append(t, elem...)
	return path.Join(t...)
}

func GetAvailablePath(elem ...string) string {
	if IsFishShell() {
		return GetFishLoginAvailablePath(elem...)
	}

	if IsBashShell() {
		return GetBashLoginAvailablePath(elem...)
	}

	return ""
}

func GetFishLoginAvailablePath(elem ...string) string {
	t := []string{fishLoginAvailableComponents}
	t = append(t, elem...)
	return GetuserdotdPath(t...)
}

func GetEnabledPath(elem ...string) string {
	if IsFishShell() {
		return GetFishLoginEnabledPath(elem...)
	}

	if IsBashShell() {
		return GetBashLoginEnabledPath(elem...)
	}

	return ""
}

func GetFishLoginEnabledPath(elem ...string) string {
	t := []string{fishLoginEnabledComponents}
	t = append(t, elem...)
	return GetuserdotdPath(t...)
}

func GetBashLoginAvailablePath(elem ...string) string {
	t := []string{fishLoginAvailableComponents}
	t = append(t, elem...)
	return GetuserdotdPath(t...)
}

func GetBashLoginEnabledPath(elem ...string) string {
	t := []string{fishLoginEnabledComponents}
	t = append(t, elem...)
	return GetuserdotdPath(t...)
}

func GetuserdotdPath(elem ...string) string {
	t := []string{GetuserdotdDir()}
	t = append(t, elem...)
	return path.Join(t...)
}
