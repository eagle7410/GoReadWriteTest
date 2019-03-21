package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func BenchmarkWriteFile (b *testing.B) {
	for i := 0; i < b.N; i++ {

		var s []byte

		for k:=0; k<10000000 ; k++  {
			s = append(s, []byte("w")...)
		}

		if err := ioutil.WriteFile("test", s, os.ModePerm); err != nil {
			log.Fatal(err)
		}

	}
}
func BenchmarkReadFile (b *testing.B) {
	for i := 0; i < b.N; i++ {

		if _, err := ioutil.ReadFile("test"); err != nil {
			log.Fatal(err)
		}

	}
}
