package strategy

import "fmt"

type DecompressionStrategy interface {
	Decompress(file string)
}

type RarDecompressionStrategy struct {
}

func (rd *RarDecompressionStrategy) Decompress(file string) {
	fmt.Println("用rar的方式解压", file)
}

type TarGzDecompressionStrategy struct {
}

func (td *TarGzDecompressionStrategy) Decompress(file string) {
	fmt.Println("用tar.gz的方式解压", file)
}

type Decompression struct {
	decompressionStrategy DecompressionStrategy
}

func NewDecompression(decompressionStrategy DecompressionStrategy) *Decompression {
	return &Decompression{decompressionStrategy}
}

func (d *Decompression) Decompression(file string) {
	d.decompressionStrategy.Decompress(file)
}

func Run() {
	file := "要解压的文件"

	NewDecompression(&RarDecompressionStrategy{}).Decompression(file)

	NewDecompression(&TarGzDecompressionStrategy{}).Decompression(file)

}
