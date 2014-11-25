#!/bin/bash
if [[ -f ./$1.xml.go ]]; then rm $1.xml.go; fi
bin2go -p braspag -l 16 -a templates/$1.xml
mv templates/$1.xml.go ./$1.xml.go