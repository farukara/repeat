package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var Reptimes map[string][2]string = map[string][2]string{
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

func main() {

    // Error handling for bad inputs
    // only one arg or more than 2 arguments
    args := os.Args
    if len(args) != 3 {
        fmt.Println("\033[7;31merror: wrong number of arguments: Repeat requires Task No and time arguments\n", args)
        return
    }
    // task no regex check
    tasknorgx := regexp.MustCompile(`\d*`)
    tasknorgxindex := tasknorgx.FindStringIndex(args[1])
    if tasknorgxindex == nil {
        fmt.Println("\033[7;31merror: task no is not recognized:\n", args[1])
    }
    // check if duration exists in the map
    if _,ok := Reptimes[args[2]]; !ok {
        fmt.Println("\033[7;31merror: time argument is not correct:\n", args[2])
        fmt.Println("\033[7;0mshould be one of:")
        for k := range Reptimes {
            fmt.Print(k, " ")
        }
        return
    }

    // app := os.Args[0]
    taskno := args[1]
    reptime := args[2]

    // get current time
    currentTime := time.Now()

    // is task overdue
    cmd := exec.Command("task", "_get", taskno + ".tags.OVERDUE")
    output, err := cmd.Output()
    if err != nil {
        fmt.Println("\033[7;31merror:", err)          
    }
    isoverdue := false
    if len(output) > 1 {isoverdue = true }

    // get tasks due date
    cmd = exec.Command("task", "_get", taskno + ".due")
    output, err = cmd.Output()
    if err != nil {
        fmt.Println("\033[7;31merror:", err)          
    }   
    duedate,err := time.Parse("2006-01-02T15:04:05", strings.TrimSpace(string((output))))
    if err != nil {
        fmt.Println("\033[7;31merror:", err)          
    }   
    // check if due date more than 1 month to avoid accidental edits
    if (duedate.Sub(time.Now())) > time.Hour * 720 {
        fmt.Println("\033[7;33mWARNING: you are editing a task that is due more than 1 month later")
        s := bufio.NewScanner(os.Stdin)
        fmt.Print("Do you want to continue (y/n): ")
        s.Scan()
        if s.Text() == "n" || s.Text() == "N"  {
            fmt.Println("\033[7;33mTask WAS NOT edited.")
            return
        } else if !(s.Text() == "y" || s.Text() == "Y") {
            fmt.Println("\033[7;31mUnknown input: Task WAS NOT edited.")
            return
        }
    }

    var hoursToSix string
    switch {
        case currentTime.Hour() < 6:
            hoursToSix = "+" + strconv.Itoa(6-currentTime.Hour()) + "h"
        case currentTime.Hour() == 6:
            hoursToSix = ""
        case currentTime.Hour() > 6:
            hoursToSix = "-" + strconv.Itoa(currentTime.Hour()-6) + "h"
    } 

    
    switch isoverdue {
    // task 12 mod due:1d-4h wait:due-4h
        case true:
            args := []string{taskno, "mod", "due:" + Reptimes[reptime][0] + hoursToSix, Reptimes[reptime][1]}
            fmt.Println("task", args)
            cmd := exec.Command("task", args...)
            output, err = cmd.Output()
            if err != nil {
                fmt.Println("\033[7;31merror:", err)          
            }   
            fmt.Println(string(output))
        case false:
            days := int(duedate.Sub(currentTime).Hours()/24)+1
            args := []string{taskno, "mod", "due:" + Reptimes[reptime][0] + "+" + strconv.Itoa(days) + "d" + hoursToSix,Reptimes[reptime][1]}
            fmt.Println("task", args)
            cmd := exec.Command("task", args...)
            output, err = cmd.Output()
            if err != nil {
                fmt.Println("\033[7;31merror:", err)          
            }   
            fmt.Println(string(output))
        default:
    }
}
