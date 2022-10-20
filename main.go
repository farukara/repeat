package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
    // app := os.Args[0]
    taskno := os.Args[1]
    reptime := os.Args[2]

    // get current time
    currentTime := time.Now()
    timeStampString := currentTime.Format("2006-01-02 15:04:05")    
    layOut := "2006-01-02 15:04:05"     
    timeStamp, err := time.Parse(layOut, timeStampString)
    if err != nil {
        fmt.Println(err)          
    }   
    year, month, day := currentTime.Date()
    hr, min, sec := timeStamp.Clock()
    fmt.Println(year, month, day, hr, min, sec)


    // is task overdue
    cmd := exec.Command("task", "_get", taskno + ".tags.OVERDUE")
    output, err := cmd.Output()
    if err != nil {
        fmt.Println(err)          
    }
    isoverdue := false
    if len(output) > 1 {isoverdue = true }
    fmt.Println("isoverdue:", isoverdue)

    // get tasks due date
    cmd = exec.Command("task", "_get", taskno + ".due")
    output, err = cmd.Output()
    if err != nil {
        fmt.Println(err)          
    }   
    duedate,err := time.Parse("2006-01-02T15:04:05", strings.TrimSpace(string((output))))
    if err != nil {
        fmt.Println(err)          
    }   
    fmt.Println("duedate:", duedate)

    var hoursToSix string
    switch {
        case currentTime.Hour() < 6:
            hoursToSix = "+" + strconv.Itoa(6-currentTime.Hour())
        case currentTime.Hour() == 6:
            hoursToSix = ""
        case currentTime.Hour() > 6:
            hoursToSix = "-" + strconv.Itoa(currentTime.Hour()-6)
    } 
    fmt.Println("hoursToSix:", hoursToSix)

    Reptimes := map[string][2]string{
        "1d": {"1d", "wait:due-4h"},
        "2d": {"2d", "wait:due-4h"},
        "3d": {"3d", "wait:due-4h"},
        "5d": {"5d", "wait:due-1d"},
        "1w": {"1w", "wait:due-1d"},
        "2w": {"2w", "wait:due-3d"},
        "3w": {"3w", "wait:due-3d"},
        "1m": {"1m", "wait:due-3d"},
        "2m": {"2m", "wait:due-1w"},
        "3m": {"3m", "wait:due-1w"},
        "4m": {"4m", "wait:due-1w"},
        "5m": {"5m", "wait:due-1w"},
        "6m": {"6m", "wait:due-1w"},
        "7m": {"7m", "wait:due-2w"},
        "8m": {"8m", "wait:due-2w"},
        "9m": {"9m", "wait:due-2w"},
        "10m": {"10m", "wait:due-2w"},
        "11m": {"11m", "wait:due-2w"},
        "1y": {"1y", "wait:due-2w"},
        "2y": {"2y", "wait:due-3w"},
        "3y": {"3y", "wait:due-3w"},
        "4y": {"4y", "wait:due-1m"},
        "5y": {"5y", "wait:due-1m"},
    }
    
    switch isoverdue {
    // task 12 mod due:1d-4h wait:due-4h
        case true:
            args := []string{taskno, "mod", "due:" + Reptimes[reptime][0] + hoursToSix + "h",Reptimes[reptime][1]}
            fmt.Println("command:", "task", args)
            cmd := exec.Command("task", args...)
            output, err = cmd.Output()
            if err != nil {
                fmt.Println(err)          
            }   
            fmt.Println("output:", string(output))
        case false:
            days := int(duedate.Sub(currentTime).Hours()/24)+1
            fmt.Println("days to due:", days)
            args := []string{taskno, "mod", "due:" + Reptimes[reptime][0] + "+" + strconv.Itoa(days) + "d" + hoursToSix + "h",Reptimes[reptime][1]}
            fmt.Println("command:", "task", args)
            cmd := exec.Command("task", args...)
            output, err = cmd.Output()
            if err != nil {
                fmt.Println(err)          
            }   
            fmt.Println("output:", string(output))
        default:
    }
}
