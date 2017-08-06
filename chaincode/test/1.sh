#!/bin/bash
set -x
echo $1
cmd="peer chaincode invoke -n charity -c '$1' -C myc" 
$cmd
