package filehandling

import (
	"fmt"
	"log"
	"os"

	"google.golang.org/protobuf/proto"
)

func WriteToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err = os.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}

func ReadFromFile(fname string, pb proto.Message) error {
	in, err := os.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can't read from file", err)
		return err
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Can't deserialise from file", err)
		return err
	}

	return nil
}
