package main

import (
	"fmt"
	"golang.org/x/sys/unix"
	"log"
	"net"
	"os"
	"syscall"
)

/*
	var ls net.Listener
	var err error
	if reusePort {
		ls, err = reuseport.Listen(network, addr)
	} else {
		ls, err = net.Listen(network, addr)
	}
	if err != nil {
		return nil, err
	}

	l, ok := ls.(*net.TCPListener)
	if !ok {
		return nil, errors.New("could not get file descriptor")
	}

	file, err := l.File()
	if err != nil {
		return nil, err
	}
	fd := int(file.Fd())
	if err = unix.SetNonblock(fd, true); err != nil {
		return nil, err
	}

	loop, err := eventloop.New()
	if err != nil {
		return nil, err
	}

	listener := &listener{
		file:     file,
		fd:       fd,
		handleC:  handlerConn,
		listener: ls,
		loop:     loop,
	}
	if err = loop.AddSocketAndEnableRead(fd, listener); err != nil {
		return nil, err
	}
*/


/*
	fd, err := unix.EpollCreate1(0)
	if err != nil {
		return nil, err
	}

	r0, _, errno := unix.Syscall(unix.SYS_EVENTFD2, 0, 0, 0)
	if errno != 0 {
		_ = unix.Close(fd)
		return nil, errno
	}
	eventFd := int(r0)

	err = unix.EpollCtl(fd, unix.EPOLL_CTL_ADD, eventFd, &unix.EpollEvent{
		Events: unix.EPOLLIN,
		Fd:     int32(eventFd),
	})
	if err != nil {
		_ = unix.Close(fd)
		_ = unix.Close(eventFd)
		return nil, err
	}

*/

func echo(fd int) {
	defer syscall.Close(fd)
	var buf [32 * 1024]byte
	for {
		nbytes, e := syscall.Read(fd, buf[:])
		if nbytes > 0 {
			fmt.Printf(">>> %s", buf)
			syscall.Write(fd, buf[:nbytes])
			fmt.Printf("<<< %s", buf)
		}
		if e != nil {
			break
		}
	}
}

func getFD() int {
	ln, err := net.Listen("tcp", ":8972")
	if err != nil {
		panic(err)
	}
	l, ok := ln.(*net.TCPListener)
	if !ok {
		log.Fatalf("listener not tcp err")
	}
	file, err := l.File()
	if err != nil {
		log.Fatalf("tcp listener to file err %v", err)
	}
	fmt.Printf("file is %s\n", file.Name())
	fd := int(file.Fd())
	if err := unix.SetNonblock(fd, true); err != nil {
		log.Fatalf("set file non block err")
	}
	return fd
}


func main() {
	epFD, err := unix.EpollCreate1(0)
	if err != nil {
		log.Fatalf("epoll create err %v", err)
	}
	fd := getFD()
	err = unix.EpollCtl(epFD, unix.EPOLL_CTL_ADD, fd, &unix.EpollEvent{
		Events: unix.EPOLLIN | unix.EPOLLOUT,
		Fd:     int32(fd),
	})
	for {
		events := make([]unix.EpollEvent, 10)
		eventNum, err := unix.EpollWait(epFD, events,  0)
		if err != nil {
			log.Fatalf("epoll wait err is %v", err)
		}
		for i := 0; i < eventNum; i++ {
			if int(events[i].Fd) == fd {
				connFd, _, err := syscall.Accept(fd)
				if err != nil {
					fmt.Println("accept: ", err)
					continue
				}
				syscall.SetNonblock(fd, true)
				event := new(syscall.EpollEvent)
				event.Events = syscall.EPOLLIN | (1 << 31)
				event.Fd = int32(connFd)
				if err := syscall.EpollCtl(
					epFD,
					syscall.EPOLL_CTL_ADD,
					connFd,
					event,
				); err != nil {
					fmt.Print("epoll_ctl: ", connFd, err)
					os.Exit(1)
				}
			} else {
				go echo(int(events[i].Fd))
			}
		}
	}
}
