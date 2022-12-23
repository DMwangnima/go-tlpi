package user_group

import (
	"bufio"
	"errors"
	"github.com/DMwangnima/go-tlpi/common"
	"io"
	"os"
	"strconv"
)

type Password struct {
	Username string
	Passwd   string
	Uid      uint32
	Gid      uint32
	Gecos    string
	Dir      string
	Shell    string
}

func Getpwnam(name string) (*Password, error) {
    file, err := os.Open("/etc/passwd")
    if err != nil {
    	return nil, err
	}
	defer file.Close()
    return getPwnamFromReader(file, name)
}

func (pw *Password) parseLine(line []byte) ([]byte, error) {
	ents, remain := common.ParseColonLine(line, 6)
	if len(ents) != 7 {
		return remain, errors.New("line is not organized in \"::::::\" form")
	}

	var uid, gid uint32
	uid64, err := strconv.ParseUint(string(ents[2]), 10, 32)
	if err != nil {
		return remain, errors.New("uid is not valid number")
	}
	uid = uint32(uid64)
	gid64, err := strconv.ParseUint(string(ents[3]), 10, 32)
	if err != nil {
		return remain, errors.New("gid id not valid number")
	}
	gid = uint32(gid64)
	pw.Username = string(ents[0])
	pw.Passwd = string(ents[1])
	pw.Uid = uid
	pw.Gid = gid
	pw.Gecos = string(ents[4])
	pw.Dir = string(ents[5])
	pw.Shell = string(ents[6])
	return remain, nil
}

func getPwnamFromReader(r io.Reader, name string) (*Password, error) {
	buf := bufio.NewReader(r)
	for {
		var wholeLine []byte
		for {
			line, isPrefix, err := buf.ReadLine()
			if err != nil {
				if err == io.EOF {
					err = nil
				}
				return nil, err
			}
			if !isPrefix && len(wholeLine) == 0 {
				wholeLine = line
				break
			}
			wholeLine = append(wholeLine, line...)
			if !isPrefix {
				break
			}
		}
        passwd := new(Password)
        _, err := passwd.parseLine(wholeLine)
        if err != nil {
        	continue
		}
		if passwd.Username == name {
			return passwd, nil
		}
	}
}
