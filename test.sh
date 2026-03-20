#!/bin/bash

set -e

echo "===== TEST 1: Single User Baseline ====="
rm -f PRINTER*
go run . 1 1 1 > /dev/null

if [ ! -f PRINTER0 ]; then
  echo "FAIL: PRINTER0 was not created"
  exit 1
fi

lines=$(wc -l < PRINTER0)
if [ "$lines" -eq 100 ]; then
  echo "PASS: PRINTER0 has 100 lines"
else
  echo "FAIL: expected 100 lines in PRINTER0, got $lines"
  exit 1
fi

a0=$(grep -h "^A0 00000$" PRINTER0 | wc -l | tr -d ' ')
a9=$(grep -h "^A9 44444$" PRINTER0 | wc -l | tr -d ' ')

if [ "$a0" -eq 2 ] && [ "$a9" -eq 2 ]; then
  echo "PASS: A0 and A9 each printed twice"
else
  echo "FAIL: expected A0/A9 counts to be 2, got A0=$a0 A9=$a9"
  exit 1
fi

echo
echo "===== TEST 2: Two Users, One Printer ====="
rm -f PRINTER*
go run . 2 2 1 > /dev/null

if [ ! -f PRINTER0 ]; then
  echo "FAIL: PRINTER0 was not created"
  exit 1
fi

lines=$(wc -l < PRINTER0)
if [ "$lines" -eq 200 ]; then
  echo "PASS: PRINTER0 has 200 lines"
else
  echo "FAIL: expected 200 lines in PRINTER0, got $lines"
  exit 1
fi

a0=$(grep -h "^A0 00000$" PRINTER0 | wc -l | tr -d ' ')
b9=$(grep -h "^B9 44444$" PRINTER0 | wc -l | tr -d ' ')

if [ "$a0" -eq 2 ] && [ "$b9" -eq 2 ]; then
  echo "PASS: A0 and B9 each printed twice"
else
  echo "FAIL: expected A0/B9 counts to be 2, got A0=$a0 B9=$b9"
  exit 1
fi

echo
echo "===== TEST 3: Two Users, Two Printers ====="
rm -f PRINTER*
go run . 2 2 2 > /dev/null

printer_count=$(ls PRINTER* 2>/dev/null | wc -l | tr -d ' ')
if [ "$printer_count" -lt 1 ]; then
  echo "FAIL: no printer output files were created"
  exit 1
fi

total=$(cat PRINTER* | wc -l | tr -d ' ')
if [ "$total" -eq 200 ]; then
  echo "PASS: total lines across printers = 200"
else
  echo "FAIL: expected total of 200 lines across printers, got $total"
  exit 1
fi

echo "Printer file line counts:"
wc -l PRINTER*

echo
echo "===== TEST 4: Stress Test (4 Users, 2 Disks, 3 Printers) ====="
rm -f PRINTER*
go run . 4 2 3 > /dev/null

printer_count=$(ls PRINTER* 2>/dev/null | wc -l | tr -d ' ')
if [ "$printer_count" -lt 1 ]; then
  echo "FAIL: no printer output files were created"
  exit 1
fi

total=$(cat PRINTER* | wc -l | tr -d ' ')
if [ "$total" -eq 400 ]; then
  echo "PASS: total lines across printers = 400"
else
  echo "FAIL: expected total of 400 lines across printers, got $total"
  exit 1
fi

echo "Printer file line counts:"
wc -l PRINTER*

echo
echo "===== ALL TESTS PASSED ====="