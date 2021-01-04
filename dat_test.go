//Package dat is a package for detection arhive type
// Copyright 2021 Marat Nagayev <nagaevmt49@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
package dat

import (
	"fmt"
	"testing"
)

func testGetTypeCode(path string, expectedCode int) error{
	actualCode,err:=getTypeCode(path)
	if err!=nil{
		return err
	}
	if actualCode!=expectedCode {
		err=fmt.Errorf("Invalid code %d, expected: %d",actualCode,expectedCode)
	}
	return err
}
func TestTarGZ(t *testing.T){
	err:=testGetTypeCode("test_files/test.tar.gz",TARGZ)
	if err!=nil{
		t.Errorf("Failed with error %v",err)
	}
}
func TestZip(t *testing.T){
	err:=testGetTypeCode("test_files/test.zip",ZIP)
	if err!=nil{
		t.Errorf("Failed with error %v",err)
	}
}
func TestZ7(t *testing.T){
	err:=testGetTypeCode("test_files/test.7z",Z7)
	if err!=nil{
		t.Errorf("Failed with error %v",err)
	}
}
func TestXZ(t *testing.T){
	err:=testGetTypeCode("test_files/test.tar.xz",XZ)
	if err!=nil{
		t.Errorf("Failed with error %v",err)
	}
}
