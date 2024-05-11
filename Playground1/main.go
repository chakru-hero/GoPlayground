package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

func main() {
	// var x int16;
	// fmt.Println("1. Swap two Numbers.")
	// fmt.Println("2. Find longest non repeating substring.")
	// fmt.Print("Select an Option : ");
	//  fmt.Scan(&x);
	// switch x{
	// 	case 1  :
	// 		basicFuncs.SwapTwoNumbers();
	// 	case 2 :
	// 		basicFuncs.NonRepeatingSubString();
	// 	default :
	// 		fmt.Println("Select an option from above");

	f, err := os.Open("/dev/input/event6")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	b := make([]byte, 24)
	for {	
		f.Read(b)
		sec := binary.LittleEndian.Uint64(b[0:8])
		usec := binary.LittleEndian.Uint64(b[8:16])
		t := time.Unix(int64(sec), int64(usec))
		fmt.Println(t)
		var value int32
		typ := binary.LittleEndian.Uint16(b[16:18])
		code := binary.LittleEndian.Uint16(b[18:20])
		binary.Read(bytes.NewReader(b[20:]), binary.LittleEndian, &value)
		fmt.Printf("type: %x\ncode: %d\nvalue: %d\n", typ, code, value)
	}

}
