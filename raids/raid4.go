package raids

import (
    "fmt"
    "osproject/shared"
)

// RAID4 uses striping with a dedicated parity disk
type RAID4 struct {
    Disks []*shared.Disk
}

// NewRAID4 creates a new RAID4 array from the given disks
func NewRAID4(disks []*shared.Disk) *RAID4 {
    if len(disks) < 2 {
        panic("RAID4 requires at least 2 disks")
    }
    return &RAID4{Disks: disks}
}

// Write stores data on a data disk and updates parity on a separate disk
func (r *RAID4) Write(blockNum int, data []byte) error {
    if len(data) != shared.BlockSize {
        return fmt.Errorf("data must be exactly one block")
    }
    numDataDisks := len(r.Disks) - 1
    diskIndex := blockNum % numDataDisks
    stripeIndex := blockNum / numDataDisks

    err := r.Disks[diskIndex].WriteBlock(stripeIndex, data)
    if err != nil {
        return err
    }

    // Calculate parity by XORing all data disks
    parity := make([]byte, shared.BlockSize)
    for i := 0; i < numDataDisks; i++ {
        if i == diskIndex {
            for j := 0; j < shared.BlockSize; j++ {
                parity[j] ^= data[j]
            }
        } else {
            d, err := r.Disks[i].ReadBlock(stripeIndex)
            if err != nil {
                return err
            }
            for j := 0; j < shared.BlockSize; j++ {
                parity[j] ^= d[j]
            }
        }
    }

    return r.Disks[len(r.Disks)-1].WriteBlock(stripeIndex, parity)
}

// Read retrieves data from the appropriate data disk
func (r *RAID4) Read(blockNum int) ([]byte, error) {
    numDataDisks := len(r.Disks) - 1
    diskIndex := blockNum % numDataDisks
    stripeIndex := blockNum / numDataDisks
    return r.Disks[diskIndex].ReadBlock(stripeIndex)
}
