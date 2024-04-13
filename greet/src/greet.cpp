#include <iostream>
#include <string>
#include <fmt/format.h>
#include "greet.hpp"
#include "math.hpp"

std::string greeting(const std::string& name) {
    int n = 5;
    int res = fib(n);
    return fmt::format("Hello {}! Did you know that fib({}) is {}?", name, n, res);
}

void greet(const std::string& name) {
    fmt::print("{}\n", greeting(name));
}
