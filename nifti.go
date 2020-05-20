package main

import (
	"fmt"
	"log"
	"os"
	"bytes"
	"encoding/binary"
	"reflect"
)

const Nifti1HeaderBytes int32 = 348
const Nifti2HeaderBytes int32 = 540

type Nifti1Header struct {
	//            						offset	size
	SizeOfHdr 		int32			//	0B		4B
	DateType		[10]byte		//	4B		10B
	DBName			[18]byte		//	14B		18B
	Extents			int32			//	32B		4B
	SessionError	int16			//	36B		2B
	Regular			byte			//	38B		1B
	DimInfo			byte			//	39B		1B
	Dim				[8]int16		//	40B		16B
	IntentP1		float32			//	56B		4B
	IntentP2		float32			//	60B		4B
	IntentP3		float32			//	64B		4B
	IntentCode		int16			//	68B		2B
	DataType		int16			//	70B		2B
	BitPix			int16			//	72B		2B
	SliceStart		int16			//	74B		2B
	PixDim			[8]float32		//	76B		32B
	VoxOffset		float32			//	108B	4B
	SclSlope		float32			//	112B	4B
	SclInter		float32			//	116B	4B
	SliceEnd		int16			//	120B	2B
	SliceCode		byte			//	122B	1B
	XYZTUnits		byte			//	123B	1B
	CalMax			float32			//	124B	4B
	CalMin			float32			//	128B	4B
	SliceDuration	float32			//	132B	4B
	TOffset			float32			//	136B	4B
	GlMax			int32			//	140B	4B
	GlMin			int32			//	144B	4B
	Descrip			[80]byte		//	148B	80B
	AuxFile			[24]byte		//	228B	24B
	QFormCode		int16			//	252B	2B
	SFormCode		int16			//	254B	2B
	QuaternB		float32			//	256B	4B
	QuaternC		float32			//	260B	4B
	QuaternD		float32			//	264B	4B
	QOffsetX		float32			//	268B	4B
	QOffsetY		float32			//	272B	4B
	QOffsetZ		float32			//	276B	4B
	SRowX			[4]float32		//	280B	16B
	SRowY			[4]float32		//	296B	16B
	SRowZ			[4]float32		//	312B	16B
	IntentName		[16]byte		//	328B	16B
	Magic			[4]byte			//	344		4B
	//							total header size = 348B
}

func ReadBytes(file os.File, number int32) []byte {
	bytes := make([]byte, number)
	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

func ReadNiftiType(file os.File) int {
	// just read the first 4 bytes to get nifti type
	var nbytes int32 = 4
	var nType int32
	rawbytes := ReadBytes(file, nbytes)
	buffer := bytes.NewBuffer(rawbytes)
	err := binary.Read(buffer, binary.LittleEndian, &nType)
	if err != nil {
		log.Fatal("ReadNiftiType failed", err)
	}
	if nType == Nifti1HeaderBytes {
		return 1
	} else if nType == Nifti2HeaderBytes {
		return 2
	} else {
		return 0
	}
}

func ReadNifti1Header(file os.File) Nifti1Header {
	header := Nifti1Header{}
	rawbytes := ReadBytes(file, Nifti1HeaderBytes) 
	buffer := bytes.NewBuffer(rawbytes)
	err := binary.Read(buffer, binary.LittleEndian, &header)
	if err != nil {
		log.Fatal("ReadNifti1Header failed", err)
	}
	return header
}

func PrintNifti1Header(h Nifti1Header) {
	fields := reflect.TypeOf(h)
	values := reflect.ValueOf(h)
	num := fields.NumField()

	for i := 0; i < num; i++ {
		field := fields.Field(i)
		value := values.Field(i)
		fmt.Print("Type:", field.Type, ",", field.Name, "=", value, "\n")
	}
}

func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Error while opening file", err)
	}

	return file
}

func main() {
    path := "sub-01 anat sub-01_T1w.nii"

	file := OpenFile(path)
	defer file.Close()

	if ReadNiftiType(*file) == 1 {
		header := ReadNifti1Header(*file)
		PrintNifti1Header(header)
	} else {
		fmt.Println("Not a Nifti1 file")
	}
}