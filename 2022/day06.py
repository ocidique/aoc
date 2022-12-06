import time
from pathlib import Path

INPUT_FILE = Path("day06_input.txt")


def get_marker(characters, sequence_length):
    for idx, _ in enumerate(characters):
        sequence = set(characters[idx : idx + sequence_length])
        if len(sequence) == sequence_length:
            return idx + sequence_length


def main():
    characters = INPUT_FILE.read_text()

    print("Part1: ", get_marker(characters, 4))
    print("Part2: ", get_marker(characters, 14))


if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")
