package app

import (
	"fmt"
	"regexp"
	"strings"
)

// TrueGeno get real genotypes from results
func CheckPoint(line string, truegeno string) string {
	s := strings.Split(line, "\t")
	drugtype := make(map[string]string)
	if v, ok := drugtype[drugname]; ok {
		fmt.Println(v)
	} else {
		if drugname != nil {
			return drugname
		}
		drugname := s[0]
		outline := "\\multicolumn{4}{|l|}{\\large\\textbf{s[0]}\\cellcolor{subtabcolor}}\\\\\n\\arrayrulecolor{tabcolor}\\hline\n"
	}
	if strings.ContainsAny(truegeno, "|"){

	}


	}
	r, _ := regexp.Compile("([rs0-9]+) & ([rs0-9]+)")
	if r.MatchString(rs) {
		match := r.FindStringSubmatch(rs)
		if strings.ContainsAny(genotype[match[0]], ";") {
			het1 := strings.Split(genotype[match[0]], ";")
			for _, v1 := range het1 {
				if strings.ContainsAny(genotype[match[1]], ";") {
					het2 := strings.Split(genotype[match[1]], ";")
					for _, v2 := range het2 {
						truegeno = truegeno + v1 + " & " + v2 + "|"
					}
				} else {
					truegeno = truegeno + v1 + " & " + genotype[match[1]] + "|"
				}
			}
		} else if strings.ContainsAny(genotype[match[1]], ";") {
			het2 := strings.Split(genotype[match[1]], ";")
			for _, v2 := range het2 {
				truegeno = truegeno + genotype[match[0]] + "&" + v2 + "|"
			}
		} else {
			truegeno = genotype[match[0]] + genotype[match[1]]
		}
	} else {
		truegeno = genotype[rs]
	}
	return truegeno
}
