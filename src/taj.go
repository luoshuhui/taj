package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"txtparse"
)

func walkfn(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}

	s := f.Name()
	/*猫和老鼠.tom and jerry.S01E02.小狗野餐.Pup On A Picnic.mkv
	----->0	-	猫和老鼠.tom and jerry
	----->1	-	01
	----->2	-	02
	----->3	-	小狗野餐
	----->4	-	Pup On A Picnic
	*/
	reg := regexp.MustCompile(`(.+)[.]S(\d*)E(\d*)[.](.+)[.](.*).mkv`)
	submatch := reg.FindAllStringSubmatch(s, -1)
	if submatch != nil {
		sea, e, err := txt.GetSeason(strings.ToLower(submatch[0][5]))

		if err == nil {
			//fmt.Println("find,sea: ", sea)
			//fmt.Println("find,e: ", e)
			newname := submatch[0][1] + ".S" + sea + "E" + e + "." + submatch[0][4] + "." + submatch[0][5] + ".mkv"

			er := os.Rename(s, newname)
			if er != nil {
				fmt.Printf("rename error:\r\n%s\r\n->%s\r\n", s, newname)
			} else {
				fmt.Printf("rename done:\r\n%s\r\n->%s\r\n", s, newname)
			}

		} else {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Printf("file:%s not match!\r\n", s)
	}
	return nil
}

var txt txtparse.TxtParse

func main() {

	txt.Init("t.txt")
	txt.Parse()
	//txt.Print()
	filepath.Walk(".\\", walkfn)
	return
}
