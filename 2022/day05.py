import time
import re
from pathlib import Path

INPUT_FILE = Path("day05_input.txt")


def get_file_inputs(input_text: str) -> list:
    return list(input_text.splitlines())


def crate_mover_9000(stacks, instructions):
    for instruction in instructions:
        amount, stack_from, stack_to = map(int, re.findall(r"\d+", instruction))

        boxes = stacks[stack_from - 1][len(stacks[stack_from - 1]) - amount :]

        # add to
        stacks[stack_to - 1] = stacks[stack_to - 1] + list(reversed(boxes))

        # remove from
        stacks[stack_from - 1] = stacks[stack_from - 1][
            : len(stacks[stack_from - 1]) - amount
        ]

    part1 = []
    for stack in stacks:
        part1.append(stack[-1])

    print("Part1 top crates: ", part1)


def crate_mover_9001(stacks, instructions):
    for instruction in instructions:
        amount, stack_from, stack_to = map(int, re.findall(r"\d+", instruction))

        boxes = stacks[stack_from - 1][len(stacks[stack_from - 1]) - amount :]

        # add to
        stacks[stack_to - 1] = stacks[stack_to - 1] + boxes

        # remove from
        stacks[stack_from - 1] = stacks[stack_from - 1][
            : len(stacks[stack_from - 1]) - amount
        ]

    part2 = []
    for stack in stacks:
        part2.append(stack[-1])

    print("Part2 top crates: ", part2)


def parse_stacks(stack_raw_rows):
    stack_rows = []

    # Stack rows into arrays
    for stack_raw_row in stack_raw_rows:
        sr = re.findall("....?", stack_raw_row)
        stack_rows.append(sr)

    # Stack rows into column stacks using zip, and convert those tuples to array
    stacks_zip = list(zip(*stack_rows))
    stacks_columns = []
    for stack_tuple in stacks_zip:
        stacks_columns.append(list(stack_tuple))

    # Reverse stack columns order and clean the clutter
    stacks = []
    for stacks_column in stacks_columns:
        stacks_column.reverse()
        stack = []
        for s in stacks_column:
            if not s.isspace():
                stack.append(re.sub(r"\s+", "", s))  # remove trailing space

        stacks.append(stack)

    return stacks


def main():
    input_text = INPUT_FILE.read_text()
    inputs = get_file_inputs(input_text)

    crate_mover_9000(parse_stacks(inputs[:8]), inputs[10:])
    crate_mover_9001(parse_stacks(inputs[:8]), inputs[10:])


if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")
