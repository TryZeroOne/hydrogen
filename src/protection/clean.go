package protection

import "os/exec"

func cleanHeader(file string) {

	exec.Command("strip", "-sxX",
		"--remove-section=.bss",
		"--remove-section=.comment",
		"--remove-section=.eh_frame",
		"--remove-section=.eh_frame_hdr",
		"--remove-section=.fini",
		"--remove-section=.fini_array",
		"--remove-section=.gnu.build.attributes",
		"--remove-section=.gnu.hash",
		"--remove-section=.gnu.version",
		"--remove-section=.gosymtab",
		"--remove-section=.got",
		"--remove-section=.note.ABI-tag",
		"--remove-section=.note.gnu.build-id",
		"--remove-section=.shstrtab",
		"--remove-section=.typelink", file).Run()

}
