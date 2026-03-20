# 141OS in Go (CS141 Extra Credit)

This project reimplements the 141OS simulation from Homework 9 in Go, using concurrency (goroutines, mutexes, condition variables) to model users, disks, printers, and resource management.

In addition to the base implementation, the system includes three extensions that add more realistic operating system behavior.

---

## How to Run

### Baseline System
```bash
go run . <numUsers> <numDisks> <numPrinters>

### Test Scripts
./test.sh              # baseline correctness
./test_priority.sh     # extension 1: print scheduling
./test_deletion.sh     # extension 2: deletion + reclamation
./test_load_balance.sh # extension 3: disk load balancing
```

## Extensions

1. Print Job Queue + Priority Scheduling
- All print requests are enqueued into a centralized queue
- Jobs are scheduled by priority if we find that the file prefix is "urgent_"

2. File Deletion + Disk Space Reclamation
- Added functionality for .delete <filename> command
- Removes file metadata from the directory
- Frees disk sectors and tracks them as reusable segments
- Future writes reuse freed space instead of always appending

3. Disk Load Balancing
- Save operations choose a disk based on:
    - Preferred: available reusable space
    - Fallback: current disk usage (nextFreeSector)
- Distributes files more evenly across disks
