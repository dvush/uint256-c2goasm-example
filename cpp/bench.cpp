#include <cstdint>

void U256Add(uint64_t* a, uint64_t* b, uint64_t* res);

#include <iostream>
#include <chrono>
using namespace std::chrono;

int main() {
  uint64_t a[4] = {
    18443354236259924217ull, 17506605513766861049ull, 18364758544493064720ull, 8198552921648689ull};
  uint64_t b[4] = {
    18443354236259924217ull, 17506605513766861049ull, 18364758544493064720ull, 8198552921648689ull};

  int N = 100000000;
  auto start = high_resolution_clock::now();

  for (int i = 0; i < N; i++) {
    U256Add(a, b, b);  
  }
  
  auto stop = high_resolution_clock::now();
  auto duration = duration_cast<nanoseconds>(stop - start);

  std::cout << "c++ benchmark" << std::endl;

  std::cout << (double)duration.count()/(double) N << std::endl;

  std::cout << b[0] << " ";
  std::cout << b[1] << " ";
  std::cout << b[2] << " ";
  std::cout << b[3] << " " << std::endl;
  

  return 0;
}
