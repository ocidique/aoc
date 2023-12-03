import time
import numpy as np


def simulate_rope(input, length):
    knots = []

    for k in range(length):
        knots.append((0, 0))

    visited = set()
    visited.add(knots[-1])

    for motion in input:
        direction, steps = motion.split()
        for _ in range(int(steps)):
            match direction:
                case "U":
                    knots[0] = (knots[0][0], knots[0][1] + 1)
                case "D":
                    knots[0] = (knots[0][0], knots[0][1] - 1)
                case "R":
                    knots[0] = (knots[0][0] + 1, knots[0][1])
                case "L":
                    knots[0] = (knots[0][0] - 1, knots[0][1])
            for k in range(length - 1):
                diff_x = knots[k][0] - knots[k + 1][0]
                diff_y = knots[k][1] - knots[k + 1][1]
                if abs(diff_x) > 1 or abs(diff_y) > 1:
                    knots[k + 1] = (
                        knots[k + 1][0] + np.sign(diff_x),
                        knots[k + 1][1] + np.sign(diff_y),
                    )
                visited.add(knots[-1])
    return len(visited)


def main():
    with open("day09_input.txt", "r") as file:
        motions = file.read().strip().split("\n")
        part1 = simulate_rope(motions, 2)
        part2 = simulate_rope(motions, 10)
        print(f"Part1: {part1}")
        print(f"Part2: {part2}")


if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")
