package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

const origHeader = `<?xml version="1.0" encoding="utf-8"?>`

const origBody = `
<jacartaset jc_model="JaCarta Reader Generic" qty="1">
  <jacarta>
    <jc_usnp date="#DATE#" time="#TIME#">#SERIAL#</jc_usnp>
  </jacarta>
</jacartaset>`

func main() {

	out := origHeader

	for i := 0; i < 100; i++ {
		s := strings.Replace(origBody, "#SERIAL#", randomHex(4), 1)
		s = strings.Replace(s, "#DATE#", randDate().Format("02.01.2006"), 1)
		s = strings.Replace(s, "#TIME#", randDate().Format("15:04:05"), 1)
		out += s
	}

	fmt.Println(out)
	ioutil.WriteFile("readers_3.xml", []byte(out), 0644)
}

func randomHex(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return strings.ToUpper(hex.EncodeToString(bytes))
}

func randDate() time.Time {
	min := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	max := time.Date(2070, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	return time.Unix(sec, 0)
}
