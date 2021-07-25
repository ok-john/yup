// Copyright (C) 2021 light-river, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package yup

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

const (
	DefaultTimeout int = 60
)

type Artifact struct {
	URL        string
	Compressed string
	RootfsPath string
}

func FromCLI(rootFsPath string) Artifact {
	return Artifact{
		URL:        "https://cloud-images.ubuntu.com/focal/current/focal-server-cloudimg-amd64-root.tar.xz",
		RootfsPath: fmt.Sprintf("/home/%v/.yup/fs", os.Getenv("USER")),
		Compressed: fmt.Sprintf("/home/%v/.yup/yupfs.tar.xz", os.Getenv("USER")),
	}
}

type Session struct {
	ToExec  string
	SubCmd  string
	SubArgs []string
	Depth   int
	FS      Artifact
}

func (k Session) CreateSubCmd() *exec.Cmd {
	return LinkIO(exec.Command(k.SubCmd, k.SubArgs...))
}

func (k Session) Marshall() []string {
	return []string{
		"-k", fmt.Sprintf("%v", k.ToExec),
		"-f", k.FS.RootfsPath,
		"-d", strconv.Itoa(k.Depth),
		"-c", k.SubCmd,
		"-a", strings.Join(k.SubArgs, " "),
	}
}

func UnMarshall() Session {
	yupFlag := flag.String("k", "-h", "The Yup command you want to run (run, child)")
	subCmdFlag := flag.String("c", "", "The command to execute in the container")
	subArgsFlag := flag.String("a", "", "The args for the command, space seperated")
	recursionDepthFlag := flag.Int("d", 1, "Recursion depth, by default 1")
	mntFlag := flag.String("f", fmt.Sprintf("/home/%v/roots/rootfs", os.Getenv("USER")), "URL to pull a compressed rootfs from")
	fileSystem := FromCLI(*mntFlag)
	flag.Parse()
	subArgs := []string{}
	if len(*subArgsFlag) >= 1 {
		subArgs = strings.Split(*subArgsFlag, " ")
	}
	return Session{
		ToExec:  *yupFlag,
		SubCmd:  *subCmdFlag,
		SubArgs: subArgs,
		Depth:   *recursionDepthFlag,
		FS:      fileSystem,
	}
}

// Unmounts Proc
func (k Session) Unmount() error {
	return syscall.Unmount("proc", 0)
}

// Recursively calls the current process from
// /proc/self/exe
func (k Session) Nest(c *exec.Cmd) *exec.Cmd {
	k.ToExec = "child"
	k.Depth--
	return Contain(exec.Command("/proc/self/exe", k.Marshall()...))
}

// Changes root to the root specificed in the session
// Mounts /proc
func (k Session) Chmount() error {
	syscall.Chroot(k.FS.RootfsPath)
	os.Chdir("/")
	syscall.Mount("proc", "proc", "proc", 0, "")
	return nil
}

// Wrapes any *exec.Cmd in it's own linux namespace
func Contain(cmd *exec.Cmd) *exec.Cmd {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER |
			syscall.CLONE_NEWNS | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWNET,
		UidMappings: []syscall.SysProcIDMap{{ContainerID: 0, HostID: os.Getuid(), Size: 1}},
		GidMappings: []syscall.SysProcIDMap{{ContainerID: 0, HostID: os.Getgid(), Size: 1}},
	}
	return LinkIO(cmd)
}

func LinkIO(cmd *exec.Cmd) *exec.Cmd {
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
