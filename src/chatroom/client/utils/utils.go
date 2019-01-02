package utils

import (
	"fmt"
	"net"
	"chatroom/common/message"	
	"encoding/binary"
	"encoding/json"
	"errors"
)

//这里定义一个结构体，将所有方法关联到这个结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn net.Conn
	Buf [8096]byte  //传输时,使用缓冲
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送给一个长度给对方
	var pkgLen uint32 
	pkgLen = uint32(len(data))
	// var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)

	//发送长度
	_,err = this.Conn.Write(this.Buf[:4])
	if err != nil{
		fmt.Println("conn.Write(resMes's length) failed,error=",err)
		return
	}

	//发送data本身
	n, err := this.Conn.Write(data)	
	if n != int(pkgLen) || err !=nil {
		fmt.Println("conn.Write(resMes's content) failed,error=",err)
		return
	}
	return
}

func (this *Transfer) ReadPkg() (mes message.Message,err error) {

	//conn.Read 在conn没有关闭的情况下才会阻塞
	//  如果客户端关闭了conn,则不会阻塞
	_,err = this.Conn.Read(this.Buf[:4])	//err再返回值列表中已经默认定义，不需要:=
	if err != nil {
		// 自定义错误并返回
		// err = errors.New("read pkg header<data's length> error")
		return
	}

	//根据buf[:4]转成一个unit32类型
	pkgLen := binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen 读取消息内容,将读到的内容放到buf中,n表示实际读取到的字节数
	n ,err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		err = errors.New("read pkg body<data's content> error")
		return
	}

	//把buf[:pkgLen]反序列化成 message.Message类型,因为mes是结构体，需要传地址
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)  
	if err != nil{
		fmt.Println("json.Unmarshal failed,error=",err)
		return 
	}
	return
}