package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

// 定义全局变量
// var name string = "张三"
// var age int = 18;
var (
	name        = "李四"
	age         = 28
	sex  string = "女"
)

func main() {

	//  :=  声明变量并且定义变量并进行赋值,自动推断类型
	i := 1
	fmt.Println("i=", i)

	// 定义多组变量
	var name, age, sex = "张三", 10, "男"
	name = name + "_ 李四 "
	fmt.Println("name=", name, "age=", age, "sex=", sex)

	var num int = 111111

	fmt.Println("num : ", num)
	// 拼接字符串
	var aa string = name + "_1"
	fmt.Println("拼接字符串 : ", aa)

	/*
		类型 : int , uint, rune, byte
	*/
	var intDesc int = 1
	fmt.Println("intDesc : ", intDesc)
	// uint  >= 0
	var uintDesc uint = 0
	fmt.Println("uintDesc : ", uintDesc)

	// Printf 制表符, %T
	fmt.Printf("uintDesc  %T ", uintDesc)

	// 在程序里查看某个变量的张红字节和数据类型
	// unsafe 包
	fmt.Printf("uinDesc 的类型 %T,  uintDesc 占用字节长度 %d ", uintDesc, unsafe.Sizeof(uintDesc))

	//  浮点数 float32 单精度,  float64 双精度
	var price float32 = 1.22222222
	fmt.Printf("price: %v\n", price)

	// 丢失精度测试  goodsPrice:  123.0009 toDoodsPrice 123.000901
	// 说明: 如果保存精度比较高的数,应该使用较高精度的类型(float64)
	var fPrice = 10.21
	fmt.Printf("\n默认浮点类型  %T", fPrice)
	var goodsPrice float32 = 123.000901
	var toDoodsPrice float64 = 123.000901
	fmt.Println(" \n丢失精度测试  goodsPrice: ", goodsPrice, "toDoodsPrice", toDoodsPrice)

	// 字符串是不可变, string  str[0]
	var stringDesc string = "hello"
	// stringDesc[0] = "1"
	fmt.Println("stringDesc: ", stringDesc)

	// 拼接 string 字符串, 需要将 + 留在上面
	var strLen string = "hello" + "word" +
		"hello2" + "word"
	fmt.Println("strLen: ", strLen)

	// 浮点数基础数据类型,默认 0
	var folat32Desc float32
	fmt.Println("folat32Desc: ", folat32Desc)

	/*
		基本数据类型默认值(零值)
			int      0
			float    0
			string   ""
			bool     false
	*/

	// =========================================================
	// ---------------------------------------------------------
	//  golang 数据类型转换 (go中数据类型只能显示转换,不可以自动类型)
	// ---------------------------------------------------------
	// =========================================================
	/*
	 表达式 T(v) 将值 v 转换为类型 T
	 T 就是数据类型, 比如 int32, int64 float32 等等
	 v 就是需要转换的变量
	 注意: 数据类型转换, 只是将 v的值转换成了 T 类型, 但是 v 值类型是不变的,值可能会进行溢出
	*/
	var intDescTr int = 100
	var floatDescTr float32 = float32(intDescTr)

	fmt.Println("数据类型转换 int -> float : ", floatDescTr)

	// =========================================================
	// ---------------------------------------------------------
	//  golang  基本数据类型互转换string 类型
	// ---------------------------------------------------------
	// =========================================================
	/*
		方式1: fmt.Sprintf("%参数", 表达式) 会返回转换后的字符串
	*/
	// 定义数据类型时, 强制要求位数后, go 不可以自动类型转换,需要强制转换并且可能存在溢出处理问题
	stu1()

	// 基本数据类型转换 sring
	baseDataToString()

	// string 转换 基本数据类型
	stringToBaseData()

	// =========================================================
	// ---------------------------------------------------------
	//  golang  指针 类型
	// ---------------------------------------------------------
	// =========================================================
	/*
		总结:
			1. 值类型, 都有对应的指针类型, 形式为 "数据类型", 比如 int 的指针就是 *int, float32 对应的 float32
		 	2. 值类型包括 : 基本数据类型,int 系列, float 系列, boo, string, 数组和结构体 struct
	*/
	// 测试指针
	var a int = 1
	fmt.Println(" a 指针: ", &a) // 0xc000012138
	var ptr *int = &a
	// 指针也存在指针地址
	fmt.Println(" ptr 指针: ", &ptr) // 0xc000006030
	ptrFunc(ptr)
	fmt.Println("修改指针 ptr: ", *ptr) // 1

	// 方式二
	var ptrNum int = 10
	var ptrFlag = &ptrNum
	*ptrFlag = 10
	fmt.Println("ptrFlag: ", ptrNum)

	// =========================================================
	// ---------------------------------------------------------
	/*
		注意 :
			1. 保持 package 的名字和目录名字保持一致, 采用有意义的包名
			2. 变量名, 函数名, 常量名 均需要采用驼峰法
			3. 如果变量, 函数, 常量首字母大写,则可以被其他包名访问, 如果首字母小写则只能在本包进行访问
			   简单理解: 首字母是大写的则是公有的, 小写字母的是私有的

	*/
	// ---------------------------------------------------------
	// =========================================================

	// =========================================================
	// ---------------------------------------------------------
	/*
	  条件判断, switch, for 条件使用
	*/
	// ---------------------------------------------------------
	// =========================================================
	//  条件判断, switch, for 条件使用
	ifAndForAndSwitch()

	// ---------------------------------------------------------
	// =========================================================

	// =========================================================
	// ---------------------------------------------------------
	/*
		   包的基本概念:
		   		1. Go 的每一个文件, 都是属于一个包的, 也就是说 go 是以包的形式来管理文件, 或者项目目录结构的
				2. Go 语言的包借助了目录树的组织形式, 一般包的名称就是其源文件的目录的名称, 虽然 Go 语言没有严格要求
					包必须和其所在的目录包同名, 但还是建议包名和所在目录同名
				3. 包名一般都是小写, 使用一个简短有意义的名称
				4. 包名一般要和所在目录同名, 也可以不同, 包名中不能包含 - 等特殊字符
				5. 包名一般使用域名作为目录名称, 这样就能保证包名的唯一性, 必须 GitHub 项目的包一般会存放在
					GOPATH/src/github.com/userName/projectName 目录下
				6.(重点) 包名为 main 的是应用程序的入口, 编译不包含源文件是不会的到可执行文件的
				7. 一个文件下所有源文件只能属于同一个包, 同理, 属于同一个包的源文件不能存放在多个目录下

			包的作用:
				1. 区分相同名字的函数,变量等标识符
				2. 当程序文件过多时, 可以很好的管理项目
				3. 控制函数, 变量等访问范围,即作用域

			包初始化:
				1. 包初始化程序从 main 函数引用的包开始，逐级查找包的引用，直到找到没有引用其他包的包，最终生成一个包引用的有向无环图。
				2. Go 编译器会将有向无环图转换为一棵树，然后从树的叶子节点开始逐层向上对包进行初始化。
				3. 单个包的初始化过程如上图所示，先初始化常量，然后是全局变量，最后执行包的 init 函数。

			包的注意事项:
				1. 在 import 包时, 路径从 $GOPATH的 src下开始, 不用带 src, 编译器会自动从 src下开始引入
				2. 在调用函数或者变量等标识符是, 需要使用包进行访问(packageName.FuncName(), 函数名或变量名必须大写 )
				3. 包名称如果很长, 可以进行起别名,但注意, 起别名的包名,无法在使用原包名
						import (
							utils_tool  "go_pack/server/utils"
						)
						在使用 utils包内变量及函数时, 还是用 utils_tool.XXX
				4. 在同一个包中无法定义同一个变量和函数


			函数注意事项和细节讨论
				1. 函数的形参列表可以是多个, 返回值列表也可以是多个, 可以使用 _ 忽略接收的返回值
				2. 形参列表和返回值雷暴的数据类型可以是值类型和引用类型
				3. 函数的名字遵守标识符规范, 首字母不能是数字, 首字母大写是该函数是 public 可以本包文件或者其他包文件使用, 首字母小写是 private 只能半文件使用
				4. 函数中的变量是局部的, 函数外不生效
				5. 基本数据类型和<<数组>>默认是值传递, 即进行值拷贝, 在函数内修改, 不会影响到原来的值(可以使用 指针进行修改 & * )
				6. 如果希望函数内的向量修改函数外的变量的值, 可以传入变量的地址 &, 函数内以指针的方式进行操作变量, 从效果上看类似于 引用
				7. Go 函数不支持重载
				8. 在 Go 语言中, 函数也是一种数据类型, 可以赋给一个变量, 则该变量就是一个函数类型的变量, 可以通过该变量对函数的的调用


	*/
	// ---------------------------------------------------------
	// =========================================================

	// =========================================================
	// ---------------------------------------------------------
	/*
		函数注意事项和细节讨论
			1. 函数的形参列表可以是多个, 返回值列表也可以是多个, 可以使用 _ 忽略接收的返回值
			2. 形参列表和返回值雷暴的数据类型可以是值类型和引用类型
			3. 函数的名字遵守标识符规范, 首字母不能是数字, 首字母大写是该函数是 public 可以本包文件或者其他包文件使用, 首字母小写是 private 只能半文件使用
			4. 函数中的变量是局部的, 函数外不生效
			5. 基本数据类型和<<数组>>默认是值传递, 即进行值拷贝, 在函数内修改, 不会影响到原来的值(可以使用 指针进行修改 & * )
			6. 如果希望函数内的向量修改函数外的变量的值, 可以传入变量的地址 &, 函数内以指针的方式进行操作变量, 从效果上看类似于 引用
			7. Go 函数不支持重载
			8. 在 Go 语言中, 函数也是一种数据类型, 可以赋给一个变量, 则该变量就是一个函数类型的变量, 可以通过该变量对函数的的调用
			9. 函数可以作为形参并且进行调用
			10. 为了支持简化数据类型定义, Go 支持自定义数据类型
						基本语法: type 自定义数据类型名 数据类型
						案例1 : type goodInt int // 相当于和 int 起个别名, 这时候 goodInt 就可以等价于 int 来使用了
						案例2 : type goodFunc func(int, int) int  这时 goodFunc 就等价于一个函数类型 func(int, int) int
			11. 支持对函数返回值命名
			13. Go 支持可变参数
	*/
	// ---------------------------------------------------------
	// =========================================================

	// 将 函数赋值给一个变量进行调用
	var funcFlag func(int, int) int = operFuncGetSum
	// funcFlag := operFuncGetSum
	result := funcFlag(10, 1)
	fmt.Println(" 测试将方法赋值给予一个变量: ", result)

	// 自定义类型
	type goodInt int
	var goodIntFlag goodInt = 10
	fmt.Println(" 自定义变量 goodIntFlag: ", goodIntFlag)

	// 将函数作为一个形参赋值(自定义数据类型)
	operFuncGetSumVarHandler(operFuncGetSumVar, 10, 9)

	// 支持函数返回值命名
	result2, goodResult := operReturnResultName(1, 2)
	fmt.Println("result2: ", result2, "goodResult: ", goodResult)

	// =========================================================
	// ---------------------------------------------------------
	/*
		 闭包概念 :
			闭包是指能够读取函数内部变量的函数，一般来说只有函数内部的子函数才能读取该函数的局部变量，GO语言就是这种方式。
			个人理解闭包 被操作对象在特性环境中定义, 而这个特定环境则是一个 函数, 当该函数被 外部函数操作时, 函数内的变量或其他对象 需要被操作时,
			在函数内部定义一个子函数(匿名函数), 而函数需要返回定义的子函数对象(返回函数) 并且 子函数在内部操作该函数环境内的对象或其他对象,
			这样就形成了闭包,说白了就是 函数内定义了一个子函数, 子函数操作该函数内变量及其他对象,且函数返回的是子函数对象(返回函数), 这样就延长了包含子函数的函数对象生命周期

	*/
	// ---------------------------------------------------------
	// =========================================================
	// 闭包
	operClosure()

}

// 题目 1 定义数据类型时, 强制要求位数后, go 不可以自动类型转换,需要强制转换并且可能存在溢出处理问题
func stu1() {

	// 定义数据类型时, 强制要求位数后, go 不可以自动类型转换,需要强制转换并且可能存在溢出处理问题
	var i1 int8 = 1
	var i2 int32 = 20
	// var n1 int32 = i1 + 20
	// var n2 int64 = i2 + 40
	var n1 int32 = int32(i1) + 20
	var n2 int64 = int64(i2) + 40

	fmt.Println("题目一测试 数据类型转换 n1 :  ", n1)
	fmt.Println("题目一测试 数据类型转换 n2 :  ", n2)

}

// 基本数据类型转换 sring
func baseDataToString() {

	var num1 int = 99
	var num2 float32 = 12.2354
	var b bool = true
	var myChar byte = 'b'
	var str string

	// %d 十进制输入
	str = fmt.Sprintf("%d", num1)
	fmt.Println(" 基本数据类型 int 类型 转换 string ", str)
	// %f 浮点数输入
	str = fmt.Sprintf("%f", num2)
	fmt.Println(" 基本数据类型 float32 类型 转换 string ", str)
	// %t 单词true或false 输入
	str = fmt.Sprintf("%t", b)
	fmt.Println(" 基本数据类型  bool 类型 转换 string ", str)
	// %c 该值对应的unicode码值
	str = fmt.Sprintf("%c", myChar)
	fmt.Println(" 基本数据类型 byte 类型 转换 string ", str)

	var numTo1 int = 99
	var numTo2 float32 = 12.2354
	var bTo bool = true
	var strTo string

	// 将 numTo1 转换 string 类型并且使用 10进制形式输出
	// func FormatInt(i int64, base int) string
	// 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
	strTo = strconv.FormatInt(int64(numTo1), 10)
	fmt.Println("strTo: ", strTo)
	// 方便, 只能适用于 int 类型
	str = strconv.Itoa(num1)
	fmt.Println("strTo Itoa : ", strTo)

	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	// bitSize表示f的来源类型（32：float32、64：float64），会据此进行舍入。
	// fmt表示格式 f 是普通输出
	// prec控制精度
	// 测试 FormatFloat 方法 第三个参数是 Float转换 string 保留精度问题,数字超过原小数位可能会存在丢失精度问题
	strTo = strconv.FormatFloat(float64(numTo2), 'f', 5, 64)
	fmt.Println("bTo:", strTo)

	strTo = strconv.FormatBool(bTo)
	fmt.Println("bTo: ", strTo)

}

// string 转换 基本数据类型
func stringToBaseData() {

	var str string = "11"

	// string 转换 int 类型
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// s : string源数据, base : 转换为base进制数, bitSize : 转换为 int位数(int8, int32, int64)
	i, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(" string 转换 基本数据类型 int64  i: ", i)

	// string 转换 float64 类型
	//func ParseFloat(s string, bitSize int) (f float64, err error)
	// s : string 源数据,  base : 转换为base进制数
	f, _ := strconv.ParseFloat(str, 10)
	fmt.Println(" string 转换 基本数据类型 float64  f:  ", f)

	// string 转换 bool 类型
	//func ParseBool(str string) (value bool, err error)
	// str string 源数据
	b, err := strconv.ParseBool(str)
	fmt.Println(" string 转换 基本数据类型  bool b:  ", b, "err", err)

}

// 修改指针
func ptrFunc(i *int) {
	var aaa int = 100
	fmt.Println(" aaa 指针: ", &aaa) // 0xc0000aa110
	i = &aaa
	// 指针不可改, 指针的值可以修改 &i 调用指针, *i调用指针的值
	// *i = 100
	fmt.Println(" i 指针: ", &i) // 0xc000006038

}

//  条件判断, switch, for 条件使用
func ifAndForAndSwitch() {
	// if 条件分支
	ifFlag := 1 > 2
	if ifFlag {
		fmt.Println("if 条件成功 ....")
	} else {
		fmt.Println("if 条件失败 ....")
	}

	// switch
	var switchFlag int = 3
	switch switchFlag {
	case 2:
		fmt.Println(" switch case switchFlag 2")
	case 1:
		fmt.Println(" switch case switchFlag 1")
	default:
		fmt.Println(" switch case switchFlag 未能命中目标 ")

	}

	// for 循环
	for i := 0; i < 3; i++ {
		fmt.Println(" for  one i: ", i)
	}

	var forFlag int = 0
	for ; forFlag < 5; forFlag++ {
		fmt.Println("for two forFlag: ", forFlag)
	}

	// wihle true
	// for ;; {}
	// for true {
	// 	fmt.Println("for true th i: ", i)
	// }

	// for if 练习
	for i := 0; i < 10; i++ {
		if i < 5 {
			fmt.Println("for if  continue   i: ", i)
			continue
		}
		if i == 7 {
			fmt.Println("for if  break   i: ", i)
			break
		}
		fmt.Println("for if  i: ", i)
	}
}

//  将 函数赋值给一个变量进行调用
func operFuncGetSum(num int, flag int) int {

	fmt.Println(" num: ", num, "flag: ", flag)
	return num + flag
}

// 自定义数据类型
type goodFunc func(int, int) int

// 将参数作为一个形参进行计算
func operFuncGetSumVarHandler(operFuncGetSumVar goodFunc, num int, flag int) int {
	result := operFuncGetSumVar(num, flag)
	fmt.Println(" 将参数作为一个形参进行计算,包含自定义数据类型 result: ", result)
	return result
}

// 声明一个用于测试将方法作为参数的函数
func operFuncGetSumVar(num int, flag int) int {

	fmt.Println(" operFuncGetSumVar: ", num, "operFuncGetSumVar: ", flag)
	return num + flag
}

// 支持函数返回值命名
func operReturnResultName(num int, flag int) (result int, goodResult int) {
	result = num + 1
	goodResult = flag + 1
	return
}

// 闭包
func operClosure() {
	fmt.Println("闭包 开始 ... ")
	// 闭包关联对象
	funcFlag := closureHandler()
	funcFlag()
	funcFlag()
	funcFlag()
	fmt.Println("闭包 结束 ... ")
}

// 闭包环境
func closureHandler() func() {
	// 该 print 只打印一次
	fmt.Println("闭包 开始操作初始化 ...  ")
	oper := 10
	return func() {
		oper++
		fmt.Println("闭包 操作 oper ++ 结果: ", oper)
	}

}
