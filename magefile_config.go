//go:build mage
// +build mage

package main

const (
	// Path for the sdfsd executable location
	installPath = "/opt/sdfs"

	// Enable installing systemd unit file
	enableSystemdUnitFile = false

	// Enable installing init.d file
	enableInitdFile = false
)
