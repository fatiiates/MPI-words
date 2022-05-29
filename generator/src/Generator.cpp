#include "Generator.hpp"
#include "Util.hpp"

int Generator::MAX_STR_LEN = 10;
int Generator::MIN_STR_LEN = 2;
int Generator::DATASET_SIZE = 50000;
char *Generator::LETTERS = (char *)malloc(Generator::LETTER_COUNT * sizeof (char));
int Generator::WORLD_SIZE = 1;
int *Generator::SCATTER_SEND_BUFFER = nullptr;
int *Generator::GATHER_RECV_COUNTS = nullptr;
int *Generator::GATHER_DISPLACEMENTS = nullptr;
char *Generator::ALL_WORDS = nullptr;


void Generator::validate(int argc, char **argv){

    if (Generator::WORLD_SIZE > 100)
        throw invalid_argument("WORLD_SIZE can be 100 at max");

    if (Generator::WORLD_SIZE < 1)
        throw invalid_argument("WORLD_SIZE can't be lowest from 1");
        
    if (argc < 2)
        throw invalid_argument("CLI arguments must have 2 arguments at least");

    Generator::DATASET_SIZE = atoi(argv[1]);

    if (strlen(argv[1]) > 7 || Generator::DATASET_SIZE > 1000000)
        throw invalid_argument("DATASET_SIZE can be 1M at max");

    if (Generator::DATASET_SIZE <= 0)
        throw invalid_argument("DATASET_SIZE can't be zero or negative");

    if (Generator::WORLD_SIZE > Generator::DATASET_SIZE)
        throw invalid_argument("DATASET_SIZE can't be lower from WORLD_SIZE");

    Generator::MAX_STR_LEN = 10;
    if (argc > 2)
       Generator::MAX_STR_LEN = atoi(argv[2]);

    if (Generator::MAX_STR_LEN > 100)
        throw invalid_argument("MAX_STR_LEN can be 100 at max");

    if (Generator::MAX_STR_LEN < 2)
        throw invalid_argument("MAX_STR_LEN can't be lowest from 2");

    Generator::MIN_STR_LEN = 2;
    if (argc > 3)
       Generator::MIN_STR_LEN = atoi(argv[3]);

    if (Generator::MIN_STR_LEN > 100)
        throw invalid_argument("MIN_STR_LEN can be 100 at max");

    if (Generator::MIN_STR_LEN < 1)
        throw invalid_argument("MIN_STR_LEN can't be lowest from 1");

    if (Generator::MIN_STR_LEN > Generator::MAX_STR_LEN)
        throw invalid_argument("MIN_STR_LEN has to be lowest or equal to MAX_STR_LEN");


}

void Generator::createLetters(){
    for (int i = 0; i < Generator::LETTER_COUNT; i++)
        Generator::LETTERS[i] = 'A' + i;
}



Generator::Generator(int argc, char **argv){
    Generator::validate(argc, argv);
    Generator::createLetters();
}

Generator::Generator(){}

void Generator::CreateSendData(){
    int total = Generator::DATASET_SIZE;
    int element_per_process = Generator::DATASET_SIZE / Generator::WORLD_SIZE; 
    Generator::SCATTER_SEND_BUFFER = (int *)malloc(Generator::WORLD_SIZE * sizeof(int));
    for (int i = 0; i < Generator::WORLD_SIZE; i++){
        Generator::SCATTER_SEND_BUFFER[i] = total > element_per_process ? element_per_process: total;
        total -= Generator::SCATTER_SEND_BUFFER[i];
    }
    if (total > 0)
        for (int i = 0; i < total; i++)
            Generator::SCATTER_SEND_BUFFER[i] += 1;
}

void Generator::CreateRecvCountsAndDisplacements(){
    if (Generator::SCATTER_SEND_BUFFER == NULL)
        Generator::CreateSendData();
    
    Generator::GATHER_DISPLACEMENTS = (int *)malloc(Generator::WORLD_SIZE * sizeof(int));
    Generator::GATHER_RECV_COUNTS = (int *)malloc(Generator::WORLD_SIZE * sizeof(int));
    
    Generator::GATHER_DISPLACEMENTS[0] = 0;
    Generator::GATHER_RECV_COUNTS[0] = Generator::SCATTER_SEND_BUFFER[0] * Generator::MAX_STR_LEN;
    int total_chars = 0;
    for (int i = 1; i < Generator::WORLD_SIZE; i++){
        total_chars += Generator::SCATTER_SEND_BUFFER[i - 1] * Generator::MAX_STR_LEN;
        Generator::GATHER_DISPLACEMENTS[i] = total_chars;//i * Generator::SCATTER_SEND_BUFFER[i - 1] * Generator::MAX_STR_LEN;
        Generator::GATHER_RECV_COUNTS[i] = Generator::SCATTER_SEND_BUFFER[i] * Generator::MAX_STR_LEN;
    }
}

void Generator::CreateWord(char **word){
    int len = rand() % (Generator::MAX_STR_LEN - Generator::MIN_STR_LEN + 1) + Generator::MIN_STR_LEN;
    *word = (char *)malloc((Generator::MAX_STR_LEN) * sizeof(char));
    int i = 0;
    for (; i < len; i++)
    {
        int r = rand() % LETTER_COUNT;
        (*word)[i] = Generator::LETTERS[r];
    }
    if (len != MAX_STR_LEN)
        (*word)[i] = '\0';
}

void Generator::CreateWords(char *arr, int buff_size){
    for (int i = 0; i < buff_size; i++){
        char *word;
        Generator::CreateWord(&word);
        for (int j = 0; j < Generator::MAX_STR_LEN; j++)
        {
            arr[i*Generator::MAX_STR_LEN + j] = word[j];
            if (word[j] == '\0')
                break;
        }
    }
}

void Generator::WorkingTime(double start_time, double end_time){
    printf("\nThat generation took %f seconds.\n", end_time - start_time);
}

void Generator::PrintWord(char *word){
    for (int i = 0; i < Generator::MAX_STR_LEN; i++)
    {
        if (word[i] == '\0')
            break;
        cout << word[i];
    }
    
    cout << endl;
}

void Generator::WriteWords(){
    ofstream f;
    string filename = Util::GetNowTime();
    f.open ("./generator/results/" + filename + ".txt");
    for (int i = 0; i < Generator::DATASET_SIZE; i++){
        for (int j = 0; j < Generator::MAX_STR_LEN; j++)
        {
            if ((Generator::ALL_WORDS + i*Generator::MAX_STR_LEN)[j] == '\0')
                break;
            f << (Generator::ALL_WORDS + i*Generator::MAX_STR_LEN)[j];
        }
        f << " ";
    }
    f.close();
    cout << endl << "Words written to file " << filename << ".txt";
}

void Generator::PrintWords(){
    for (int i = 0; i < Generator::DATASET_SIZE; i++){
        Generator::PrintWord((Generator::ALL_WORDS + i*Generator::MAX_STR_LEN));
    }
}

void Generator::PrintWords(char *arr, int buff_size){
    for (int i = 0; i < buff_size; i++)
        Generator::PrintWord((arr + i*Generator::MAX_STR_LEN));
}

Generator::~Generator(){
    if(Generator::SCATTER_SEND_BUFFER != NULL)
        free(Generator::SCATTER_SEND_BUFFER);

    if(Generator::ALL_WORDS != NULL)
        free(Generator::ALL_WORDS);

    if(Generator::LETTERS != NULL)
        free(Generator::LETTERS);
}
