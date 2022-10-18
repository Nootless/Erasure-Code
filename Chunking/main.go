package main

import ("fmt"
		"log"
		"io"
 		"os"
		"errors"
		"encoding/gob"
		"bytes"
)
const chunkSize int = 4;

type Node struct {
	info interface{}
	next *Node
}

type List struct {
	head *Node
}

func (l *List) Insert(d interface{}) {
	list := &Node{info: d, next: nil}
	if l.head == nil {
		l.head = list
	} else {
		p := l.head
		for p.next != nil {
			p = p.next
		}
		p.next = list
	}
}

func Show(l *List) {
	p := l.head
	for p != nil {
		fmt.Printf("-> %v ", p.info)
		p = p.next
	}
}
func GetBytes(key interface{}) ([]byte, error) {
    var buf bytes.Buffer
    enc := gob.NewEncoder(&buf)
    err := enc.Encode(key)
    if err != nil {
        return nil, err
    }
    return buf.Bytes(), nil
}

// function that breaks down the files into smaller 4 bit parts
func chunk(f string, chunkSize int)(List){
	// open file
	
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
		} // if
	defer file.Close()	
	// get file length

	buf := make([]byte,chunkSize)
	// var chunkList *List;
	var chunkyList List

	for  {
        // read content to buffer
    	_, err := file.Read(buf)
		fmt.Print(buf)
		chunkyList.Insert(buf)
        if err != nil {
			if err != io.EOF {
				fmt.Println(err)
            }
            break
        } // if
		// create insert into list
	} // for
	
	Show(&chunkyList)
	return chunkyList
} // breakdown

func dechunk(output string, chunkList List) {
	// checks if file exists
	_, error := os.Stat(output)
	if !errors.Is(error, os.ErrNotExist) {
		os.Remove(output)
	} // if
	file,error := os.Create(output)
	if error != nil {
		log.Fatal(error)
		} // if
	defer file.Close()

	// loop over all chunks
	// for chunkList.head != nil {
	// 	byteChunk, err := GetBytes(chunkList.head.key)
	// 	if err != nil { fmt.Println(err) }
	// 	print
	// 	chunkList.head = chunkList.head.next
	// } // for
	// file.Write()
	// delete files
	
	
} // dechunk

func main() {
	chunky := chunk("help.txt", chunkSize)
	// dechunk("help_output.txt",chunky)
	fmt.Println(chunky)
}