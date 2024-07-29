package cmd

import "flag"

type Config struct {
	ArrayLength         int
	ArrayMaxElementSize int
}

func ParseConfig() Config {
	length := flag.Int("length", 100, "Sets the array length")
	size := flag.Int("max-size", 100, "Sets array elements max size")

	flag.Parse()

	return Config{
		ArrayLength:         *length,
		ArrayMaxElementSize: *size,
	}
}
