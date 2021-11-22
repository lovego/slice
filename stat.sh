#!/bin/bash
shopt -s extglob
for i in !(*_test).go; do
    egrep -c '^func ' $i ${i/.go/_test.go};
done
