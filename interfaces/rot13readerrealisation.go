package interfaces


import (
	"fmt"
	"io"
	"os"
	"strings"
)
//var rot13 = map[byte]byte{
//
//	"Bell Labs": Vertexx{
//		40.68433, -74.39967,
//	},
//	"Google": Vertexx{
//		37.42202, -122.08408,
//	},
//}
type NotEqualsLengthStrings struct {
	lengthFirst  int
	lengthSecond int
}

func (ne NotEqualsLengthStrings) Error() string {
	return fmt.Sprintf("Not equal length strings: %d, %d", ne.lengthFirst, ne.lengthSecond)
}

type rot13Reader struct {
	r io.Reader
	mapwsh map[byte]byte
}
func (rot rot13Reader) Read(b []byte) (int, error){

	n, err := rot.r.Read(b)
	if err == nil {
		for i := 0; i < n; i++ {
			_, ok := rot.mapwsh[b[i]]
			if ok {
				b[i] = rot.mapwsh[b[i]]
			}
		}
		return len(b), nil
	}
	return 0, nil
}


func createEncodedMap(decoded []byte, encoded []byte) (map[byte]byte, error){
	if len(decoded) != len(encoded) {
		return nil, NotEqualsLengthStrings{
			len(decoded),
			len(encoded),
		}
	}
	m := make(map[byte]byte)
	for i, _ := range decoded {
		m[encoded[i]] = decoded[i]
	}
	return m, nil
}

func testRot13() {
	s1 := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	s2 := []byte("NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm")
	bytemap , error := createEncodedMap(s1,s2)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(bytemap, error)
		s := strings.NewReader("Lbh penpxrq gur pbqr!")
		r := rot13Reader{
					s,
					bytemap,
				}
		io.Copy(os.Stdout, &r)
	}

}

