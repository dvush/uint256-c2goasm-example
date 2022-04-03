#include <cstdint>

void U256Add(uint64_t* a, uint64_t* b, uint64_t* res) {
  uint64_t carry_in = 0, carry_out = 0;
  // using clang builtin for __builtin_addcll
  res[0] = __builtin_addcll(a[0], b[0], carry_in, &carry_out);
  carry_in = carry_out;
  res[1] = __builtin_addcll(a[1], b[1], carry_in, &carry_out);
  carry_in = carry_out;
  res[2] = __builtin_addcll(a[2], b[2], carry_in, &carry_out);
  carry_in = carry_out;
  res[3] = __builtin_addcll(a[3], b[3], carry_in, &carry_out);
}


