// This is a makeshift parser for use with example. Use at own risk
package main

import (
	"fmt"
	lex "github.com/soypat/lexer"
	"os"
)

type parser struct {
	items <-chan lex.Item
	itemList []lex.Item
}

func parse(items <-chan lex.Item) *parser {
	p := parser{items: items}
	return &p
}

func (p *parser)run() {
	for {
		j := 0
		select {
		case item := <- p.items:
			switch item.Type() {
			case lex.ItemError:
				fmt.Printf("Recieved error in parser\n%s",item.Value())
				return
			case lex.ItemText:
				j++
				continue
			case lex.ItemEOF:
				fmt.Printf("EOF reached. %d texts found",j)
				return
			default:
				p.itemList = append(p.itemList,item)
			}
		}
	}
}

func (p *parser)listStructure(f *os.File) {
	for _,v:= range p.itemList {
		_,_ = f.WriteString(v.Value() + fmt.Sprintf("\t%s",v.Type()))
		f.WriteString("\n")
	}
}