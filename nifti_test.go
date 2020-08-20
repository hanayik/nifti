package nifti

import (
	"testing"
)

func TestReadBytes(t *testing.T) {
	nifti1Path := "./test_data/sub-01 anat sub-01_T1w.nii"
	file := openFile(nifti1Path)
	defer file.Close()

	var n int32 = 4
	b := readBtyes(file, n)
	var wantBytes = 4

	if len(b) != wantBytes {
		t.Errorf("readBytes got: %d bytes, want: %d bytes.", len(b), wantBytes)
	}

}
func TestNifti1SizeOfHeader(t *testing.T) {
	// test that ReadNiftiType returns 1 for a
	// Nifti1 image formatted file
	nifti1Path := "./test_data/sub-01 anat sub-01_T1w.nii"
	ntype := ReadNiftiType(nifti1Path)
	if ntype != 1 {
		t.Errorf("ReadNiftiType got: %d, want: %d.", ntype, 1)
	}
}
