package runtime

import (
	"fmt"
)

func NewSystem() *System {
	return &System{Out: &PrintStream{}}
}

func (s *System) Call(Method string, args ...any) (any, error) {
	switch Method {
	case"java.lang.System.out.println":
		s.Out.println(args[0])
		return nil, nil
	}
	return nil, fmt.Errorf("Method not found: %s", Method)
}

type System struct {
	Out *PrintStream
}

type PrintStream struct {}

func (p *PrintStream) println(args ...any) {
	for _, arg := range args {
		fmt.Printf("%v ", arg)
	}
	fmt.Print("\n")
}
