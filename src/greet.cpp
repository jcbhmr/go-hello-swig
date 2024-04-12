#include <iostream>
#include <string>
#include "greet.hpp"
#include "math.hpp"

std::string greeting(const std::string& name) {
    int n = 5;
    int res = fib(n);
    return "Hello " + name + "! Did you know that fib(" + std::to_string(n) + ") is " + std::to_string(res) + "?";
}

void greet(const std::string& name) {
    std::cout << greeting(name) << "\n";
}
