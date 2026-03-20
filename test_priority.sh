#!/bin/bash

set -e

run_priority_case() {
  local test_file="$1"
  local label="$2"
  local expected_order="$3"

  echo "===== $label ====="

  if [ ! -f users/USER0 ]; then
    echo "FAIL: users/USER0 not found"
    exit 1
  fi

  if [ ! -f "$test_file" ]; then
    echo "FAIL: $test_file not found"
    exit 1
  fi

  rm -f PRINTER* priority_demo_output.log
  cp users/USER0 users/USER0.backup
  cp "$test_file" users/USER0

  cleanup_case() {
    mv users/USER0.backup users/USER0
  }
  trap cleanup_case RETURN

  go run . 1 1 1 | tee priority_demo_output.log

  echo
  echo "===== SCHEDULER ORDER ($label) ====="
  grep "Scheduling print job:" priority_demo_output.log || true

  actual_order=$(grep "Scheduling print job:" priority_demo_output.log | awk '{print $4}' | paste -sd "," -)

  echo
  echo "Expected order: $expected_order"
  echo "Actual order:   $actual_order"

  if [ "$actual_order" = "$expected_order" ]; then
    echo "PASS: scheduler order matched expected priority behavior"
  else
    echo "FAIL: scheduler order did not match expected order"
    exit 1
  fi

  echo
  echo "===== PRINTER OUTPUT ($label) ====="
  if [ -f PRINTER0 ]; then
    cat PRINTER0
  else
    echo "PRINTER0 was not created"
    exit 1
  fi

  echo
}

echo "===== BASELINE TEST ====="
rm -f PRINTER*
go run . 1 1 1 > /dev/null

if [ ! -f PRINTER0 ]; then
  echo "FAIL: baseline PRINTER0 was not created"
  exit 1
fi

lines=$(wc -l < PRINTER0)
if [ "$lines" -eq 100 ]; then
  echo "PASS: baseline PRINTER0 has 100 lines"
else
  echo "FAIL: expected 100 baseline lines, got $lines"
  exit 1
fi

echo
echo "This script runs two priority scheduling demos:"
echo "1. Simple urgent vs normal"
echo "2. Multiple queued urgent + normal jobs"
echo

run_priority_case \
  "users/USER0_PRIORITY" \
  "PRIORITY TEST (SIMPLE)" \
  "urgent_A,normal_B"

run_priority_case \
  "users/USER0_PRIORITY_MULTI" \
  "PRIORITY TEST (MULTI)" \
  "urgent_B,urgent_D,normal_A,normal_C"

echo "===== ALL PRIORITY TESTS PASSED ====="