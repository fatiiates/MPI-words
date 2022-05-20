all: build run

build:
	clear
#	mpic++ -I ./include -o ./lib/GenerateData.o -c GenerateData.cpp

	mpic++ -I ./include -o ./bin/gen_data GenerateData.cpp

run:
	clear
	mpirun -n $(MPI_WORLD_SIZE) ./bin/gen_data $(DATASET_SIZE)