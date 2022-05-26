#ifndef UTIL
#define UTIL

#ifndef __INCLUDES__
    #include <iostream>
    #include <fstream>
    #include <cstring>
    #include <time.h>
    #include <assert.h>
    using namespace std;
#endif

class Util {
  public:
    void static Clear();
    void static PrintArray(int *arr, int size);
    string static GetNowTime();
};

#endif