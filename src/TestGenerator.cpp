#include "Generator.hpp"
#include "Util.hpp"
#include <mpi.h>
#include <unistd.h>
int main(int argc, char **argv)
{
    Util::Clear();

    cout << "<<<==========GENERATION TEST START=========>>>" << endl;

    double start_time, end_time;

    MPI_Init(NULL, NULL);
    int world_rank;
    MPI_Comm_rank(MPI_COMM_WORLD, &world_rank);
    MPI_Comm_size(MPI_COMM_WORLD, &Generator::WORLD_SIZE);

    Generator *generator = new Generator(argc, argv);

    start_time = MPI_Wtime();

    Generator::SCATTER_SEND_BUFFER = NULL;
    Generator::CreateSendData();
    assert(Generator::SCATTER_SEND_BUFFER != NULL);

    int scatter_recv_buffer;
    MPI_Barrier(MPI_COMM_WORLD);

    MPI_Scatter(Generator::SCATTER_SEND_BUFFER, 1, MPI_INT, &scatter_recv_buffer, 1, MPI_INT, 0, MPI_COMM_WORLD);
    srand((world_rank + 1) * time(0));

    char *words = (char *)malloc(sizeof(char) * scatter_recv_buffer * Generator::MAX_STR_LEN);
    Generator::CreateWords(words, scatter_recv_buffer);

    if (world_rank == 0)
        Generator::ALL_WORDS = (char *)malloc(sizeof(char) * Generator::DATASET_SIZE * Generator::MAX_STR_LEN);

    MPI_Gather(words, scatter_recv_buffer * Generator::MAX_STR_LEN, MPI_CHAR, Generator::ALL_WORDS, scatter_recv_buffer * Generator::MAX_STR_LEN, MPI_CHAR, 0, MPI_COMM_WORLD);
    MPI_Barrier(MPI_COMM_WORLD);

    if (world_rank == 0)
    {
        end_time = MPI_Wtime();
        Generator::WriteWords();
        Generator::WorkingTime(start_time, end_time);
        cout << "<<<===========GENERATION TEST END==========>>>" << endl;
    }

    MPI_Finalize();
    free(words);
}