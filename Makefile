all: build run

build:
	clear
	mpic++ -I ./include -o ./lib/Generator.o -c ./src/Generator.cpp
	mpic++ -I ./include -o ./lib/Util.o -c ./src/Util.cpp

	mpic++ -I ./include -o ./bin/gen_data ./lib/Generator.o ./lib/Util.o ./src/TestGenerator.cpp

MAX_STR_LEN ?= 10
MIN_STR_LEN ?= 2
run:
	clear
	mpirun --oversubscribe -np $(MPI_WORLD_SIZE) ./bin/gen_data $(DATASET_SIZE) $(MAX_STR_LEN) $(MIN_STR_LEN) 