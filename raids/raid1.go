package raids

import (
    "osproject/shared"
)

type RAID1 struct {
    Disks []*shared.Disk
}

func NewRAID1(disks []*shared.Disk) *RAID1 {
    return &RAID1{Disks: disks}
}

func (r *RAID1) Write(blockNum int, data []byte) error {
    for _, disk := range r.Disks {
        if err := disk.WriteBlock(blockNum, data); err != nil {
            return err
        }
    }
    return nil
}

func (r *RAID1) Read(blockNum int) ([]byte, error) {
    return r.Disks[0].ReadBlock(blockNum)
}
