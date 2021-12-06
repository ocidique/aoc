import time
from pathlib import Path
from typing import List, Tuple

INPUT_FILE = Path("day03_input.txt")


def get_file_inputs(input_text: str) -> List[str]:
    return list(input_text.splitlines())


def binary_to_decimal(lines: List[str]) -> List[int]:
    return [int(line, base=2) for line in lines]


def get_gamma(bits: List[int], data: List[int]) -> int:
    return sum(bit for bit in bits if sum(d & bit for d in data) // bit >= len(data) / 2)


def get_epsilon(bits: List[int], data: List[int]) -> int:
    return sum(bit for bit in bits if sum(d & bit for d in data) // bit <= len(data) / 2) 


def get_part1(bits: List[int], data: List[int]) -> Tuple[int, int]:
    return (get_gamma(bits, data), get_epsilon(bits, data))


def filter_data_bitwise(bits: List[int], data: List[int], filter_by_most_common: bool = True) -> List[int]:
    filtered = [x for x in data]
    for bit in reversed(bits):
        ratio = sum(1 for num in filtered if num & bit) / len(filtered)
        wanted_bit_value = bit * int((ratio >= 0.5) == filter_by_most_common)
        filtered = [x for x in filtered if x & bit == wanted_bit_value]
        if len(filtered) == 1:
            break
    return filtered


def get_part2(bits: List[int], data: List[int]) -> Tuple[int, int]:
    filtered_by_most_common = filter_data_bitwise(bits, data)
    filtered_by_least_common = filter_data_bitwise(bits, data, filter_by_most_common = False)

    return (filtered_by_most_common[0], filtered_by_least_common[0])


def main():
    input_text = INPUT_FILE.read_text()
    lines = get_file_inputs(input_text)
    digit_count = len(lines[0].strip())
    data = binary_to_decimal(lines)
    bits = [2 ** n for n in range(digit_count)] # [1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048]

    (gamma, epsilon) = get_part1(bits, data)
    print("Part 1:", gamma * epsilon)

    (o, co2) = get_part2(bits, data)
    print("Part 2:", o * co2)


if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")