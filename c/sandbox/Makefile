compile:
	gcc main.c -o main

run: compile
	./main ; rm -f main

compile-to-asm:
	gcc -S main.c -o main.s

compile-asm:
	gcc main.s -o main
