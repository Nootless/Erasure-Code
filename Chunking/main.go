package main

import ("fmt"
		"log"
		"io"
 		"os"
)
const chunkSize int64 = 4;

// function that breaks down the files into smaller 4 bit parts
func break_down(f string, chunkSize int64)(){
	// open file
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
		} // if
	defer file.Close()	
	// get file length
	fi,err := file.Stat()
	cSize := fi.Size() / chunkSize
	fileArray := [cSize][chunkSize]byte{}

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
		fmt.Println(string(buf[:readTotal])) 
	}
}

func main() {
	break_down("help.txt", chunkSize)
	
}