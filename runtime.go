// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package aclog

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	runtimeDocker       = "docker"
	runtimeRkt          = "rkt"
	runtimeRunC         = "runc"
	runtimeContainerD   = "containerd"
	runtimeLXC          = "lxc"
	runtimeLXD          = "lxd"
	runtimeOpenVZ       = "openvz"
	runtimeUndetermined = "undetermined"
)

func getRuntime() string {
	if _, err := os.Stat("/.dockerinit"); err == nil {
		return runtimeDocker
	}

	if _, err := os.Stat("/.dockerenv"); err == nil {
		return runtimeDocker
	}

	cgroup, _ := ioutil.ReadFile("/proc/self/cgroup")
	if strings.Contains(string(cgroup), "docker") {
		return runtimeDocker
	}

	if ac := os.Getenv("AC_METADATA_URL"); ac != "" {
		return runtimeRkt
	}

	if ac := os.Getenv("AC_APP_NAME"); ac != "" {
		return runtimeRkt
	}

	return runtimeUndetermined
}