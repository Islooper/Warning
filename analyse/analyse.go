package analyse

import (
	"github.com/Islooper/warning/push"
)

type WarningProcess struct {
	R chan string
	W chan string
	push.Writer
}

func NewWarningProcess(r chan string, w chan string, writer push.Writer) *WarningProcess {
	return &WarningProcess{R: r, W: w, Writer: writer}
}

// 分析工程
func (w *WarningProcess) Analyse() {
	for x := range w.R {
		w.W <- x
	}
}
