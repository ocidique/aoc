import time
from pathlib import Path
import re

INPUT_FILE = Path("day04_input.txt")


def get_file_inputs(input_text: str) -> list:
    return list(input_text.splitlines())

def main():
    input_text = INPUT_FILE.read_text()
    assignments = get_file_inputs(input_text)

    complete_overlap_sum = 0
    overlap_at_all_sum = 0

    for assignment in assignments:
        elf1_start, elf1_end, elf2_start, elf2_end = map(int, re.split(r"[,-]", assignment))

        if (((elf1_start <= elf2_start <= elf1_end) and (elf1_start <= elf2_end <= elf1_end)) or
            ((elf2_start <= elf1_start <= elf2_end) and (elf2_start <= elf1_end <= elf2_end))):
            complete_overlap_sum += 1

        if ((elf1_start <= (elf2_start or elf2_end) <= elf1_end) or
            (elf2_start <= (elf1_start or elf1_end) <= elf2_end)):
            overlap_at_all_sum += 1

    print("Part1: ", complete_overlap_sum)
    print("Part2: ", overlap_at_all_sum)

if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")
