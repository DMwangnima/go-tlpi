package user_group

import (
	"reflect"
	"strings"
	"testing"
)

const (
	passwdCase string = "root:x:0:0:root:/root:/bin/bash\ndaemon:x:1:1:daemon:/usr/sbin:/usr/sbin/nologin\nbin:x:2:2:bin:/bin:/usr/sbin/nologin\nsys:x:3:3:sys:/dev:/usr/sbin/nologin\nsync:x:4:65534:sync:/bin:/bin/sync\ngames:x:5:60:games:/usr/games:/usr/sbin/nologin\nman:x:6:12:man:/var/cache/man:/usr/sbin/nologin\nlp:x:7:7:lp:/var/spool/lpd:/usr/sbin/nologin\nmail:x:8:8:mail:/var/mail:/usr/sbin/nologin\nnews:x:9:9:news:/var/spool/news:/usr/sbin/nologin\nuucp:x:10:10:uucp:/var/spool/uucp:/usr/sbin/nologin\nproxy:x:13:13:proxy:/bin:/usr/sbin/nologin\nwww-data:x:33:33:www-data:/var/www:/usr/sbin/nologin\nbackup:x:34:34:backup:/var/backups:/usr/sbin/nologin\nlist:x:38:38:Mailing List Manager:/var/list:/usr/sbin/nologin\nirc:x:39:39:ircd:/run/ircd:/usr/sbin/nologin\ngnats:x:41:41:Gnats Bug-Reporting System (admin):/var/lib/gnats:/usr/sbin/nologin\nnobody:x:65534:65534:nobody:/nonexistent:/usr/sbin/nologin\n_apt:x:100:65534::/nonexistent:/usr/sbin/nologin\nsystemd-network:x:101:102:systemd Network Management,,,:/run/systemd:/usr/sbin/nologin\nsystemd-resolve:x:102:103:systemd Resolver,,,:/run/systemd:/usr/sbin/nologin\nmessagebus:x:103:104::/nonexistent:/usr/sbin/nologin\nsystemd-timesync:x:104:105:systemd Time Synchronization,,,:/run/systemd:/usr/sbin/nologin\npollinate:x:105:1::/var/cache/pollinate:/bin/false\nsshd:x:106:65534::/run/sshd:/usr/sbin/nologin\nsyslog:x:107:113::/home/syslog:/usr/sbin/nologin\nuuidd:x:108:114::/run/uuidd:/usr/sbin/nologin\ntcpdump:x:109:115::/nonexistent:/usr/sbin/nologin\ntss:x:110:116:TPM software stack,,,:/var/lib/tpm:/bin/false\nlandscape:x:111:117::/var/lib/landscape:/usr/sbin/nologin\nusbmux:x:112:46:usbmux daemon,,,:/var/lib/usbmux:/usr/sbin/nologin\nlxd:x:999:100::/var/snap/lxd/common/lxd:/bin/false\nntp:x:113:118::/nonexistent:/usr/sbin/nologin\n_chrony:x:114:124:Chrony daemon,,,:/var/lib/chrony:/usr/sbin/nologin\nfwupd-refresh:x:115:125:fwupd-refresh user,,,:/run/systemd:/usr/sbin/nologin"
)

func TestGetpwnamFromReader(t *testing.T) {
	rd := strings.NewReader(passwdCase)
	expectPasswd := &Password{
		Username: "root",
		Passwd:   "x",
		Uid:      0,
		Gid:      0,
		Gecos:    "root",
		Dir:      "/root",
		Shell:    "/bin/bash",
	}
	res, err := getPwnamFromReader(rd, "root")
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(res, expectPasswd) {
		t.Errorf("results are not equal")
	}
}
