// MarshalIndent 的作用跟 Marshal 类似， 但每个 XML 元素都会出现在一个新的缩进行里面，
// 这个缩进行以 prefix 为前缀， 后跟一个或多个 indent 组成的缩进字符串， indent 出现的数量由行的嵌套深度决定。
package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
}
