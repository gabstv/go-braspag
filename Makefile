
all:
	echo "PHONY"
build:
	./compile_template.sh authorize
	./compile_template.sh capturecc
	./compile_template.sh query_getboletodata
