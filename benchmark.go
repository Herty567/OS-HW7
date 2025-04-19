package main

import (
    "fmt"
    "math/rand"
    "time"
    "osproject/raids"
    "osproject/shared"
)

func BenchmarkRAID(r raids.RAID, totalMB int) {
    blockCount := (totalMB * 1024 * 1024) / shared.BlockSize
    data := make([]byte, shared.BlockSize)

    fmt.Printf("Writing %d blocks (%dMB)...\n", blockCount, totalMB)
    startWrite := time.Now()
    for i := 0; i < blockCount; i++ {
        rand.Read(data)
        if err := r.Write(i, data); err != nil {
            fmt.Println("Write error:", err)
            return
        }
    }
    durationWrite := time.Since(startWrite)

    fmt.Println("Reading data back...")
    startRead := time.Now()
    for i := 0; i < blockCount; i++ {
        _, err := r.Read(i)
        if err != nil {
            fmt.Println("Read error:", err)
            return
        }
    }
    durationRead := time.Since(startRead)

    fmt.Printf("Write time: %.2fs (%.2f ms/block)\n", durationWrite.Seconds(), durationWrite.Seconds()*1000/float64(blockCount))
    fmt.Printf("Read  time: %.2fs (%.2f ms/block)\n", durationRead.Seconds(), durationRead.Seconds()*1000/float64(blockCount))
}
