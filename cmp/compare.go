package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadAndParse reads the file and parse each line, returns the map
func ReadAndParse(infile string) map[string]string {
	var totcount int = 0
	f, err := os.Open(infile)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	m := make(map[string]string)
	var repeatlist []string
	for {
		var s_new []string
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.Replace(line, "\t\t\t\t\t\t", "\t-\t-\t-\t-\t-\t-", 1)
		if strings.HasPrefix(line, "#") {
			continue
		}
		s := strings.Split(line, "\t")
		id := s[0]
		for _, v := range s {
			if v == "" {
				v = "-"
			}
			s_new = append(s_new, v)
		}
		line = strings.Join(s_new, "\t")
		//	line = strings.TrimSpace(line)
		_, ok := m[id]
		if ok {
			totcount++
			repeatlist = append(repeatlist, id)
			//	fmt.Printf("\n%s is repeated in %s\n", id, infile)
		} else {
			m[id] = line
		}
	}
	fmt.Printf("repeated in %s: #%d\n%s\n", infile, totcount, repeatlist)
	return m
}

func main() {
	argsWithProg := os.Args
	infile1 := argsWithProg[1]
	infile2 := argsWithProg[2]

	var (
		scanid map[string]string
		listid map[string]string
	)
	scanid = ReadAndParse(infile1)
	listid = ReadAndParse(infile2)

	// find scaned samples in list
	file1, error := os.Create("./InScanNotInList.txt")
	if error != nil {
		fmt.Println(error)
	}
	file2, error := os.Create("./InListNotInScan.txt")
	if error != nil {
		fmt.Println(error)
	}
	file3, error := os.Create("./InListInScan.txt")
	if error != nil {
		fmt.Println(error)
	}
	file4, error := os.Create("./UniqueInList.txt")
	if error != nil {
		fmt.Println(error)
	}

	count1 := 0
	for k := range scanid {
		_, ok := listid[k]
		count1 = count1 + 1
		if !ok {
			writestring := k
			file1.WriteString(writestring)
		}
	}
	file1.Close()

	// find scaned samples not in list

	count2 := 0
	count3 := 0
	for k, val := range listid {
		_, ok := scanid[k]
		file4.WriteString(val)
		if ok {
			count2 = count2 + 1
			file3.WriteString(listid[k])
		} else {
			count3 = count3 + 1
			writestring := listid[k]
			file2.WriteString(writestring)
		}
	}
	file2.Close()
	file3.Close()
	file4.Close()
	fmt.Printf("scaned in list: %d\n", count2)
	fmt.Printf("total in scaned: %d\n", count1)
	fmt.Printf("total in list: %d\n", count2+count3)
}
