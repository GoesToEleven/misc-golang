package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	f, err := os.Create("all-text.txt")
	if err != nil {
		log.Panic("couldn't create file")
	}
	defer f.Close()

	errOut := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		// error in walking the path
		if err != nil {
			fmt.Println(err)
			return err
		}

		if d.IsDir() {
			return nil
		}

		// _, err = f.WriteString(filepath.Ext(d.Name())+"\n")
		// if err != nil {
		// 	log.Println(err)
		// }

		if filepath.Ext(d.Name()) != ".vtt" {
			return nil
		}

		_, err = f.WriteString(path + "\n" + d.Name() + "\n")
		if err != nil {
			log.Println(err)
		}

		// get content as string from file
		bs, err := os.ReadFile(path)
		if err != nil {
			log.Println(err)
		}

		_, err = f.WriteString(string(bs))
		if err != nil {
			log.Println(err)
		}

		_, err = f.WriteString("\n\n\n\n\n\n")
		if err != nil {
			log.Println(err)
		}

		return nil
	})

	if errOut != nil {
		fmt.Printf("error walking the path: %v\n", errOut)
		return
	}
}
