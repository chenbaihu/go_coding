#include "fun.h"

#include <stdio.h>

char* MySecret() {
      return "my secret not tell anybody";
}


void HelloWorld(const char* str) {
    printf("HelloWorld: %s\n", str);
}
