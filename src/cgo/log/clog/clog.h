#ifndef __CLOG_H
#define __CLOG_H

#ifndef __cplusplus

//void LogDebug(const char* logname, int level, const char* file, int line, const char* content, ...) __attribute__ ((format (printf, 5, 6)));
void LogDebug(const char* logname, int level, const char* file, int line, const char* content);

#endif

#endif

