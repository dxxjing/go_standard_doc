package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main () {


	//bufioRead()
	bufioWrite()


}
//缓冲的读取
func bufioRead() {
	file, err := os.Open("testfile.txt")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()

	//创建缓冲读
	reader := bufio.NewReader(file)

	var content []byte
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		//str, err := reader.ReadString('\n')
		//fmt.Println(buf[:n], err)
		if err == io.EOF{
			break
		}
		content = append(content, buf[:n]...)
	}
	fmt.Println(string(content))
}

//写缓冲
func bufioWrite () {
	file, err := os.OpenFile("writefile.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 6)
	if err != nil {
		fmt.Println(err)
		return
	}
	//创建默认大小的缓冲
	writer := bufio.NewWriter(file)
	//写入缓冲
	writer.WriteString("hello jdx\n")
	writer.WriteString("hello jdx\n")
	fmt.Println(writer.Buffered())
	//将缓冲区内容写入文件
	writer.Flush()
}
