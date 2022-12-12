import time
from pathlib import Path
from functools import reduce
import numpy as np

INPUT_FILE = Path("day08_input.txt")


def get_file_inputs(input_text: str) -> list:
    return list(input_text.splitlines())


def create_grid(lines):
    grid = np.array([[int(tree) for tree in line.strip()] for line in lines])
    return grid


def main():
    input_text = INPUT_FILE.read_text()
    lines = get_file_inputs(input_text)
    GRID = create_grid(lines)
    MASK = np.zeros_like(GRID)
    dists = [np.zeros_like(GRID) for _ in range(4)]

    for k, dist in enumerate(dists):
        grid, mask = np.rot90(GRID, k=k), np.rot90(MASK, k=k)
        for row, (g, m, d) in enumerate(zip(grid, mask, dist)):
            current = 0
            for col, h in enumerate(g):
                if col == 0 or h > current:
                    m[col] = 1
                    current = h
                current_height = h
                counter = 0
                for i in g[col + 1 :]:
                    counter += 1
                    if i >= current_height:
                        break
                d[col] = counter

    print("Part1: ", np.sum(mask))

    dists = [np.rot90(d, k=-k) for k, d in enumerate(dists)]
    print("Part2: ", np.max(reduce(np.multiply, dists)))


if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")
