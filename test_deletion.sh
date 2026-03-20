#!/bin/bash

set -e

echo "===== DELETE + RECLAMATION TEST ====="

if [ ! -f users/USER0 ]; then
  echo "FAIL: users/USER0 not found"
  exit 1
fi

if [ ! -f users/USER0_DELETION ]; then
  echo "FAIL: users/USER0_DELETION not found"
  exit 1
fi

rm -f PRINTER* deletion_output.log

cp users/USER0 users/USER0.backup
cp users/USER0_DELETION users/USER0

cleanup() {
  mv users/USER0.backup users/USER0
}
trap cleanup EXIT

go run . 1 1 1 | tee deletion_output.log

echo
echo "===== VALIDATION ====="

if grep -q "Deleted file metadata for: temp_A" deletion_output.log; then
  echo "PASS: delete command executed"
else
  echo "FAIL: delete command not detected"
  exit 1
fi

if grep -q "Freed disk space on disk 0 start 0 length 2" deletion_output.log; then
  echo "PASS: disk space was reclaimed"
else
  echo "FAIL: reclaimed disk space not detected"
  exit 1
fi

if grep -q "Starting save on disk 0 at sector 0" deletion_output.log; then
  echo "PASS: reused freed disk space"
else
  echo "FAIL: did not reuse freed sector 0"
  exit 1
fi

if grep -q "File not found" deletion_output.log; then
  echo "PASS: deleted file cannot be printed"
else
  echo "FAIL: deleted file still accessible"
  exit 1
fi

echo
echo "===== PRINTER OUTPUT ====="
if [ -f PRINTER0 ]; then
  cat PRINTER0
else
  echo "PRINTER0 not created"
fi

echo
echo "===== DELETE TEST COMPLETE ====="