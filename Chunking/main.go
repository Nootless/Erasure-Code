package main

import ("fmt"
		"log"
		"io"
 		"os"
)
const chunkSize int = 4;

// function that breaks down the files into smaller 4 bit parts
func break_down(f string, chunkSize int)([][]byte){
	// open file
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
		} // if
	defer file.Close()	
	// get file length
	fileArray := [][]byte{}
	fi,err := file.Stat()
	cSize := int(fi.Size() / int64(chunkSize))
	fmt.Println(cSize)

	// Make array of length filesize / chunksize

	buf := make([]byte,chunkSize)

	for {
        // read content to buffer
        readTotal, err := file.Read(buf)
        if err != nil {
            if err != io.EOF {
                fmt.Println(err)
            }
            break
        }
		fileArray = append(fileArray, buf[:readTotal])
	}
	fmt.Println(fileArray)
	return fileArray
}

func main() {
	break_down("help.txt", chunkSize)
	
}