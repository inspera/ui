// 11 december 2015

package ui

// #include "pkgui.h"
import "C"

// Quit queues a return from Main. It does not exit the program.
// It also does not immediately cause Main to return; Main will
// return when it next can. Quit must be called from the GUI thread.
func Quit() {
	C.uiQuit()
}
