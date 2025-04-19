package raids

import (
    "osproject/shared"
)

// RAID1 mirrors data across all disks for redundancy
type RAID1 struct {
    Disks []*shared.Disk
}

// NewRAID1 returns a RAID1 instance
func NewRAID1(disks []*shared.Disk) *RAID1 {
    return &RAID1{Disks: disks}
}

// Write writes the same block to every disk (mirroring)
func (r *RAID1) Write(blockNum int, data []byte) error {
    for _, disk := range r.Disks {
        if err := disk.WriteBlock(blockNum, data); err != nil {
            return err
        }
    }
    return nil
}

// Read from the first disk (in practice, could use any)
func (r *RAID1) Read(blockNum int) ([]byte, error) {
    return r.Disks[0].ReadBlock(blockNum)
}
