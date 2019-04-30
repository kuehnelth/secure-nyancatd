package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
	"unsafe"

	"github.com/gliderlabs/ssh"
	"github.com/kr/pty"
)

var (
	port            = flag.Int("port", 22, "SSH server port")
	hostKeyFilePath = flag.String("host-key", "/etc/secure-nyancatd/ssh_host_rsa_key", "ID RSA SSH Host key")
)

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func main() {
	flag.Parse()

	ssh.Handle(func(s ssh.Session) {
		log.Println("New SSH session from " + s.RemoteAddr().String())
		cmd := exec.Command("nyancat")
		ptyReq, winCh, isPty := s.Pty()
		if isPty {
			cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
			f, err := pty.Start(cmd)
			start := time.Now()
			if err != nil {
				panic(err)
			}
			go func() {
				for win := range winCh {
					setWinsize(f, win.Width, win.Height)
				}
			}()
			go func() {
				io.Copy(f, s) // stdin
			}()
			io.Copy(s, f) // stdout
			cmd.Wait()
			elapsed := time.Since(start)
			log.Printf("%s nyaned for %f seconds", s.RemoteAddr().String(), elapsed.Seconds())
		} else {
			io.WriteString(s, "No PTY requested.\n")
			s.Exit(1)
		}
	})

	portString := strconv.Itoa(*port)
	log.Println("starting nyancat server on port " + portString)
	log.Fatal(ssh.ListenAndServe(":" + portString, nil, ssh.HostKeyFile(*hostKeyFilePath)))
}
