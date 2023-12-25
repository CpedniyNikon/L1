package tasks

import "fmt"

type USB interface {
	ConnectWithUsbCable()
}

type MemoryCard struct {
}

func (m MemoryCard) Insert() {
	fmt.Println("inserted")
}

func (m MemoryCard) CopyData() {
	fmt.Println("copied")
}

type CardReader struct {
	MemoryCard MemoryCard
}

func NewCardReader(memoryCard MemoryCard) *CardReader {
	return &CardReader{MemoryCard: memoryCard}
}

func (r CardReader) ConnectWithUsbCable() {
	r.MemoryCard.Insert()
	r.MemoryCard.CopyData()
}

func Task21() {
	CardReader := NewCardReader(MemoryCard{})
	CardReader.ConnectWithUsbCable()
}
