package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

const CORNER = "+"

const SIDE_LINE = "-"

const PARTITION = "|"

const SPACE = " "

const LF = "\r\n"

type Table struct {
	headerList    []string
	bodyList      [][]string
	separatorList []int
	colSize       int
}

func NewTable() *Table {
	t := &Table{}
	return t
}

func (t *Table) SetHeader(headerList []string) {
	t.headerList = headerList
}

func (t *Table) SetBody(bodyList [][]string) {
	t.bodyList = bodyList
}

func (t *Table) Render() {
	t.parseSeparator()
	t.writeHeader()
	t.writeBody()
	t.printSeparateLine()
}

func (t *Table) writeHeader() {
	t.printSeparateLine()

	t.printLF()
	t.printPartition()

	for i, v := range t.headerList {
		fmt.Printf(v)
		separator := t.separatorList[i]
		fmt.Printf(strings.Repeat(SPACE, separator-len(v)))
		t.printPartition()
	}

	t.printLF()

	t.printSeparateLine()

}

func (t *Table) writeBody() {
	t.printLF()

	for _, v := range t.bodyList {
		rows := v
		for i, v := range rows {
			t.printPartition()
			fmt.Printf(v)
			separator := t.separatorList[i]
			fmt.Printf(strings.Repeat(SPACE, separator-len(v)))
		}
		t.printPartition()
		t.printLF()
	}

}

func (t *Table) printSeparateLine() {
	fmt.Printf(CORNER)
	for _, v := range t.separatorList {
		fmt.Printf(strings.Repeat(SIDE_LINE, v))
		fmt.Printf(CORNER)
	}
}

func (t *Table) printLF() {
	fmt.Printf(LF)
}

func (t *Table) printPartition() {
	fmt.Printf(PARTITION)
}


func (t *Table) parseSeparator() {
	separatorList := make([]int, len(t.headerList))

	sumList := append([][]string{t.headerList}, t.bodyList...)

	for _, v := range sumList {
		rows := v
		for i, v := range rows {
			separator := separatorList[i]
			wordSize := utf8.RuneCountInString(v)
			//数値→文字列変換
			//fmt.Println("wordSize:"+strconv.Itoa(wordSize))
			if wordSize > separator {
				separatorList[i] = wordSize
			}
		}

	}
	t.separatorList = separatorList
}
