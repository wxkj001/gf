// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

// 进程管理/通信.
// 本进程管理从syscall, os.StartProcess, exec.Cmd都使用过，
// 最后采用了exec.Cmd来实现多进程管理，这是一个顶层的跨平台封装，兼容性更好，另外两个是偏底层的接口。
package gproc

import (
    "os"
    "gitee.com/johng/gf/g/util/gconv"
)

const (
    gPROC_ENV_KEY_PPID_KEY = "gproc.ppid"
)

// 获取当前进程ID
func Pid() int {
    return os.Getpid()
}

// 获取父进程ID(gproc父进程，不存在时则使用系统父进程)
func PPid() int {
    // gPROC_ENV_KEY_PPID_KEY为gproc包自定义的父进程
    ppidValue := os.Getenv(gPROC_ENV_KEY_PPID_KEY)
    if ppidValue != "" {
        return gconv.Int(ppidValue)
    }
    return os.Getppid()
}

// 获取父进程ID(系统父进程)
func PpidOfOs() int {
    return os.Getppid()
}

// 判断当前进程是否为gproc创建的子进程
func IsChild() bool {
    return os.Getenv(gPROC_ENV_KEY_PPID_KEY) != ""
}
