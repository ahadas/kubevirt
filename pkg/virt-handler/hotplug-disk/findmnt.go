package hotplug_volume

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

const (
	rhcosPrefix = "/ostree/deploy/rhcos"
)

var (
	sourceRgx = regexp.MustCompile(`\[(.+)\]`)

	findMntByVolume = func(volumeName string, pid int) ([]byte, error) {
		return exec.Command("/usr/bin/findmnt", "-T", fmt.Sprintf("/%s", volumeName), "-N", strconv.Itoa(pid), "-J").CombinedOutput()
	}

	findMntByDevice = func(deviceName string) ([]byte, error) {
		return exec.Command("/usr/bin/findmnt", "-S", deviceName, "-N", "1", "-J").CombinedOutput()
	}
)

type FindmntInfo struct {
	Target  string `json:"target"`
	Source  string `json:"source"`
	Fstype  string `json:"fstype"`
	Options string `json:"options"`
}

type FileSystems struct {
	Filesystems []FindmntInfo `json:"filesystems"`
}

func LookupFindmntInfoByVolume(volumeName string, pid int) ([]FindmntInfo, error) {
	mntInfoJson, err := findMntByVolume(volumeName, pid)
	if err != nil {
		return make([]FindmntInfo, 0), errors.Wrapf(err, "Error running findmnt for volume %s", volumeName)
	}
	return parseMntInfoJson(mntInfoJson)
}

func LookupFindmntInfoByDevice(deviceName string) ([]FindmntInfo, error) {
	mntInfoJson, err := findMntByDevice(deviceName)
	if err != nil {
		return make([]FindmntInfo, 0), errors.Wrapf(err, "Error running findmnt for device %s", deviceName)
	}
	return parseMntInfoJson(mntInfoJson)
}

func parseMntInfoJson(mntInfoJson []byte) ([]FindmntInfo, error) {
	mntinfo := FileSystems{}
	if err := json.Unmarshal(mntInfoJson, &mntinfo); err != nil {
		return mntinfo.Filesystems, errors.Wrapf(err, "unable to unmarshal [%v]", mntInfoJson)
	}
	return mntinfo.Filesystems, nil
}

func (f *FindmntInfo) GetSource() string {
	match := sourceRgx.FindStringSubmatch(f.Source)
	if len(match) != 2 {
		return strings.TrimPrefix(f.Source, rhcosPrefix)
	}
	return strings.TrimPrefix(match[1], rhcosPrefix)
}

func (f *FindmntInfo) GetOptions() []string {
	return strings.Split(f.Options, ",")
}
