#include <iostream>
#include <string>
#include "greet.hpp"

std::string greeting(const std::string& name) {
    return "Hello " + name + "!";
}

void greet(const std::string& name) {
    std::cout << greeting(name) << "\n";
}
