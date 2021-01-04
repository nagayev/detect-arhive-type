//Package dat is a package for detection arhive type
// Copyright 2021 Marat Nagayev <nagaevmt49t@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
package dat

import (
	"os"
	"bufio"
	"bytes"
)
const (
	//UNKNOWN file, default value
	UNKNOWN=0
	//TARGZ is .tar.gz file
	TARGZ=1
	//ZIP is .zip file
	ZIP=2
	//GZ is .gz file
	GZ=3
	//Z7 is .7z file (7Zip)
	Z7=4
)

func getTypeCode(path string) (int,error) {
	//You can find the list of signatures here:
	//https://en.wikipedia.org/wiki/List_of_file_signatures
	targzSignature:=[]byte{31,139}
	zipSignature:=[]byte{0x50,0x4B,0x03,0x04}
	gzSignature:=[]byte{0x1f,0x8b,0x08}
	z7Signature:=[]byte{0x37,0x7a,0xbc,0xaf,0x27,0x1c}

	file, err := os.Open(path)
	if err!=nil{
		return UNKNOWN,err
	}
	fileType:=UNKNOWN
	reader := bufio.NewReader(file)
	buffer:=make([]byte,8)
	//NOTE: elements' order is special (same as constants) 
	formats:=[][]byte{targzSignature,zipSignature,gzSignature,z7Signature}
	reader.Read(buffer)
	for i,format:=range formats{
		if bytes.HasPrefix(buffer,format){
			fileType=i+1
			break
		}
	}
	return fileType,nil
}
func getTypeString(path string) (string, error) {
	code,err:=getTypeCode(path)
	if err!=nil{
		return "unknown",err
	}
	typesAsString:=[]string{"unknown","tar.gz","zip","gz","7z"}
	//len(typesAsString) should be equal to len(formats)
	return typesAsString[code],nil
}