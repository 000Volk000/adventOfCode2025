#include <fstream>
#include <iostream>
#include <string>

int main(int argc, char *argv[]) {
  std::ifstream fich;
  fich.open("day1.txt");
  if (fich.is_open()) {
    std::cout << "File opened successfully\n\n";
  } else {
    std::cout << "Couldn't open file\n";
    return -1;
  }

  std::string instruction;
  int cont = 0, lock = 50;
  while (std::getline(fich, instruction)) {
    if (instruction[0] == 'R') {
      instruction.erase(0, 1);
      int sum = std::stoi(instruction);
      lock = (lock + sum) % 100;
    } else if (instruction[0] == 'L') {
      instruction.erase(0, 1);
      int subs = std::stoi(instruction);
      lock = (lock - subs) % 100;
    } else {
      std::cout << "Wrong input format\n";
      return -1;
    }

    if (lock == 0) {
      cont++;
    }
  }

  fich.close();
  std::cout << "The number 0 appears " << cont << " times\n";
  return 0;
}
