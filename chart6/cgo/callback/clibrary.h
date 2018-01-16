#ifndef CLIBRARY_H
#define CLIBRARY_H
typedef int (*callback_func)(int);
void do_c_func(callback_func);
#endif