package raids

import (
    "fmt"
    "osproject/shared"
)

// RAID5 uses striping with distributed parity
type RAID5 struct {
    Disks []*shared.Disk
}

// NewRAID5 constructs a RAID5 setup with distributed parity
func NewRAID5(disks []*shared.Disk) *RAID5 {
    if len(disks) < 3 {
        panic("RAID5 requires at least 3 disks")
    }
    return &RAID5{Disks: disks}
}

// Write stores data with rotating parity
func (r *RAID5) Write(blockNum int, data []byte) error {
    if len(data) != shared.BlockSize {
        return fmt.Errorf("data must be exactly one block")
    }
    n := len(r.Disks)
    stripe := blockNum / (n - 1)
    offset := blockNum % (n - 1)

    parityDisk := stripe % n
    dataDisk := offset
    if dataDisk >= parityDisk {
        dataDisk++
    }

    if err := r.Disks[dataDisk].WriteBlock(stripe, data); err != nil {
        return err
    }

    // XOR all other blocks in the stripe to compute parity
    parity := make([]byte, shared.BlockSize)
    for i := 0; i < n; i++ {
        if i == parityDisk {
            continue
        }
        d, err := r.Disks[i].ReadBlock(stripe)
        if err != nil {
            return err
        }
        for j := 0; j < shared.BlockSize; j++ {
            parity[j] ^= d[j]
        }
    }

    return r.Disks[parityDisk].WriteBlock(stripe, parity)
}

// Read retrieves the correct block using RAID 5 striping logic
func (r *RAID5) Read(blockNum int) ([]byte, error) {
    n := len(r.Disks)
    stripe := blockNum / (n - 1)
    offset := blockNum % (n - 1)

    parityDisk := stripe % n
    dataDisk := offset
    if dataDisk >= parityDisk {
        dataDisk++
    }

    return r.Disks[dataDisk].ReadBlock(stripe)
}
