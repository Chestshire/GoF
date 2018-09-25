package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/supertech/common/logs"
)

type sudoku struct {
	row  []map[int]bool
	col  []map[int]bool
	sqr  []map[int]bool
	sudo [][]int
}

const (
	gong    = 9
	logName = "sudoku.log"
)

//sudoku 单例
var (
	S sudoku
)

//----init-----
func init() {
	//初始化 log 文件
	logs.L.InitLogger(logName, "debug")
	//初始化 sudoku 单例
	sudokuInit()
}

//初始化 sudoku 单例
func sudokuInit() {
	S = sudoku{
		row:  make([]map[int]bool, 0, gong),
		col:  make([]map[int]bool, 0, gong),
		sqr:  make([]map[int]bool, 0, gong),
		sudo: make([][]int, 0, gong),
	}

	//三元素的初始化
	for i := 0; i < gong; i++ {
		for j := 0; j < gong; j++ {
			S.row[i][j] = true
			S.col[i][j] = true
			S.sqr[i][j] = true
		}
	}

	gongArray := make([]int, 0, gong)

	for row := 0; row < gong; row++ {
		S.sudo = append(S.sudo, gongArray)
	}
}

//---------------------------------------- 程序入口 ---------------------------------------
func main() {
	RunOperation()
}

//--------------------------------------运行封装函数-------------------------------------

//RunOperation 开始运行
func RunOperation() {
	inputReader := S.readInput()
	//将sudoku的值带入
	S.inputSudoku(inputReader)
	//将sudoku的其他元素列表初始化
	S.sudokuInitSet()

	S.writeOutput()
}

//-------------------------------- 初始设置数独结构体 --------------------------------
func (s *sudoku) sudokuInitSet() {

	//设置 row
	for row := 0; row < gong; row++ {
		for col := 0; col < gong; col++ {
			numExists := S.sudo[row][col]
			if numExists != 0 { //如果该位置的数存在则删除它
				deleteElement(S.row, row, numExists)
			}
		}
	}

	//设置 col
	for col := 0; col < gong; col++ {
		for row := 0; row < gong; row++ {

		}
	}

	//设置 sqr
	for i := 0; i < gong; i++ {
		for j := 0; j < gong; j++ {

		}
	}

}

//inputSudoku 将sudoku的值带入
func (s *sudoku) inputSudoku(inputReader *bufio.Reader) {
	for row := 0; row < gong; row++ {
		buf, errin := inputReader.ReadString('\n')
		if errin != nil {
			logs.L.Error("输入有问题")
		}

		rowArray := make([]int, 0, 9)

		//给 sudoku 切片增加元素
		for col := 0; col < gong; col++ {
			num, errcon := strconv.Atoi(buf[col : col+1])
			if errcon != nil {
				logs.L.Fatal("输入数据中出现了非数字")
			}
			rowArray = append(rowArray, num)
		}
		s.sudo[row] = rowArray
	}
}

//----------------------------------------- 命令行信息 ------------------------------------
//readInput 对话框接受信息
func (s *sudoku) readInput() (inputReader *bufio.Reader) {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("请输入数据")
	return
}

//writeOutput 对话框返回信息
func (s *sudoku) writeOutput() {
	fmt.Printf("\n输出为：\n")

	for row := 0; row < gong; row++ {
		for col := 0; col < gong; col++ {
			fmt.Printf("%v", s.sudo[row][col])
		}
		fmt.Println()
	}

	//带上方框

	/* 	fmt.Println("带上方框：")
	   	for row := 0; row < gong; row++ {
	   		if row > 0 && (row%3) == 0 {
	   			fmt.Printf("-----------------------------\n")
	   		}
	   		for col := 0; col < gong; col++ {
	   			if col > 0 && (col%3) == 0 {
	   				fmt.Printf("|")
	   			}
	   			fmt.Printf(" %v ", s.sudo[row][col])
	   		}
	   		fmt.Printf("\n")
	   	}
	   	fmt.Println() */

}

//------------------------------------------ 运算函数 ------------------------------------
//deleteElement "删除"一个元素
func deleteElement(element []map[int]bool, row, col int) {
	element[row][col] = false
}

//sudokuOperation 数独的运算
func sudokuOperation() {

	//关闭通道
}

//sdkOperRecur 数独递归函数
func sdkOperRecur() {

	//return 写完

	//如果检测数独已经完成，则使用通道传输结果
}

//------------------------------------------ 工具函数 ------------------------------------
//deleteElement 根据row/col/sqr 以及第几个 和 坐标 删除元素

//updateSudoku 更新 数独内容

//getInfo 根据row/col/sqr 以及第几个 查询 返回 长度，一个内容切片

//copySudoku 返回备份的数独结构体信息

//resultCheck 查看sudoku是否已经完成
func resultCheck() {

	//如果完成了，但数据不合法则报错（完全不应该发生）
}

/*---------------------------------------------- 文档 -----------------------------------------
row 1 - 9		[]map
col 1 - 9
sqr	1 - 9

sudoku 81 个元素的 切片 存放数据

一个主要的递归函数 给予	一个 sudoku 切片， row col sqr 为了减少运算也会给予 但这里可能有内存占用过多的问题

读进数据

输出数据
	通过通道传输结果

查询指定row/col/sqr的长度(可选数字量)
排序，然后从最短的开始走

删除 row/col/sqr 中的某个元素

*/
