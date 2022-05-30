all: builder runner

builder: build_generator build_counter
runner: run_generator run_counter

BR_generator: build_generator run_generator

BR_counter: build_counter run_counter

build_counter:
	cd counter && go build -o ./bin
	clear

WORLD_SIZE ?= 1
run_counter:
	cd counter && ./bin/counter $(WORLD_SIZE) $(GENERATED_FILE_PATH)

build_generator:
	clear
	mpic++ -I generator/include -o generator/lib/Generator.o -c generator/src/Generator.cpp
	mpic++ -I generator/include -o generator/lib/Util.o -c generator/src/Util.cpp

	mpic++ -I generator/include -o generator/bin/gen_data generator/lib/Generator.o generator/lib/Util.o generator/src/TestGenerator.cpp

DATASET_SIZE ?= 50000
MAX_STR_LEN ?= 10
MIN_STR_LEN ?= 2
run_generator:
	clear
	mpirun --oversubscribe -np $(WORLD_SIZE) generator/bin/gen_data $(DATASET_SIZE) $(MAX_STR_LEN) $(MIN_STR_LEN) 