package Socket

import (
	"net/http"
	"io"
	"os"
)

func main() {
	resp, err := http.Get("https://www.baidu.com/")
	if err != nil {
		return
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
