
all:
	echo "PHONY"
build:
	if [[ -f ./authorize.xml.go ]]; then rm authorize.xml.go; fi
	bin2go -p braspag -l 16 -a templates/authorize.xml
	mv templates/authorize.xml.go ./authorize.xml.go
