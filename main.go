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
package main

import (
	"log"
	"os"
	"os/exec"
	"yup/yup"
)

var (
	CMD_FUNCTIONS map[string]func(yup.Session) = map[string]func(yup.Session){
		"pull":  pull,
		"run":   run,
		"child": child,
	}
)

func main() {
	sesh := yup.UnMarshall()
	as("\t+main")
	as("\t^->+" + sesh.ToExec)
	CMD_FUNCTIONS[sesh.ToExec](sesh)
}

func pull(sesh yup.Session) {
	nilOr(exec.Command(".env/pull").Run())
}

func run(sesh yup.Session) {
	nilOr(sesh.Nest(sesh.CreateSubCmd()).Run())
}

func child(sesh yup.Session) {
	nilOr(sesh.Chmount())
	nilOr(sesh.CreateSubCmd().Run())
	nilOr(sesh.Unmount())
}

func nilOr(err error) {
	as("+>nilOr")
	if err != nil {
		panic(err)
	}
}

func as(f string) {
	log.Printf("%v(uid: %v, gid: %v, ppid:%v, pid: %v)", f, os.Getuid(), os.Getgid(), os.Getppid(), os.Getpid())
}
