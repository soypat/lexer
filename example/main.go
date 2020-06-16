package main

import (
	"fmt"
	"github.com/soypat/lexer"
	"os"
)

const args = `@(X,U,t)`
const f = `[X(7);X(8);X(9);X(10);X(11);X(12);(U(1)*sin(U(2))*(cos(X(4))*sin(X(6))-cos(X(6))*sin(X(4))*sin(X(5))))/4+(U(1)*cos(U(3)+U(2))*(sin(X(4))*sin(X(6))+cos(X(4))*cos(X(6))*sin(X(5))))/4+(U(1)*cos(X(5))*cos(X(6))*sin(U(3)))/4;(U(1)*cos(X(5))*sin(X(6))*sin(U(3)))/4-(U(1)*cos(U(3)+U(2))*(cos(X(6))*sin(X(4))-cos(X(4))*sin(X(5))*sin(X(6))))/4-(U(1)*sin(U(2))*(cos(X(4))*cos(X(6))+sin(X(4))*sin(X(5))*sin(X(6))))/4;(U(1)*cos(U(3)+U(2))*cos(X(4))*cos(X(5)))/4-(U(1)*sin(X(5))*sin(U(3)))/4-(U(1)*cos(X(5))*sin(X(4))*sin(U(2)))/4-979/100;(cos(X(4))*sin(X(6))-cos(X(6))*sin(X(4))*sin(X(5)))*((3*U(1)*sin(U(3)))/20-(39271*X(12)^2*cos(X(4))*cos(X(5))*sin(X(5)))/40000+(39271*X(11)^2*cos(X(4))*cos(X(5))*sin(X(5))*sin(X(6))^2)/40000+(39271*X(10)*X(12)*cos(X(4))*cos(X(5))^2*cos(X(6)))/40000-(39271*X(10)*X(11)*cos(X(5))*cos(X(6))^2*sin(X(4)))/40000-(39271*X(10)*X(12)*cos(X(4))*cos(X(6))*sin(X(5))^2)/40000+(39271*X(11)*X(12)*cos(X(4))*cos(X(5))^2*sin(X(6)))/40000+(39271*X(10)*X(11)*cos(X(5))*sin(X(4))*sin(X(6))^2)/40000-(39271*X(11)*X(12)*cos(X(4))*sin(X(5))^2*sin(X(6)))/40000+(39271*X(10)^2*cos(X(5))*cos(X(6))*sin(X(4))*sin(X(6)))/40000-(39271*X(11)^2*cos(X(5))*cos(X(6))*sin(X(4))*sin(X(6)))/40000+(39271*X(11)*X(12)*cos(X(6))*sin(X(4))*sin(X(5)))/40000-(39271*X(10)*X(12)*sin(X(4))*sin(X(5))*sin(X(6)))/40000+(39271*X(10)^2*cos(X(4))*cos(X(5))*cos(X(6))^2*sin(X(5)))/40000+(39271*X(10)*X(11)*cos(X(4))*cos(X(5))*cos(X(6))*sin(X(5))*sin(X(6)))/20000)+cos(X(5))*cos(X(6))*((39271*X(12)^2*cos(X(4))*cos(X(5))^2*sin(X(4)))/40000-(39271*X(11)^2*cos(X(4))*cos(X(6))^2*sin(X(4)))/40000-(3*U(1)*sin(U(2)))/20-(39271*X(10)^2*cos(X(4))*sin(X(4))*sin(X(6))^2)/40000-(39271*X(10)^2*cos(X(4))^2*cos(X(6))*sin(X(5))*sin(X(6)))/40000+(39271*X(11)^2*cos(X(4))^2*cos(X(6))*sin(X(5))*sin(X(6)))/40000+(39271*X(10)^2*cos(X(6))*sin(X(4))^2*sin(X(5))*sin(X(6)))/40000-(39271*X(11)^2*cos(X(6))*sin(X(4))^2*sin(X(5))*sin(X(6)))/40000+(39271*X(11)*X(12)*cos(X(4))^2*cos(X(5))*cos(X(6)))/40000-(39271*X(10)*X(12)*cos(X(4))^2*cos(X(5))*sin(X(6)))/40000-(39271*X(11)*X(12)*cos(X(5))*cos(X(6))*sin(X(4))^2)/40000+(39271*X(10)*X(12)*cos(X(5))*sin(X(4))^2*sin(X(6)))/40000+(39271*X(10)^2*cos(X(4))*cos(X(6))^2*sin(X(4))*sin(X(5))^2)/40000+(39271*X(11)^2*cos(X(4))*sin(X(4))*sin(X(5))^2*sin(X(6))^2)/40000+(39271*X(10)*X(11)*cos(X(4))^2*cos(X(6))^2*sin(X(5)))/40000-(39271*X(10)*X(11)*cos(X(4))^2*sin(X(5))*sin(X(6))^2)/40000-(39271*X(10)*X(11)*cos(X(6))^2*sin(X(4))^2*sin(X(5)))/40000+(39271*X(10)*X(11)*sin(X(4))^2*sin(X(5))*sin(X(6))^2)/40000+(39271*X(10)*X(11)*cos(X(4))*cos(X(6))*sin(X(4))*sin(X(6)))/20000+(39271*X(10)*X(11)*cos(X(4))*cos(X(6))*sin(X(4))*sin(X(5))^2*sin(X(6)))/20000+(39271*X(10)*X(12)*cos(X(4))*cos(X(5))*cos(X(6))*sin(X(4))*sin(X(5)))/20000+(39271*X(11)*X(12)*cos(X(4))*cos(X(5))*sin(X(4))*sin(X(5))*sin(X(6)))/20000)+(800*U(1)*U(4)*(sin(X(4))*sin(X(6))+cos(X(4))*cos(X(6))*sin(X(5))))/243;cos(X(5))*sin(X(6))*((39271*X(12)^2*cos(X(4))*cos(X(5))^2*sin(X(4)))/40000-(39271*X(11)^2*cos(X(4))*cos(X(6))^2*sin(X(4)))/40000-(3*U(1)*sin(U(2)))/20-(39271*X(10)^2*cos(X(4))*sin(X(4))*sin(X(6))^2)/40000-(39271*X(10)^2*cos(X(4))^2*cos(X(6))*sin(X(5))*sin(X(6)))/40000+(39271*X(11)^2*cos(X(4))^2*cos(X(6))*sin(X(5))*sin(X(6)))/40000+(39271*X(10)^2*cos(X(6))*sin(X(4))^2*sin(X(5))*sin(X(6)))/40000-(39271*X(11)^2*cos(X(6))*sin(X(4))^2*sin(X(5))*sin(X(6)))/40000+(39271*X(11)*X(12)*cos(X(4))^2*cos(X(5))*cos(X(6)))/40000-(39271*X(10)*X(12)*cos(X(4))^2*cos(X(5))*sin(X(6)))/40000-(39271*X(11)*X(12)*cos(X(5))*cos(X(6))*sin(X(4))^2)/40000+(39271*X(10)*X(12)*cos(X(5))*sin(X(4))^2*sin(X(6)))/40000+(39271*X(10)^2*cos(X(4))*cos(X(6))^2*sin(X(4))*sin(X(5))^2)/40000+(39271*X(11)^2*cos(X(4))*sin(X(4))*sin(X(5))^2*sin(X(6))^2)/40000+(39271*X(10)*X(11)*cos(X(4))^2*cos(X(6))^2*sin(X(5)))/40000-(39271*X(10)*X(11)*cos(X(4))^2*sin(X(5))*sin(X(6))^2)/40000-(39271*X(10)*X(11)*cos(X(6))^2*sin(X(4))^2*sin(X(5)))/40000+(39271*X(10)*X(11)*sin(X(4))^2*sin(X(5))*sin(X(6))^2)/40000+(39271*X(10)*X(11)*cos(X(4))*cos(X(6))*sin(X(4))*sin(X(6)))/20000+(39271*X(10)*X(11)*cos(X(4))*cos(X(6))*sin(X(4))*sin(X(5))^2*sin(X(6)))/20000+(39271*X(10)*X(12)*cos(X(4))*cos(X(5))*cos(X(6))*sin(X(4))*sin(X(5)))/20000+(39271*X(11)*X(12)*cos(X(4))*cos(X(5))*sin(X(4))*sin(X(5))*sin(X(6)))/20000)-(cos(X(4))*cos(X(6))+sin(X(4))*sin(X(5))*sin(X(6)))*((3*U(1)*sin(U(3)))/20-(39271*X(12)^2*cos(X(4))*cos(X(5))*sin(X(5)))/40000+(39271*X(11)^2*cos(X(4))*cos(X(5))*sin(X(5))*sin(X(6))^2)/40000+(39271*X(10)*X(12)*cos(X(4))*cos(X(5))^2*cos(X(6)))/40000-(39271*X(10)*X(11)*cos(X(5))*cos(X(6))^2*sin(X(4)))/40000-(39271*X(10)*X(12)*cos(X(4))*cos(X(6))*sin(X(5))^2)/40000+(39271*X(11)*X(12)*cos(X(4))*cos(X(5))^2*sin(X(6)))/40000+(39271*X(10)*X(11)*cos(X(5))*sin(X(4))*sin(X(6))^2)/40000-(39271*X(11)*X(12)*cos(X(4))*sin(X(5))^2*sin(X(6)))/40000+(39271*X(10)^2*cos(X(5))*cos(X(6))*sin(X(4))*sin(X(6)))/40000-(39271*X(11)^2*cos(X(5))*cos(X(6))*sin(X(4))*sin(X(6)))/40000+(39271*X(11)*X(12)*cos(X(6))*sin(X(4))*sin(X(5)))/40000-(39271*X(10)*X(12)*sin(X(4))*sin(X(5))*sin(X(6)))/40000+(39271*X(10)^2*cos(X(4))*cos(X(5))*cos(X(6))^2*sin(X(5)))/40000+(39271*X(10)*X(11)*cos(X(4))*cos(X(5))*cos(X(6))*sin(X(5))*sin(X(6)))/20000)-(800*U(1)*U(4)*(cos(X(6))*sin(X(4))-cos(X(4))*sin(X(5))*sin(X(6))))/243;(800*U(1)*U(4)*cos(X(4))*cos(X(5)))/243-cos(X(5))*sin(X(4))*((3*U(1)*sin(U(3)))/20-(39271*X(12)^2*cos(X(4))*cos(X(5))*sin(X(5)))/40000+(39271*X(11)^2*cos(X(4))*cos(X(5))*sin(X(5))*sin(X(6))^2)/40000+(39271*X(10)*X(12)*cos(X(4))*cos(X(5))^2*cos(X(6)))/40000-(39271*X(10)*X(11)*cos(X(5))*cos(X(6))^2*sin(X(4)))/40000-(39271*X(10)*X(12)*cos(X(4))*cos(X(6))*sin(X(5))^2)/40000+(39271*X(11)*X(12)*cos(X(4))*cos(X(5))^2*sin(X(6)))/40000+(39271*X(10)*X(11)*cos(X(5))*sin(X(4))*sin(X(6))^2)/40000-(39271*X(11)*X(12)*cos(X(4))*sin(X(5))^2*sin(X(6)))/40000+(39271*X(10)^2*cos(X(5))*cos(X(6))*sin(X(4))*sin(X(6)))/40000-(39271*X(11)^2*cos(X(5))*cos(X(6))*sin(X(4))*sin(X(6)))/40000+(39271*X(11)*X(12)*cos(X(6))*sin(X(4))*sin(X(5)))/40000-(39271*X(10)*X(12)*sin(X(4))*sin(X(5))*sin(X(6)))/40000+(39271*X(10)^2*cos(X(4))*cos(X(5))*cos(X(6))^2*sin(X(5)))/40000+(39271*X(10)*X(11)*cos(X(4))*cos(X(5))*cos(X(6))*sin(X(5))*sin(X(6)))/20000)-sin(X(5))*((39271*X(12)^2*cos(X(4))*cos(X(5))^2*sin(X(4)))/40000-(39271*X(11)^2*cos(X(4))*cos(X(6))^2*sin(X(4)))/40000-(3*U(1)*sin(U(2)))/20-(39271*X(10)^2*cos(X(4))*sin(X(4))*sin(X(6))^2)/40000-(39271*X(10)^2*cos(X(4))^2*cos(X(6))*sin(X(5))*sin(X(6)))/40000+(39271*X(11)^2*cos(X(4))^2*cos(X(6))*sin(X(5))*sin(X(6)))/40000+(39271*X(10)^2*cos(X(6))*sin(X(4))^2*sin(X(5))*sin(X(6)))/40000-(39271*X(11)^2*cos(X(6))*sin(X(4))^2*sin(X(5))*sin(X(6)))/40000+(39271*X(11)*X(12)*cos(X(4))^2*cos(X(5))*cos(X(6)))/40000-(39271*X(10)*X(12)*cos(X(4))^2*cos(X(5))*sin(X(6)))/40000-(39271*X(11)*X(12)*cos(X(5))*cos(X(6))*sin(X(4))^2)/40000+(39271*X(10)*X(12)*cos(X(5))*sin(X(4))^2*sin(X(6)))/40000+(39271*X(10)^2*cos(X(4))*cos(X(6))^2*sin(X(4))*sin(X(5))^2)/40000+(39271*X(11)^2*cos(X(4))*sin(X(4))*sin(X(5))^2*sin(X(6))^2)/40000+(39271*X(10)*X(11)*cos(X(4))^2*cos(X(6))^2*sin(X(5)))/40000-(39271*X(10)*X(11)*cos(X(4))^2*sin(X(5))*sin(X(6))^2)/40000-(39271*X(10)*X(11)*cos(X(6))^2*sin(X(4))^2*sin(X(5)))/40000+(39271*X(10)*X(11)*sin(X(4))^2*sin(X(5))*sin(X(6))^2)/40000+(39271*X(10)*X(11)*cos(X(4))*cos(X(6))*sin(X(4))*sin(X(6)))/20000+(39271*X(10)*X(11)*cos(X(4))*cos(X(6))*sin(X(4))*sin(X(5))^2*sin(X(6)))/20000+(39271*X(10)*X(12)*cos(X(4))*cos(X(5))*cos(X(6))*sin(X(4))*sin(X(5)))/20000+(39271*X(11)*X(12)*cos(X(4))*cos(X(5))*sin(X(4))*sin(X(5))*sin(X(6)))/20000)]`
const f2 = `(U(1)*sin(U(2))*(cos(X(4))*sin(X(6))-cos(X(6))*sin(X(4))*sin(X(5))))/4+(U(1)*cos(U(3)+U(2))*(sin(X(4))*sin(X(6))+cos(X(4))*cos(X(6))*sin(X(5))))/4+(U(1)*cos(X(5))*cos(X(6))*sin(U(3)))/4`
const fs = `U(X(4))*2`


func main() {
	if err := run(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

// actual program
func run() error {
	lexy := lexer.NewStringLexer("matlab func", f)

	// add identifiers
	var variables = []string{"X","U"}
	var functions = []string{"sin","cos"}
	for _,v:= range variables {
		_ = lexy.NewVariableID(v)
	}
	for _,v:= range functions {
		_ = lexy.NewFunctionID(v)
	}
	// lexer pushes item tokens to channel. parse picks em up
	parser := parse(lexy.ItemChannel())
	go lexy.Run()
	parser.run()

	fo, err := os.Create("output.txt")
	if err != nil {
		return err
	}
	defer fo.Close()
	parser.listStructure(fo)
	return nil
}
