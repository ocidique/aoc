import time
from pathlib import Path

INPUT_FILE = Path("day01_input.txt")


def get_file_inputs(input_text: str) -> list:
    return list(input_text.splitlines())

def main():
    input_text = INPUT_FILE.read_text()
    calories = get_file_inputs(input_text)

    sums = []
    elf = 0
    for calorie in calories:
        if calorie:
            elf += int(calorie)
        else:
            sums.append(elf)
            elf = 0

    sums.sort()

    print("Most calories:", sums[-1])
    print("Top3 total calories:", sum(sums[-3:]))

if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")
