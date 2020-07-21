package main 
import (
	"fmt"
"encoding/base64"

)

type SID struct {
	RevisionLevel     int
	SubAuthorityCount int
	Authority         int
	SubAuthorities    []int
	RelativeID        *int
}

func (sid SID) String() string {
	s := fmt.Sprintf("S-%d-%d", sid.RevisionLevel, sid.Authority)
	for _, v := range sid.SubAuthorities {
		s += fmt.Sprintf("-%d", v)
	}
	return s
}

func (sid SID) RID() int {
	l := len(sid.SubAuthorities)
	return sid.SubAuthorities[l-1]
}

func Decode(b []byte) SID {

	var sid SID

	sid.RevisionLevel = int(b[0])
	sid.SubAuthorityCount = int(b[1]) & 0xFF

	for i := 2; i <= 7; i++ {
		sid.Authority = sid.Authority | int(b[i])<<uint(8*(5-(i-2))) //fix (shift count type int, must be unsigned integer)
	}

	var offset = 8
	var size = 4
	for i := 0; i < sid.SubAuthorityCount; i++ {
		var subAuthority int
		for k := 0; k < size; k++ {
			subAuthority = subAuthority | (int(b[offset+k])&0xFF)<<uint(8*k) //fix (shift count type int, must be unsigned integer)
		}
		sid.SubAuthorities = append(sid.SubAuthorities, subAuthority)
		offset += size
	}

	return sid
}

func main() {
	// A objectSID in base64
	// This is just here for the purpose of an example program
	b64sid := `AQUAAAAAAAUVAAAAMhzjG7Vz0nTGP9dbUw0AAA==`

	// Convert the above into binary form. This is the value you would
	// get from a LDAP query on AD.
	bsid, _ := base64.StdEncoding.DecodeString(b64sid)

	// Decode the binary objectsid into a SID object
	sid := Decode(bsid)

	// Print out just one component of the SID
	fmt.Println("Authority:" , sid.Authority)

	// Print out the relative identifier
	fmt.Println("RID:" ,sid.RID())

	// Print the entire ObjectSID
	fmt.Println("SID:" , sid)
}

