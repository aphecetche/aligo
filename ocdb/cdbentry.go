package ocdb

import (
	"fmt"
	"io"

	"go-hep.org/x/hep/groot/root"
)

func (entry *Entry) Object() root.Object { return entry.obj }
func (entry *Entry) Id() ID              { return entry.id }

func (entry *Entry) Display(w io.Writer) {
	fmt.Fprintf(w, `=== Entry ===
ID: %v
Owner: %v
`,
		entry.id, entry.owner,
	)
	if entry.meta != nil {
		fmt.Fprintf(w, "MetaData:\n")
		entry.meta.Display(w)
	}
	if entry.obj != nil {
		fmt.Fprintf(w, "Object: %T\n%v\n===\n", entry.obj, entry.obj)
	}
}
