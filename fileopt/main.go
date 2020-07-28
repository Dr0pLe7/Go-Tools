package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/deckarep/golang-set"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	r = rand.New(rand.NewSource(time.Now().Unix()))
}


func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

var(
	r *rand.Rand
 	f1 string
 	f2 string
 	mode string
	scanner *bufio.Scanner
)



func main(){

	flag.StringVar(&mode, "m", "1+2", "1+2 2+1 1-2 2-1 1=2 2=1 1!=2 2!=1")
	flag.StringVar(&f1, "f1", "", "Read data from a file 1")
	flag.StringVar(&f2, "f2", "", "Read data from a file 2")
	flag.Parse()
	if f1 == "" || f2 == ""{
		flag.Usage()
		os.Exit(0)
	}

	startTime := time.Now() // get current time
	//set 1
	set1 := mapset.NewSet()
	file1, err := os.Open(f1)
	defer file1.Close()
	if err != nil {
		log.Fatalf("Error, cannot open file %v", err)
	}
	scanner = bufio.NewScanner(file1)
	for scanner.Scan() {

		token := strings.TrimSpace(scanner.Text())
		set1.Add(token)
	}

	//set 2
	set2 := mapset.NewSet()
	file2, err := os.Open(f2)
	defer file2.Close()
	if err != nil {
		log.Fatalf("[!] Error, cannot open file %v", err)
	}
	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		token := strings.TrimSpace(scanner.Text())
		set2.Add(token)
	}

	var retval mapset.Set
	switch mode {
	case "1+2":
		retval = set1.Union(set2)
		break
	case "2+1":
		retval = set1.Union(set2)
		break
	case "1-2":
		retval = set1.Difference(set2)
		break
	case "2-1":
		retval = set2.Difference(set1)
		break
	case "1=2":
		retval = set2.Intersect(set1)
		break
	case "2=1":
		retval = set2.Intersect(set1)
		break
	case "1!=2":
		retval = set1.Difference(set2).Union(set2.Difference(set1))
		break
	case "2!=1":
		retval = set1.Difference(set2).Union(set2.Difference(set1))
		break
	default:
		fmt.Println("[!] mode err")
		flag.Usage()
	}
	wfname := RandString(6) + ".txt"
	f, err := os.Create(wfname)
	if err != nil{
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	total := 0
	for i := range retval.Iter(){
		fval := fmt.Sprintf("%s\n",i)
		count, err := w.WriteString(fval)
		if err != nil{
			panic(err)
		}
		total += count
		w.Flush()
	}
	elapsed := time.Since(startTime)

	fmt.Printf("%s write %d bytes\n",wfname,total)
	fmt.Println("App elapsed: ", elapsed)
}