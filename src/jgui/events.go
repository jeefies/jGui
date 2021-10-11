package jgui

import (
    "bytes"
    "time"
    "fmt"

    "sdl"
)

/*
type Event struct {
    EventName string
    EventCode int
    SDLEvent * sdl.Event
}
*/

const MaxEventCoroutine int = 10

var gEvents = make([](*sdl.Event), MaxEventCoroutine)
var gLocks = make(chan int, MaxEventCoroutine)

var gStdoutLock = make(chan int)
var gStdoutTimeout = make(chan int)
var gStdoutBuf bytes.Buffer

func init() {
    for i := 0; i < MaxEventCoroutine; i++ {
        gEvents[i] = new(sdl.Event)
        gLocks <- i
    }
}

// PollEvent Returns (*sdl.Event, int, bool)
// the int is the id of the event, CloseEvent should give the id in
// bool stands for whether it's useable or not
func PollEvent() (*sdl.Event, int, bool) {
    i := <- gLocks
    remained := gEvents[i].Poll()
    logger.Printf("Locked %d, type = %d, avaliable = %v\n", i, gEvents[i].Type(), remained)
    return gEvents[i], i, remained
}

func PollEventWait() (*sdl.Event, int) {
    i := <- gLocks
    gEvents[i].WaitPoll()
    return gEvents[i], i
}

func CloseEvent(i int) {
    gLocks <- i
    logger.Printf("Unlocked %d\n", i)
}

func Print(s string) {
    gStdoutBuf.WriteString(s)
    gStdoutLock <- 1
}


func Printf(s string, sth ...interface{}) {
    gStdoutBuf.WriteString(fmt.Sprintf(s, sth...))
    gStdoutLock <- 1
}

func flushOut() {
    go func() {
        time.Sleep(500)
        gStdoutTimeout <- 1
    }()
    select {
        case <-gStdoutLock:
        case <-gStdoutTimeout:
            return
    }
    print(gStdoutBuf.String())
    gStdoutBuf.Reset()
}
