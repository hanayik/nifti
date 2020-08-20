package main

import (
	"fmt"
	"nifti"
)

func main() {
	path := "../test_data/sub-01 anat sub-01_T1w.nii"
	//path := "MNI152_T1_1mm_nifti2.nii"

	if nifti.ReadNiftiType(path) == 1 {
		header := nifti.ReadNifti1Header(path)
		nifti.PrintNifti1Header(header)
	} else {
		fmt.Println("Not a Nifti1 file")
	}
}
