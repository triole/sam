package transform

import "testing"

func TestDirName(t *testing.T) {
	assert(tr.DirName("/home/username/document.txt"), "/home/username", t)
	assert(tr.DirName("/home/username"), "/home", t)
	assert(tr.DirName("/home/username/"), "/home", t)
	assert(tr.DirName("/home/username///"), "/home", t)
}

func TestTidyFilePath(t *testing.T) {
	assert(tr.TidyFilePath("/tmp/HEllO   WORLD___.md"), "/tmp/hello_world_.md", t)
	assert(tr.TidyFilePath("/tmp/A B C D E.docx"), "/tmp/a_b_c_d_e.docx", t)
	assert(tr.TidyFilePath("/tmp/1!2@3#4$5.docx"), "/tmp/1_2_3_4_5.docx", t)
	assert(tr.TidyFilePath("/tmp/1!!!2!!!.docx"), "/tmp/1_2_.docx", t)
	assert(tr.TidyFilePath("/tmp/ä-_-Ö_-_ü.md"), "/tmp/ae_oe_ue.md", t)
}
