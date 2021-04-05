package strategy

import "fmt"

type Decompression interface {
	Decompress(file string)
}

type RarDecompression struct {
}

func (rd *RarDecompression) Decompress(file string) {
	fmt.Println("用rar的方式解压", file)
}

type TarGzDecompression struct {
}

func (td *TarGzDecompression) Decompress(file string) {
	fmt.Println("用tar.gz的方式解压", file)
}

func Run() {
	file := "要解压的文件"

	rd := &RarDecompression{}
	rd.Decompress(file)

	td := &TarGzDecompression{}
	td.Decompress(file)

}
