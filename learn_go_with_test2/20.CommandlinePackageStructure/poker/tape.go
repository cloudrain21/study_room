package poker

import (
	"io"
	"os"
)

type tape struct {
	file *os.File
}

func (t *tape)Write(p []byte) (int,error) {
	t.file.Truncate(0)
	t.file.Seek(0,io.SeekStart)
	return t.file.Write(p)
}
