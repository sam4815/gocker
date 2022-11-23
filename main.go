package main

import (
  "os"
  "fmt"
  "os/exec"
  "syscall"
)

func main() {
  switch os.Args[1] {
  case "run":
    parent()
  case "child":
    child()
  default:
    panic("Unsupported argument")
  }
}

func parent() {
  // The first argument refers to the currently-executing program, so
  // this command self-executes the program with "child" as an argument,
  // which provides an isolated UTS (hostname/domain name) namespace
  cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  // Create our own UTS, PID and MNT namespaces
  cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
  }

  if err := cmd.Run(); err != nil {
    fmt.Println("Error running parent command:", err)
    os.Exit(1)
  }
}

func child() {
  syscall.Sethostname([]byte("container"))
  must(syscall.Chdir("/"))
  // Mount a new filesystem; see https://man7.org/linux/man-pages/man2/mount.2.html
  must(syscall.Mount("proc", "proc", "proc", 0, ""))

  cmd := exec.Command(os.Args[2], os.Args[3:]...)
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  if err := cmd.Run(); err != nil {
    fmt.Println("Error running child command:", err)
    os.Exit(1)
  }
}

func must(err error) {
  if err != nil {
    panic(err)
  }
}

