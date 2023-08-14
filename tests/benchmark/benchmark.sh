#!/bin/bash

# RUN GO BENCHMARK FILES
echo "=== GO BENCHMARK ==="
go test
echo "=== END GO BENCHMARK ==="
echo ""

# RUN PY BENCHMARK FILES
echo "=== PY BENCHMARK ==="
py_files=`ls ./*.py`
for p in $py_files
do
  python3 $p
done
echo "=== END PY BENCHMARK ==="
