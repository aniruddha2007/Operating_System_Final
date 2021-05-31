package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println("選一個功能\n")
		var pgm float64
		var sharedmem, Mthreading, exit = "0]Shared-Memory", "1]Multi-Threading", "2] Exit"
		fmt.Println(sharedmem)
		fmt.Println(Mthreading)
		fmt.Println(exit)
		fmt.Scanln(&pgm)
		fmt.Print("Your Current Selection is: ", pgm)

		//program assign
		// sharedmem Command
		if pgm == 0 {
			fmt.Println(sharedmem, "starting")
			time.Sleep(1 * time.Second)
			cmd := exec.Command("./shm")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout

			if err := cmd.Run(); err != nil {
				fmt.Println("Error:", err)
			}
		}
		//Mthreadingsion Command
		if pgm == 1 {
			fmt.Println(Mthreading, "starting")
			time.Sleep(1 * time.Second)
			cmd := exec.Command("./threading")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout

			if err := cmd.Run(); err != nil {
				fmt.Println("Error:", err)
			}
		}
		//Exit Command
		if pgm == 3 {
			fmt.Println(exit, "starting")
			time.Sleep(1 * time.Second)
			cmd := exec.Command("exit")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stdout

			if err := cmd.Run(); err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}
