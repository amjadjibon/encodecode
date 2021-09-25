package main

import (
	"fmt"
	"github.com/amjadjibon/encodecode/flatbuff/models"
	flatbuffers "github.com/google/flatbuffers/go"
	"reflect"
)

func MakeUser(b *flatbuffers.Builder, name []byte, age int32) []byte {
	// re-use the already-allocated Builder:
	b.Reset()

	// create the name object and get its offset:
	namePosition := b.CreateByteString(name)

	// write the User object:
	models.PersonStart(b)
	models.PersonAddName(b, namePosition)
	models.PersonAddAge(b, age)
	userPosition := models.PersonEnd(b)

	// finish the write operations by our User the root object:
	b.Finish(userPosition)

	// return the byte slice containing encoded data:
	return b.Bytes[b.Head():]
}

func ReadUser(buf []byte) (name []byte, age int32) {
	// initialize a User reader from the given buffer:
	user := models.GetRootAsPerson(buf, 0)

	// point the name variable to the bytes containing the encoded name:
	name = user.Name()

	// copy the user's id (since this is just a uint64):
	age = user.Age()

	return
}

func main()  {
	b := flatbuffers.NewBuilder(0)
	buf := MakeUser(b, []byte("Arthur Dent"), 42)
	name, id := ReadUser(buf)
	fmt.Println(reflect.TypeOf(buf))

	fmt.Printf("%s has id %d. The encoded data is %d bytes long.\n", name, id, len(buf))
}
