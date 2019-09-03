package main

import (
	"crypto/rand"
	"encoding/hex"
	"io/ioutil"
	"strings"
)

const origHeader = `<?xml version="1.0" encoding="utf-8"?>`

const origBody = `
<jacartaset jc_model="JaCarta Reader Generic" qty="1">
  <jacarta>
    <jc_usnp>kirk</jc_usnp>
  </jacarta>
</jacartaset>`

func main() {

	out := origHeader

	for i := 0; i < 1000; i++ {
		out += strings.Replace(origBody, "kirk", randomHex(4), 1)
	}

	//fmt.Println(out)
	ioutil.WriteFile("readers.xml", []byte(out), 0644)
}

func randomHex(n int) string {
	bytes := make([]byte, n)
	rand.Read(bytes)
	return strings.ToUpper(hex.EncodeToString(bytes))
}
