import time
from pathlib import Path

INPUT_FILE = Path("day02_input.txt")


def get_file_inputs(input_text: str) -> list:
    return iter(input_text.splitlines())


def main():
    input_text = INPUT_FILE.read_text()
    instructions = get_file_inputs(input_text)

    position_part1 = {
        "x": 0,
        "z": 0,
    }

    position_part2 = {
        "x": 0,
        "z": 0,
        "aim": 0,
    }

    for instruction in instructions:
        instruction = instruction.split()
        
        direction = instruction[0]
        distance = int(instruction[1])


        if direction == "forward":
            position_part1["x"] += distance

            position_part2["x"] += distance
            position_part2["z"] += distance * position_part2["aim"]  

        elif direction == "down":
            position_part1["z"] += distance

            position_part2["aim"] += distance

        elif direction == "up":
            position_part1["z"] -= distance

            position_part2["aim"] -= distance


    print("Part1: ", position_part1["x"] * position_part1["z"])
    print("Part2: ", position_part2["x"] * position_part2["z"])


if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")