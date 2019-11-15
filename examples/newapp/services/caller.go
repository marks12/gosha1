package services

import "runtime"

func GetCallerFunction() string {

    pc := make([]uintptr, 10)
    runtime.Callers(3, pc)
    f := runtime.FuncForPC(pc[0])
    return f.Name()
}

