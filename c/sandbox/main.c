#include <stdio.h>
#include <limits.h>

// Assumes little endian
void printBits(size_t const size, void const * const ptr)
{
    unsigned char *b = (unsigned char*) ptr;
    unsigned char byte;
    int i, j;

    for (i = size-1; i >= 0; i--) {
        for (j = 7; j >= 0; j--) {
            byte = (b[i] >> j) & 1;
            printf("%u", byte);
        }
    }
    puts("");
}

int main() {
   printf("Integers\n");

   int cI = INT_MAX;
   printf("%d\n", cI);
   printBits(sizeof(cI), &cI);

   cI += 1;
   printf("%d\n", cI);
   printBits(sizeof(cI), &cI);

   printf("Floats\n");

   float cF = -0.0;
   printf("%f\n", cF);
   printBits(sizeof(cF), &cF);

   return 0;
}
