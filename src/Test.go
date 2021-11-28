package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func main() {

	//testPrint("")

	// 测试处理文件名称
	//var splitName  string= testSplit("1224file.txt",1)
	//fmt.Println("去掉开头指定长度后的文件名称",splitName)

	// 读取文件夹中全部文件名称
	var baseFilePath = "/Users/may/Desktop/ray/go/music/file/"

	var nameList []string = renameAllFiles(baseFilePath,1)
	fmt.Println(nameList)


	// 字符串拼接 参考 https://www.cnblogs.com/mambakb/p/10352138.html
	//var baseFilePath= "/Users/may/Desktop/ray/go/music/file/"

	//for i := 0; i <= 10; i++ {
	//	var byteBuffer bytes.Buffer
	//	byteBuffer.WriteString(baseFilePath)
	//	byteBuffer.WriteString(strconv.Itoa(i))
	//	byteBuffer.WriteString("file.txt")
	//	newName:=byteBuffer.String()
	//	fmt.Println(newName)
	//
	//	copyFile2(fullFilename,newName)
	//
	//}

}

func rename(source string, target string) bool {
	// 重命名文件
	file := source
	err1 := os.Rename(file, target)
	if err1 != nil {
		panic(err1)
		return false
	} else {
		return true
		println(`文件重命名成功`)
	}
	return true
}

func renameDir(source string, target string) bool {
	// 重命名文件夹
	folder := `./新建文件夹`
	err2 := os.Rename(folder, `重命名文件夹`)
	if err2 != nil {
		panic(err2)
	} else {
		println(`文件夹重命名成功`)
	}
	return  true
}

func renameAllFiles(dir string,removeLength int) []string {
	var nameList []string

	// 读取当前目录中的所有文件和子目录
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	// 获取文件，并输出它们的名字
	for _, file := range files {

		//需要先导入Strings包
		var build strings.Builder
		build.WriteString(dir)
		build.WriteString(testSplit(file.Name(), removeLength))
		var fullFileNameNew = build.String()
		var buildSource strings.Builder
		buildSource.WriteString(dir)
		buildSource.WriteString(file.Name())


		fmt.Println(fullFileNameNew)
		nameList = append(nameList, fullFileNameNew)
		os.Rename(buildSource.String(), fullFileNameNew)

	}

	return nameList
}

func testSplit(filename string, removeLength int) string {
	var len = len(filename)
	fmt.Print("原始名称=", filename, " 长度是=", len)
	//1224file.txt  --> file.txt

	var start int = removeLength
	var end int = len

	filename = filename[start:end]
	fmt.Println(" start=", start, " end=", end, " 处理后:", filename)

	return filename

}

func testPrintFile(name string) {
	fmt.Println("参数 name =", name)

	fullFilename := "/Users/may/Desktop/ray/go/music/file/file1.txt"
	fmt.Println("文件全名 =", fullFilename)
	var filenameWithSuffix string
	filenameWithSuffix = path.Base(fullFilename) //获取文件名带后缀
	fmt.Println("文件名称+后缀filenameWithSuffix =", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
	fmt.Println("文件后缀fileSuffix =", fileSuffix)

	var filenameOnly string
	filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix) //获取文件名
	fmt.Println("filenameOnly =", filenameOnly)
}

func copyFile2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1)
}
