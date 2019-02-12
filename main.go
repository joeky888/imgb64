package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		usage()
	}

	for _, filename := range os.Args[1:] {

		f, err := os.Open(filename)
		if err != nil {
			print(filename + ": " + err.Error())
			continue
		}

		reader := bufio.NewReader(f)
		content, _ := ioutil.ReadAll(reader)

		encoded := base64.StdEncoding.EncodeToString(content)

		fmt.Printf("<img src=\"data:image/%s;base64,%s\"/>\n", getFiletype(content), encoded)
	}
}

func usage() {
	println("Usage: imgb64 FILE_NAME")
	os.Exit(0)
}

func getFiletype(contentByte []byte) string {

	lenb := len(contentByte)
	/*---------------Read file end------------------------*/

	if lenb > 28 && HasPrefix(contentByte, "\x89PNG\x0d\x0a\x1a\x0a") {
		// print("PNG image data")
		return "png"
	} else if lenb > 16 &&
		(HasPrefix(contentByte, "GIF87a") || HasPrefix(contentByte, "GIF89a")) {
		// print("GIF image data")
		return "gif"
	} else if lenb > 32 && HasPrefix(contentByte, "\xff\xd8") {
		// print("JPEG image data")
		return "jpeg"
	}
	return "unknown"
}

func HasPrefix(s []byte, prefix string) bool {
	return len(s) >= len(prefix) && Equal(s[:len(prefix)], prefix)
}

func Equal(a []byte, b string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range []byte(b) {
		if v != a[i] {
			return false
		}
	}
	return true
}
