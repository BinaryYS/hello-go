package fileop

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("文件打开失败", err)
		return ""
	}

	defer file.Close()
	//bufio缓冲区读取数据
	reader := bufio.NewReader(file)
	cnt := 0

	var content string
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("总行数: %d \n", cnt)
			return content
		}

		if err != nil {
			fmt.Println("文件读取失败", err)
			return content
		}

		//fmt.Print(str
		content += str
		cnt++
	}
}

func Read() string {
	//
	file, err := os.Open("/Users/ssg/Desktop/trade.jmx")
	if err != nil {
		fmt.Println("文件打开失败", err)
		return ""
	}

	//main函数结束前执行文件资源释放
	defer file.Close()
	fmt.Println("文件打开成功")
	var result [1024]byte

	n, err := file.Read(result[:])
	//文件读取完毕
	if err == io.EOF {
		fmt.Println("读取文件失败", err)
		return ""
	}
	//文件读取中出现异常
	if err != nil {
		fmt.Println("读取文件失败", err)
		return ""
	}

	fmt.Printf("字节数%d\n", n)

	return string(string(result[:]))
}
