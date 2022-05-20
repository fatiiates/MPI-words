#include <mpi.h>
#include <stdio.h>
#include<vector>

using namespace std;

#define LETTER_COUNT 26

void validateArgs(int argc, char **argv)
{
    if (argc < 2)
        throw invalid_argument("CLI arguments must have 2 arguments at least");

    int DATASET_SIZE = atoi(argv[1]);

    if (strlen(argv[1]) > 7 || DATASET_SIZE > 1000000)
        throw invalid_argument("MPI_DATASET_SIZE can be 1M at max");

    if (DATASET_SIZE <= 0)
        throw invalid_argument("MPI_DATASET_SIZE can't be zero or negative");
}

void createLetters(char *letters)
{
    for (int i = 0; i < LETTER_COUNT; i++)
        letters[i] = 'A' + i;
}

string createWord(char *letters)
{
    int len = rand() % 10 + 3;
    string word = "";
    for (int i = 0; i < len; i++)
    {
        int r = rand() % LETTER_COUNT;
        word += letters[r];
    }

    return word;
}

void createSendData(int *arr, int DATASET_SIZE, int world_size)
{
    for (int i = 0; i < world_size; i++)
        arr[i] = DATASET_SIZE / world_size;
}

void workingTime(double start_time, double end_time)
{
    printf("That took %f seconds\n", end_time - start_time);
}

void Clear()
{
#if defined _WIN32
    system("cls");
#elif defined (__LINUX__) || defined(__gnu_linux__) || defined(__linux__)
    system("clear");
#elif defined (__APPLE__)
    system("clear");
#endif
}

int main(int argc, char **argv)
{
    Clear();
    validateArgs(argc, argv);

    char letters[LETTER_COUNT];
    createLetters(letters);

    double start_time, end_time;

    MPI_Init(NULL, NULL);
    int world_rank;
    MPI_Comm_rank(MPI_COMM_WORLD, &world_rank);
    int world_size;
    MPI_Comm_size(MPI_COMM_WORLD, &world_size);
    int DATASET_SIZE = atoi(argv[1]);
    int ELEMENTS_PER_PROC = DATASET_SIZE / world_size;

    start_time = MPI_Wtime();

    int scatter_send_buf[world_size];
    createSendData(scatter_send_buf, DATASET_SIZE, world_size);

    int scatter_recv_buf[1];

    MPI_Barrier(MPI_COMM_WORLD);

    MPI_Scatter(scatter_send_buf, ELEMENTS_PER_PROC, MPI_INT, scatter_recv_buf, ELEMENTS_PER_PROC, MPI_INT, 0, MPI_COMM_WORLD);

    string words[scatter_recv_buf[0]];
    for (int i = 0; i < scatter_recv_buf[0]; i++)
        words[i] = createWord(letters);

    vector<string> *all_words = NULL;

    if (world_rank == 0)
        all_words = new vector<string>(scatter_recv_buf[0] * world_size);

    MPI_Gather(&words, scatter_recv_buf[0], MPI_CHAR, all_words, scatter_recv_buf[0], MPI_CHAR, 0, MPI_COMM_WORLD);
    if (world_rank == 0)
    {
        end_time = MPI_Wtime();
        workingTime(start_time, end_time);

        for (int i = 0; i < DATASET_SIZE; i++)
        {
            cout << "d => " << DATASET_SIZE << "testest" << i<<endl;
            cout << all_words << endl;
            cout << "yeni test";
        }
    }

    // // tüm veriler burada toparlanarak bir dosyaya kaydedilecek

    MPI_Barrier(MPI_COMM_WORLD);

    MPI_Finalize();

    return 0;
}