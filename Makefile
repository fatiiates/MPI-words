all: build_generator run_generator

build_generator:
	clear
	mpic++ -I generator/include -o generator/lib/Generator.o -c generator/src/Generator.cpp
	mpic++ -I generator/include -o generator/lib/Util.o -c generator/src/Util.cpp

	mpic++ -I generator/include -o generator/bin/gen_data generator/lib/Generator.o generator/lib/Util.o generator/src/TestGenerator.cpp

MAX_STR_LEN ?= 10
MIN_STR_LEN ?= 2
run_generator:
	clear
	mpirun --oversubscribe -np $(WORLD_SIZE) generator/bin/gen_data $(DATASET_SIZE) $(MAX_STR_LEN) $(MIN_STR_LEN) 