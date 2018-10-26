package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"
)

func main () {
	b, err := ioutil.ReadFile("ip.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}
	ipInGit := string(b)

	for {
		resp, err := http.Get("https://api.ipify.org/")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if string(body) != ipInGit {
			ioutil.WriteFile("ip.txt", body, 0666)
			exec.Command("git","add","ip.txt").Run()
			exec.Command("git","commit", "-m", "updateip").Run()
			exec.Command("git","push").Run()
		}
		time.Sleep(5*60*time.Second)
	}

}
