// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package aclog

var (
	// GitCommit is the git commit that was compiled. This will be filled in by
	// the compiler.
	GitCommit string

	// GitDescribe is the git description that was compiled. This will be filled
	// in by the compiler.
	GitDescribe string
)

const (
	// Version is the semantic version number being executed.
	Version = "1.1.4"

	// VersionPrerelease marks the pre-release version. If this is ""
	// (empty string) then it is a final release. Otherwise, this is a pre-release
	// such as "dev" (in development), "beta", "rc1", etc.
	VersionPrerelease = "alpha"
)
