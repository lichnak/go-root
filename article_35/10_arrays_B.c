#include <stdio.h>
#include <stdint.h>

int main(void) {
	int32_t a[10] = {0};
        int i;

	puts("Pole pred upravou:");
        for (i=0; i<10; i++) {
            printf(" %d", a[i]);
        }
        puts("\n");

        for (i=0; i<10; i++) {
            a[i] = i * 2;
        }

	puts("Pole po uprave:");
        for (i=0; i<10; i++) {
            printf(" %d", a[i]);
        }
        return 0;
}