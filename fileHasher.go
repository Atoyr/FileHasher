package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "FileHasher"
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			return nil
		}
		for _, file := range c.Args() {
			h, err := hash(file)
			if err != nil {
				log.Fatal(err)
			}
			size, err := getBytes(file)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s : %d bytes : %s \n", file, size, h)
		}
		return nil
	}

	app.Run(os.Args)
}

func hash(file string) (string, error) {
	hasher := sha1.New()
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer f.Close()
	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil)), nil
}

func getBytes(file string) (int64, error) {
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	info, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
