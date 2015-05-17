#include "clog.h"

#include <stdio.h>

//void LogDebug(const char* logname, int level, const char* file, int line, const char* content, ...) {
void LogDebug(const char* logname, int level, const char* file, int line, const char* content) {
    fprintf(stdout, content);
    fprintf(stdout, "\n");
    fprintf(stdout, "logname=%s\tleve=%d\tfile=%s\tline=%d\tcontent=%s\n", logname, level, file, line, content);
}

#ifdef _DEBUG_

int main(int argc, char* argv[]) 
{
    LogDebug("clog.clog_test", 0, __FILE__, __LINE__, "md5\t123456\twin");
    return 0;
}

#endif
