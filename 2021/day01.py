import time
from pathlib import Path
from typing import Iterator

INPUT_FILE = Path("day01_input.txt")


def get_file_inputs(input_text: str) -> list:
    return list(map(int, input_text.splitlines()))


def get_couple_increases(measurements: Iterator[int]) -> int:
    couples = zip(measurements[:-1], measurements[1:])
    return sum(couple[1] > couple[0] for couple in couples)


def get_triple_increases(measurements: Iterator[int]) -> int:
    triples = zip(measurements[:-1], measurements[1:], measurements[2:])

    sum_triples = list()

    for triples in triples:
        sum_triples.append(sum(triples))

    return get_couple_increases(sum_triples)


def main():
    input_text = INPUT_FILE.read_text()
    measurements = get_file_inputs(input_text)
    
    couple_increases = get_couple_increases(measurements)
    triple_increases = get_triple_increases(measurements)

    print("Part1: increases", couple_increases, "times")
    print("Part2: increases", triple_increases, "times")


if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")
