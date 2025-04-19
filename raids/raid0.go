package raids

import (
    "fmt"
    "osproject/shared"
)

// RAID0 uses striping without redundancy for maximum speed and capacity
type RAID0 struct {
    Disks []*shared.Disk
}

// NewRAID0 returns a RAID0 object with a slice of disks
func NewRAID0(disks []*shared.Disk) *RAID0 {
    return &RAID0{Disks: disks}
}

// Write stripes the data across all disks
func (r *RAID0) Write(blockNum int, data []byte) error {
    if len(data) != shared.BlockSize {
        return fmt.Errorf("data must be exactly one block")
    }
    diskIndex := blockNum % len(r.Disks)
    stripeIndex := blockNum / len(r.Disks)
    return r.Disks[diskIndex].WriteBlock(stripeIndex, data)
}

// Read retrieves striped data from the appropriate disk
func (r *RAID0) Read(blockNum int) ([]byte, error) {
    diskIndex := blockNum % len(r.Disks)
    stripeIndex := blockNum / len(r.Disks)
    return r.Disks[diskIndex].ReadBlock(stripeIndex)
}
