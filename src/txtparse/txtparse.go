// txt_parse.go
package txtparse

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type TxtParse struct {
	filename string
	sea      []SSeason
}

type SSeason struct {
	season string
	e      []string
	name   []string
}

func (txt *TxtParse) Init(fn string) bool {
	if fn != "" {
		txt.filename = fn
		txt.sea = make([]SSeason, 2, 10)
		for _, itor := range txt.sea {
			itor.name = make([]string, 5, 1024)
			itor.e = make([]string, 5, 1024)
		}
		return true
	} else {
		return false
	}

}
func (txt *TxtParse) GetSeason(name string) (string, string, error) {
	if "" == name {
		return "", "", errors.New("no name")
	}
	for _, itor := range txt.sea {
		for idx, itor2 := range itor.name {
			if itor2 == name {
				return itor.season, itor.e[idx], nil
			}
		}
	}
	return "", "", errors.New("no match")
}
func (txt *TxtParse) Print() {
	for _, itor := range txt.sea {
		fmt.Println("season:", itor.season)
		for idx, itor2 := range itor.name {
			fmt.Println("name:", itor2)
			fmt.Println("e: ", itor.e[idx])
		}
	}
}
func (txt *TxtParse) Parse() (bool, error) {

	f, ok := os.Open(txt.filename)
	if ok != nil {
		return false, errors.New("no that file")
	}
	
	defer f.Close()
	r := bufio.NewReader(f)

	reg1 := regexp.MustCompile(`(Season).(\d*).`)
	reg2 := regexp.MustCompile(`(\d*).(.*)[.]`)
	var (
		current *SSeason
		seacnt  int
		namecnt int
	)
	for {
		line, ok := r.ReadString('\n')
		if ok != nil {
			break
		}

		submatch1 := reg1.FindAllStringSubmatch(line, -1)
		if submatch1 != nil {
			namecnt = 0
			if seacnt >= len(txt.sea) {
				var newone SSeason
				newone.name = make([]string, 5, 1024)
				newone.e = make([]string, 5, 1024)
				txt.sea = append(txt.sea, newone)
			}
			for idx, itor := range txt.sea {
				if itor.season == "" {
					txt.sea[idx].season = submatch1[0][2]
					current = &txt.sea[idx]
					namecnt = 0
					seacnt = seacnt + 1
					break
				}
			}

		} else if submatch2 := reg2.FindAllStringSubmatch(line, -1); submatch2 != nil {
			if namecnt >= len(current.name) {
				var newone1 string
				var newone2 string
				current.name = append(current.name, newone1)
				current.e = append(current.e, newone2)
			}
			for idx2, itor2 := range current.name {
				if itor2 == "" {
					current.e[idx2] = submatch2[0][1]
					current.name[idx2] = strings.ToLower(submatch2[0][2])
					namecnt = namecnt + 1
					break
				}
			}

		} else {
			fmt.Printf("line:%s is not match\r\n", line)
		}
	}
	return true, nil
}
