#!/bin/bash

set -e

echo "===== DISK LOAD BALANCING TEST ====="

if [ ! -f users/USER0 ]; then
  echo "FAIL: users/USER0 not found"
  exit 1
fi

if [ ! -f users/USER0_LOAD_BALANCE ]; then
  echo "FAIL: users/USER0_LOAD_BALANCE not found"
  exit 1
fi

rm -f PRINTER* load_balance_output.log

cp users/USER0 users/USER0.backup
cp users/USER0_LOAD_BALANCE users/USER0

cleanup() {
  mv users/USER0.backup users/USER0
}
trap cleanup EXIT

go run . 1 2 1 | tee load_balance_output.log

echo
echo "===== VALIDATION ====="

if grep -q "Chose disk" load_balance_output.log; then
  echo "PASS: load-balancing disk choice logged"
else
  echo "FAIL: no disk choice log found"
  exit 1
fi

if grep -q "has reusable space" load_balance_output.log; then
  echo "PASS: reclaimed space was preferred"
else
  echo "WARN: no reclaimed-space preference observed in this run"
fi

echo
echo "===== SAVE PLACEMENT LOGS ====="
grep "Starting save on disk" load_balance_output.log || true

echo
echo "===== PRINTER OUTPUT ====="
if [ -f PRINTER0 ]; then
  cat PRINTER0
else
  echo "PRINTER0 not created"
fi

echo
echo "===== LOAD BALANCING TEST COMPLETE ====="