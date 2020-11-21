package main

import (
	"./send_to_js"
	"fmt"
	"github.com/ganeshrvel/go-mtpx"
	jsoniter "github.com/json-iterator/go"
)

//	#include "stdint.h"
import "C"

var container deviceContainer

//export Initialize
func Initialize(ptr int64) {
	_, err := _initialize(mtpx.Init{DebugMode: false})
	if err != nil {
		send_to_js.SendError(ptr, err)

		return
	}

	dInfo, err := _fetchDeviceInfo()
	if err != nil {
		send_to_js.SendError(ptr, err)

		return
	}

	send_to_js.SendInitialize(ptr, dInfo)
}

//export FetchDeviceInfo
func FetchDeviceInfo(ptr int64) {
	dInfo, err := _fetchDeviceInfo()
	if err != nil {
		send_to_js.SendError(ptr, err)

		return
	}

	send_to_js.SendDeviceInfo(ptr, dInfo)
}

//export FetchStorages
func FetchStorages(ptr int64) {
	storages, err := _fetchStorages()
	if err != nil {
		send_to_js.SendError(ptr, err)

		return
	}

	send_to_js.SendStorages(ptr, storages)
}

//export MakeDirectory
func MakeDirectory(ptr int64, json *C.char) {
	i := MakeDirectoryInput{}

	var j = jsoniter.ConfigFastest
	err := j.UnmarshalFromString(C.GoString(json), &i)
	if err != nil {
		send_to_js.SendError(ptr, fmt.Errorf("error occured while Unmarshalling MakeDirectory input data %+v: ", err))

		return
	}

	if err := _makeDirectory(i.StorageId, i.FullPath); err != nil {
		send_to_js.SendError(ptr, err)

		return
	}

	send_to_js.SendMakeDirectory(ptr)
}

//export FileExists
func FileExists(ptr int64, json *C.char) {
	i := FileExistsInput{}

	var j = jsoniter.ConfigFastest
	err := j.UnmarshalFromString(C.GoString(json), &i)
	if err != nil {
		send_to_js.SendError(ptr, fmt.Errorf("error occured while Unmarshalling FileExists input data %+v: ", err))

		return
	}

	var fProps []mtpx.FileProp
	for _, f := range i.Files {
		fProp := mtpx.FileProp{FullPath: f}

		fProps = append(fProps, fProp)
	}

	fc, err := _fileExists(i.StorageId, fProps)

	if err != nil {
		send_to_js.SendError(ptr, err)

		return
	}

	send_to_js.FileExists(ptr, fc, i.Files)
}

//export Dispose
func Dispose(ptr int64) {
	if err := _dispose(); err != nil {
		send_to_js.SendError(ptr, err)

		return
	}

	container.dev = nil
	container.deviceInfo = nil

	send_to_js.SendDispose(ptr)
}

////export Walk
//func Walk(fullPath *C.char, storageId int, ptr int64) {
//	_fullPath := C.GoString(fullPath)
//	pretty.Println("_fullPath:", _fullPath)
//	pretty.Println("storageId:", uint32(storageId))
//
//	decode_native.SendWalkResult(ptr, d, uint32(storageId), _fullPath)
//}
//
////export UploadFiles
//func UploadFiles(ptr int64) {
//	start := time.Now()
//	uploadFile1 := getTestMocksAsset("test-large-files")
//	sources := []string{uploadFile1}
//	destination := "/mtp-test-files/temp_dir/test_UploadFiles"
//	objectIdDest, totalFiles, totalSize, err := mtpx.UploadFiles(d, Sid,
//		sources,
//		destination,
//		func(fi *mtpx.TransferredFileInfo, err error) error {
//			pretty.Printf("Filename: %s\n", fi.FileInfo.Name)
//			pretty.Printf("Speed: %d MB/s\n", fi.Speed)
//
//			decode_native.UploadFilesCb(ptr, fi)
//			return nil
//		},
//	)
//
//	if err != nil {
//		return
//	}
//
//	pretty.Println(objectIdDest)
//	pretty.Println(totalFiles)
//	pretty.Println(totalSize)
//	pretty.Println("time elapsed: ", time.Since(start).Seconds())
//}
//
///// [sourcesJson] source list string slice in json format
////export DownloadFiles
//func DownloadFiles(sourcesJson, destination *C.char, ptr int64) {
//	var sources []string
//
//	_sourcesJson := C.GoString(sourcesJson)
//	_destination := C.GoString(destination)
//
//	var json = jsoniter.ConfigFastest
//	err := json.UnmarshalFromString(_sourcesJson, &sources)
//	if err != nil {
//		return
//	}
//
//	_, _, err = mtpx.DownloadFiles(d, Sid,
//		sources, _destination,
//		func(downloadFi *mtpx.TransferredFileInfo, err error) error {
//			fmt.Printf("Current filepath: %s\n", downloadFi.FileInfo.FullPath)
//			fmt.Printf("%f MB/s\n", downloadFi.Speed)
//
//			return nil
//		},
//	)
//	if err != nil {
//		return
//	}
//
//	//pretty.Println(totalFiles)
//	//pretty.Println(totalSize)
//	//pretty.Println("time elapsed: ", time.Since(start).Seconds())
//}

func main() {}
