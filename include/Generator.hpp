#ifndef GENERATOR
#define GENERATOR

#ifndef __INCLUDES__
    #include <iostream>
    #include <fstream>
    #include <cstring>
    #include <time.h>
    #include <assert.h>
    using namespace std;
#endif

class Generator {
  private:
    void static validate(int argc, char **argv);
    void static createLetters();
    string static getNowTime();

  public:
    static int const LETTER_COUNT = 26;
    static int MAX_STR_LEN;
    static int MIN_STR_LEN;
    static int DATASET_SIZE;
    static char *LETTERS;
    static int WORLD_SIZE;
    static int *SCATTER_SEND_BUFFER;
    static char *ALL_WORDS;
    static int ELEMENTS_PER_PROC;
    Generator(int argc, char **argv);
    Generator();
    void static CreateSendData();
    void static CreateWord(char **word);
    void static CreateWords(char *arr, int buff_size);
    void static WorkingTime(double start_time, double end_time);
    void static PrintWord(char *word);
    void static WriteWords();
    void static PrintWords();
    void static PrintWords(char *arr, int buff_size);
    ~Generator();
};

#endif