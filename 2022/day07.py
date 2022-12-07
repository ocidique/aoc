import time
from pathlib import Path
from collections import defaultdict

INPUT_FILE = Path("day07_input.txt")


def get_file_inputs(input_text: str) -> list:
    return list(input_text.splitlines())


def traverse_directories(inputs: list) -> dict:
    directories = defaultdict(int)
    pwd = []
    for input in inputs:
        match input.split():
            case [_, _, "/"]:
                pwd = []
            case [_, _, ".."]:
                pwd.pop()
            case [_, _, x]:
                pwd.append(x)
            case [a, _] if a.isdigit():
                for i in range(len(pwd) + 1):
                    path = "/" + "/".join(pwd[:i])
                    directories[path] += int(a)

    return directories


def get_part1(directories: dict) -> int:
    MAX_DIR_SIZE = 100000
    return sum(filter(lambda v: v <= MAX_DIR_SIZE, directories.values()))


def get_part2(directories: dict) -> int:
    TOTAL_DISK_SPACE = 70000000
    REQUIRED_DISK_SPACE = 30000000
    unused_space = TOTAL_DISK_SPACE - directories["/"]
    required_space = REQUIRED_DISK_SPACE - unused_space

    return min(filter(lambda v: v >= required_space, directories.values()))


def main():
    input_text = INPUT_FILE.read_text()
    inputs = get_file_inputs(input_text)

    print("Part1: ", get_part1(traverse_directories(inputs)))
    print("Part2: ", get_part2(traverse_directories(inputs)))


if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")
