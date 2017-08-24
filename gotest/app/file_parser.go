package app

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadLine wdadawda
func ReadLine(line string) (key string, value string) {
	sliceA := strings.Split(line, "\t")
	return sliceA[3], line
}

// ReadAndParse reads the file and parse each line, returns the map
func ReadAndParse(infile string) (map[string]string, error) {
	f, err := os.Open(infile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	m := make(map[string]string)
	for {
		line, err := buf.ReadString('\n')
		if err == io.EOF {
			break
		}
		if strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimSpace(line)
		k, v := ReadLine(line)
		m[k] = v
	}

	return m, nil
}
