package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"
	"reflect"
	"time"
	//	"github.com/stacktitan/ldapauth"
	//	"os"
	"net/http"
	"os/exec"
	"strings"
)

type error interface {
	Error() string
}

func throwError() error {
	return errors.New("Error")
}

func foo() error {
	return errors.New("Some Error Occurred")
}

func read(s []byte) {

	fmt.Println(string(s))
	fmt.Println("\n")

}

func f() {
	fmt.Println("f function")
}

func strlen(s string, c chan string) {
	result := s + "AppleBois"
	c <- result
}

func scan(s string, c chan string) {
	_, err := net.Dial("tcp", s)
	if err != nil {
		result := "Error on :" + s
		c <- result
	}
	c <- s
}

func main() {
	//	os.Setenv("HTTP_PROXY", "http://127.0.0.1:8080")

	fmt.Println("Hello, Black Hat Gophers!")
	var x string
	fmt.Scanln(&x)
	switch x {
	case "foo":
		fmt.Println("Found foo")
	case "bar":
		fmt.Println("Found bar")
	default:
		fmt.Println("Default case")
	}
	fmt.Println(reflect.TypeOf(x))

	var i int
	for i = 0; i < 10; i++ {
		fmt.Println(i)
	}

	nums := []int{2, 4, 6, 8}
	for _, test := range nums {
		fmt.Println(test)
	}

	go f()
	time.Sleep(1 * time.Second)

	c := make(chan string)
	go strlen("Salutations", c)
	go strlen("World", c)
	x, y := <-c, <-c
	fmt.Println(x, y)

	err := throwError()
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}

	a, err := net.Dial("tcp", "scanme.nmap.org:80")
	//	a, err = net.Dial("tcp",  "applebois.com:80") // No such host
	fmt.Println(err)
	if err == nil {
		fmt.Println("Connection successful") // Connection Successful, tcp on port 80 is good
	}
	fmt.Println(a)
	/*
		for i := 1; i <= 1024; i++ {
			address := fmt.Sprintf("scanme.nmap.org:%d", i)
			c = make(chan string)
			go scan(address, c)
			z := <-c
			fmt.Println(z)

		}

		r1, err := http.Get("http://www.google.com/robots.txt")
		data, err := ioutil.ReadAll(r1.Body)
		if err != nil {
			fmt.Println("Error")
		}
		fmt.Println(r1.Status)
		read(data)
		defer r1.Body.Close()

		r2, err := http.Head("http://www.google.com/robots.txt")
		data, err = ioutil.ReadAll(r2.Body)
		if err != nil {
			fmt.Println("Error")
		}
		fmt.Println(r2.Status)
		read(data)
		fmt.Println("\n")
		defer r2.Body.Close()
	*/

	form := url.Values{}
	form.Add("foo", "bar")
	r3, err := http.Post("http://blog.orange.tw/", "application/x-www-form-urlencoded", strings.NewReader(form.Encode()))
	data, err := ioutil.ReadAll(r3.Body)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(r3.Status)
	read(data)
	fmt.Println("\n")
	defer r3.Body.Close()

	request, err := http.NewRequest("GET", "https://tptgtwa.com/keycheck.php", nil)
	request.Header.Set("Apikey", "90403c2632ea")
	checkError(err)

	var client http.Client
	resp, err := client.Do(request)
	data, err = ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	var check int
	check = (strings.Index(string(data), "TOOL SUCCESSFULLY SETUP!"))

	if check == -1 {
		fmt.Println("API Header not configured correctly")
	} else {
		fmt.Println("Setup completed")
	}

	cmd := exec.Command("id")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// Print the output
	fmt.Println(string(stdout))

}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
