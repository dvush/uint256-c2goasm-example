C2GOASM_FLAGS := -masm=intel -mno-red-zone -mstackrealign -mllvm -inline-threshold=1000 -fno-asynchronous-unwind-tables -fno-exceptions -fno-rtti  
OTHER_FLAGS := -O3 -march=native -std=c++11

all: goasmu256 bench_cpp

bench_cpp: cpp/uint256.cpp cpp/bench.cpp
	clang++ ${C2GOASM_FLAGS} ${OTHER_FLAGS} cpp/uint256.cpp cpp/bench.cpp -o bench_cpp

cpp/uint256.s: cpp/uint256.cpp
	clang++ ${C2GOASM_FLAGS} ${OTHER_FLAGS} -S cpp/uint256.cpp -o cpp/uint256.s

Uint256.s: cpp/uint256.s
	c2goasm -a cpp/uint256.s Uint256.s

goasmu256: *.go Uint256.s
	go build .
