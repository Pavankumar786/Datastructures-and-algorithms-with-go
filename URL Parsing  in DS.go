//1.Encoding parsing
package main
import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	queryStr := "name=Pavankumar%20keerthi&phone=%2B919703043048&phone=%2B919866392604"
	params, err := url.ParseQuery(queryStr)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Query Params: ")
	for key, value := range params {
		fmt.Printf("  %v = %v\n", key, value)
	}
}
$ro run main.go
Query Params: 
  name = [pavankumar keerthi]
  phone = [+919703043048 +919866392604]

//2.Decoding parsing
package main
import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	encodedValue := "Hell%C3%B6+W%C3%B6rld%40Golang"
	decodedValue, err := url.QueryUnescape(encodedValue)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(decodedValue)
}
$ro run main.go
Hellö Wörld@Golang

//3.path segment
package main
import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	path := "path%20with%3Freserved+characters"
	unescapedPath, err := url.PathUnescape(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(unescapedPath)
}
$go run main.go
path with?reserved+characters

//4.complete url parsing
package main
import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	u, err := url.Parse("https://www.website.com/person?name=Pavankumar%20keerthi&phone=%2B919999999999&phone=%2B918888888888")

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Scheme: ", u.Scheme)
	fmt.Println("Host: ", u.Host)

	queries := u.Query()
	fmt.Println("Query Strings: ")
	for key, value := range queries {
		fmt.Printf("  %v = %v\n", key, value)
	}
	fmt.Println("Path: ", u.Path)
}
$ro run main.go
Scheme:  https
Host:  www.website.com
Query Strings: 
  name = [Pavankumar keerthi]
  phone = [+919999999999 +918888888888]
Path:  /person