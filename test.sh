#!/bin/bash -e

goyacc calc.y > /dev/null
go build -o llvmcalcc ./cmd

function assert()
{
  input=$1
  expect=$2

  real=`./llvmcalcc "${input}" && lli ./a.ll; echo $?`
  if [[ $real != $expect ]]; then
    echo "[$input]: expect ($expect), but got ($real)"
    exit 1
  fi
}

assert '1 + 2' '3'
assert '3 - 1' '2'
assert '2 * 3' '6'
assert '6 / 2' '3'

assert '3 + 2 - 1' '4'
assert '3 + 1 - 2' '2'
assert '3 * 1 - 2' '1'
assert '5 - 3 * 1' '2'

assert '5 - 10 / 2' '0'
assert '10 - 4 / 2' '8'
assert '(10 - 4) / 2' '3'
assert '(10 - 4) * (10 - 9) / (1 + 2)' '2'

assert '1.5 + 1.5' '3'
assert '1.5 + 1.2' '2'
