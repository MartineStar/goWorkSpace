package main
import (
	"fmt"
	"os"
	"bufio"
	"io"
)
//直接return，会返回指定返回类型的默认值
func mytest() (abc int64,err error){
	return 	//abc: 0 err: <nil>
}

func CopyFile(dstFileName string,srcFileName string) (written int64,err error){
	srcFile,err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n",err)
		return
	}
	//开启文件成功后及时关闭
	defer srcFile.Close()

	reader := bufio.NewReader(srcFile)
	//打开dstFileName,如果不存在则创建
	dstFile,err :=os.OpenFile(dstFileName,os.O_WRONLY | os.O_CREATE,0666)
	if err != nil {
		fmt.Println("open file err=",err)
		return 
	}
	defer dstFile.Close()
	writer := bufio.NewWriter(dstFile)

	return io.Copy(writer,reader)
}
func main(){
	// 将e:/11.jpg写入到f:/rabbit.jpg
	srcFile := "f:/office20f13.zip"
	dstFile := "e:/office2014.zip"
	written,err := CopyFile(dstFile,srcFile)
	if err == nil{
		fmt.Println("拷贝完成")
	}else {
		fmt.Println("拷贝失败：",err)
	}

	w,e := mytest()
	fmt.Println("w:",w,"e:",e)
}