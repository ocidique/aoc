import time
from pathlib import Path
import string

INPUT_FILE = Path("day03_input.txt")


def get_file_inputs(input_text: str) -> list:
    return list(input_text.splitlines())

def part1(rucksacks, alphabet):
    sum = 0
    for content in rucksacks:
        first, second = content[:len(content)//2], content[len(content)//2:]
        common = list(set(first) & set(second))[0]
        sum += 1 + alphabet.index(common)

    print("Part1: ", sum)
        
def part2(rucksacks, alphabet):
    sum = 0
    for i in range(0, len(rucksacks), 3):
        common = list(set(rucksacks[i]) & set(rucksacks[i+1]) & set(rucksacks[i+2]))[0]
        sum += 1 + alphabet.index(common)
    
    print("Part2: ", sum)

def main():
    input_text = INPUT_FILE.read_text()
    rucksacks = get_file_inputs(input_text)
    alphabet = list(string.ascii_letters)

    part1(rucksacks, alphabet)
    part2(rucksacks, alphabet)

if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")
