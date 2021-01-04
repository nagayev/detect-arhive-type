// Copyright 2015 Geofrey Ernest <geofreyernest@live.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package dat

import (
	"fmt"
	"testing"
)

func testGetTypeCode(path string,expectedCode int) error{
	actualCode,err:=getTypeCode(path)
	//expectedCode:=TARGZ
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
//TODO: 
func TestGZ(t *testing.T){}
//TODO:

func TestZ7(t *testing.T){
	err:=testGetTypeCode("test_files/test.7z",Z7)
	if err!=nil{
		t.Errorf("Failed with error %v",err)
	}
}
