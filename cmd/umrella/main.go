package main

import (
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/k0kubun/pp"
)

type Config struct {
	Foo string `toml:"foo"`
}

func main() {
	f, err := os.Open("foo.toml")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	var cfg Config
	if _, err := toml.Decode(string(b), &cfg); err != nil {
		panic(err)
	}
	pp.Println(cfg)
}
