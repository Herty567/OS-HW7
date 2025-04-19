package raids

import (
    "fmt"
    "osproject/shared"
)

type RAID0 struct {
    Disks []*shared.Disk
}

func NewRAID0(disks []*shared.Disk) *RAID0 {
    return &RAID0{Disks: disks}
}

func (r *RAID0) Write(blockNum int, data []byte) error {
    if len(data) != shared.BlockSize {
        return fmt.Errorf("data must be exactly one block")
    }
    diskIndex := blockNum % len(r.Disks)
    stripeIndex := blockNum / len(r.Disks)
    return r.Disks[diskIndex].WriteBlock(stripeIndex, data)
}

func (r *RAID0) Read(blockNum int) ([]byte, error) {
    diskIndex := blockNum % len(r.Disks)
    stripeIndex := blockNum / len(r.Disks)
    return r.Disks[diskIndex].ReadBlock(stripeIndex)
}
