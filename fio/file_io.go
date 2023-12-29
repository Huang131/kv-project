package fio

import "os"

// FileIo 标准文件IO
type FileIo struct {
	fd *os.File // 系统文件描述符
}

func NewFileIoManger(fileName string) (*FileIo, error) {
	fd, err := os.OpenFile(fileName,
		os.O_CREATE|os.O_RDWR|os.O_APPEND,
		DataFilePerm,
	)
	if err != nil {
		return nil, err
	}

	return &FileIo{fd: fd}, nil
}

func (fio *FileIo) Read(bytes []byte, offset int64) (int, error) {
	return fio.fd.ReadAt(bytes, offset)
}

func (fio *FileIo) Write(bytes []byte) (int, error) {
	return fio.fd.Write(bytes)
}

func (fio *FileIo) Sync() error {
	return fio.fd.Sync()
}

func (fio *FileIo) Close() error {
	return fio.fd.Close()
}
