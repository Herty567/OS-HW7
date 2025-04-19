package shared

import (
    "fmt"
    "os"
)

const BlockSize = 4096

type Disk struct {
    File *os.File
    ID   int
}

func OpenDisk(filename string, id int) (*Disk, error) {
    f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0666)
    if err != nil {
        return nil, err
    }

    // Pre-size the file to 100MB (25600 * 4096 bytes)
    const totalBlocks = 25600
    err = f.Truncate(int64(totalBlocks * BlockSize))
    if err != nil {
        return nil, err
    }

    return &Disk{File: f, ID: id}, nil
}

func (d *Disk) WriteBlock(blockNum int, data []byte) error {
    if len(data) != BlockSize {
        return fmt.Errorf("invalid block size: expected %d, got %d", BlockSize, len(data))
    }
    offset := int64(blockNum * BlockSize)
    _, err := d.File.WriteAt(data, offset)
    if err != nil {
        return err
    }
    return d.File.Sync()
}

func (d *Disk) ReadBlock(blockNum int) ([]byte, error) {
    offset := int64(blockNum * BlockSize)
    buf := make([]byte, BlockSize)
    _, err := d.File.ReadAt(buf, offset)
    return buf, err
}
