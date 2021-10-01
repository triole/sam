package transform

import "testing"

func TestDirName(t *testing.T) {
	assert(tr.DirName("/home/username/document.txt"), "/home/username", t)
	assert(tr.DirName("/home/username"), "/home", t)
	assert(tr.DirName("/home/username/"), "/home", t)
	assert(tr.DirName("/home/username///"), "/home", t)
}

func TestTidyFileName1(t *testing.T) {
	assert(tr.TidyFileName1("///etc//crontab"), "/etc/crontab", t)
	assert(tr.TidyFileName1("/tmp/HEllO   WORLD___.md"), "/tmp/HEllO   WORLD___.md", t)
	assert(tr.TidyFileName1("/tmp/A B C D E.docx"), "/tmp/A B C D E.docx", t)
	assert(tr.TidyFileName1("/tmp//1!2@3#4$5.docx"), "/tmp/1!2@3#4$5.docx", t)
	assert(tr.TidyFileName1("//tmp///1!!!2!!!.docx"), "/tmp/1!!!2!!!.docx", t)
}

func TestTidyFileName2(t *testing.T) {
	assert(tr.TidyFileName2("/tmp/HEllO   WORLD___.md"), "/tmp/HEllO___WORLD___.md", t)
	assert(tr.TidyFileName2("/tmp/A B C D E.docx"), "/tmp/A_B_C_D_E.docx", t)
	assert(tr.TidyFileName2("/tmp//1!2@3#4$5.docx"), "/tmp/1_2_3_4_5.docx", t)
	assert(tr.TidyFileName2("//tmp///1!!!2!!!.docx"), "/tmp/1___2___.docx", t)

	assert(tr.TidyFileName2("ùûüÿ€àâæçéèêëïîôœ"), "uuuey_aa_ceeeeiio_", t)
}

func TestTidyFileName3(t *testing.T) {
	assert(tr.TidyFileName3("/tmp/HEllO   WORLD___.md"), "/tmp/hello___world___.md", t)
	assert(tr.TidyFileName3("/tmp/A B C D E.docx"), "/tmp/a_b_c_d_e.docx", t)
	assert(tr.TidyFileName3("/tmp//1!2@3#4$5.docx"), "/tmp/1_2_3_4_5.docx", t)
	assert(tr.TidyFileName3("//tmp///1!!!2!!!.docx"), "/tmp/1___2___.docx", t)
}

func TestTidyFileName4(t *testing.T) {
	assert(tr.TidyFileName4("/tmp/HEllO   WORLD___.md"), "/tmp/hello_world_.md", t)
	assert(tr.TidyFileName4("/tmp/A B C D E.docx"), "/tmp/a_b_c_d_e.docx", t)
	assert(tr.TidyFileName4("/tmp/1!2@3#4$5.docx"), "/tmp/1_2_3_4_5.docx", t)
	assert(tr.TidyFileName4("/tmp/1!!!2!!!.docx"), "/tmp/1_2_.docx", t)
}
