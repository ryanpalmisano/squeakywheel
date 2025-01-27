package main

import (
        "fmt"
        "log"
        "net"
        "os"
        "os/exec"
        "path/filepath"
        "strings"
        "syscall"
)

func main() {
        // 1. Attempt to read and write to dummy file
        dummyFile := "/root/ring0_rights/user_group/dummy_file"
        err := os.OpenFile(dummyFile, os.O_RDWR, 0644)
        if err != nil {
                log.Printf("Error accessing dummy file: %v", err)
        } else {
                log.Println("Successfully accessed dummy file.")
        }

        // 2. Check file system permissions
        err = filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
                if err != nil {
                        return err
                }

                if info.Mode()&0777 != 0755 && !strings.HasPrefix(path, "/proc") && !strings.HasPrefix(path, "/sys") {
                        log.Printf("Incorrect permissions: %s (Mode: %o)", path, info.Mode())
                }
                return nil
        })
        if err != nil {
                log.Fatal(err)
        }

        // 3. Attempt remote shell connection
        remoteAddr, err := net.ResolveTCPAddr("tcp", "192.168.1.100:22") // Replace with target IP
        if err != nil {
                log.Fatal("Error resolving remote address:", err)
        }

        cmd := exec.Command("ssh", "root@"+remoteAddr.IP.String())
        err = cmd.Run()
        if err != nil {
                log.Printf("SSH connection failed: %v", err)
        } else {
                log.Println("SSH connection successful.")
        }

        // 4. Attempt firewall test (Example: ping)
        cmd = exec.Command("ping", "-c", "1", "192.168.1.100") // Replace with target IP
        err = cmd.Run()
        if err != nil {
                log.Printf("Ping failed: %v", err)
        } else {
                log.Println("Ping successful.")
        }
}
