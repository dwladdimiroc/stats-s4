package main

import (
	"fmt"
	"github.com/dwladdimiroc/stats-s4/exec"
	"github.com/dwladdimiroc/stats-s4/write_file"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	duration, errParser := strconv.Atoi(os.Args[1])
	if errParser != nil {
		fmt.Println("Error... No es un número válido el ingresado como parámetro...")
		fmt.Println(errParser)
		os.Exit(2)
	}

	var wg sync.WaitGroup

	// 1) Execute Zk
	appCmdZk := "sh"
	argsCmdZk := []string{"zk.sh"}
	dirCmdZk := "apache-s4-0.6.0"
	var outputZk string

	wg.Add(1)
	go func(appCmd string, argsCmd []string, dirCmd string, output *string, duration int) {
		defer wg.Done()
		*output = exec.Start(appCmd, argsCmd, dirCmd, duration+20)
		fmt.Println("Terminado " + argsCmd[0])
	}(appCmdZk, argsCmdZk, dirCmdZk, &outputZk, duration)

	// 2) Execute Stats
	appCmdStats := "sh"
	argsCmdStats := []string{"statistics.sh"}
	dirCmdStats := "statisticsHW"
	var outputStats string

	wg.Add(1)
	go func(appCmd string, argsCmd []string, dirCmd string, output *string, duration int) {
		defer wg.Done()
		*output = exec.Start(appCmd, argsCmd, dirCmd, duration+20)
		fmt.Println("Terminado " + argsCmd[0])
	}(appCmdStats, argsCmdStats, dirCmdStats, &outputStats, duration)

	time.Sleep(2 * time.Second)

	// 3) Execute Node
	appCmdNode := "sh"
	argsCmdNode := []string{"node.sh"}
	dirCmdNode := "experimentJournal"
	var outputNode string

	wg.Add(1)
	go func(appCmd string, argsCmd []string, dirCmd string, output *string, duration int) {
		defer wg.Done()
		*output = exec.Start(appCmd, argsCmd, dirCmd, duration+15)
		fmt.Println("Terminado " + argsCmd[0])
	}(appCmdNode, argsCmdNode, dirCmdNode, &outputNode, duration)

	time.Sleep(3 * time.Second)

	// 4) Execute Compile
	appCmdCompile := "sh"
	argsCmdCompile := []string{"compile.sh"}
	dirCmdCompile := "experimentJournal"
	outputCompile := exec.Execute(appCmdCompile, argsCmdCompile, dirCmdCompile)
	fmt.Println("Terminado " + argsCmdCompile[0])
	files.WriteOutput("compile", outputCompile)

	// 5) Execute Adapter
	appCmdAdapter := "sh"
	argsCmdAdapter := []string{"adapterDynamic.sh"}
	dirCmdAdapter := "experimentJournal"
	var outputAdapter string

	wg.Add(1)
	go func(appCmd string, argsCmd []string, dirCmd string, output *string, duration int) {
		defer wg.Done()
		*output = exec.Start(appCmd, argsCmd, dirCmd, duration)
		fmt.Println("Terminado " + argsCmd[0])
	}(appCmdAdapter, argsCmdAdapter, dirCmdAdapter, &outputAdapter, duration)

	wg.Wait()

	files.WriteOutput("zk", outputZk)
	files.WriteOutput("node", outputNode)
	files.WriteOutput("adapter", outputAdapter)
	files.WriteOutput("stats", outputStats)
}
