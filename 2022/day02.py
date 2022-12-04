import time
from pathlib import Path

INPUT_FILE = Path("day02_input.txt")


def get_file_inputs(input_text: str) -> list:
    return list(input_text.splitlines())

def get_home_selections(selection: str, score_row: list) -> int:
    match selection:
        case "X":
            return score_row[0]
        case "Y":
            return score_row[1]
        case "Z":
            return score_row[2]

def get_away_selections(selection: str, score_matrix: list) -> int:
    match selection:
        case "A":
            return score_matrix[0]
        case "B":
            return score_matrix[1]
        case "C":
            return score_matrix[2]

def home_score(players: list, score_matrix: list) -> int:
    selections_away = get_away_selections(players[0], score_matrix)
    selections_home = get_home_selections(players[1], selections_away)

    return selections_home

def get_truth_matrix_based_score(players: list, truth_matrix: list, score_matrix: list) -> int:    
    scores_away = get_away_selections(players[0], score_matrix)
    truths_away = get_away_selections(players[0], truth_matrix)

    # X to lose
    # Y to draw
    # Z to win
    match players[1]:
        case "X":
            return scores_away[truths_away.index("0")]
        case "Y":
            return scores_away[truths_away.index("X")]
        case "Z":            
            return scores_away[truths_away.index("1")]
        

def main():
    input_text = INPUT_FILE.read_text()
    games = get_file_inputs(input_text)

    score_home_part1 = 0
    score_home_part2 = 0

    # Use matrices in a opposite order ABC being away player and XYZ home player for easier list handling
    
    # Truth matrix for rock paper scissors
    """    X, Y, Z   """
    """ A [X, 0, 1], """
    """ B [1, X, 0], """
    """ C [0, 1, X], """

    truth_matrix = [
        ["X", "1", "0"],
        ["0", "X", "1"],
        ["1", "0", "X"],
    ]   

    # Score matrix for home

    # play_rock = 1
    # play_paper = 2
    # play_scissors = 3

    # lost_bonus = 0
    # draw_bonus = 3
    # win_bonus = 6
    
    """    X, Y, Z   """
    """ A [4, 8, 3], """
    """ B [1, 5, 9], """
    """ C [7, 2, 6], """

    score_matrix = [
        [4, 8, 3],
        [1, 5, 9],
        [7, 2, 6],
    ]

    for game in games:
        players = game.split()    
        score_home_part1 += home_score(players, score_matrix)
        score_home_part2 += get_truth_matrix_based_score(players, truth_matrix, score_matrix)

    print("Home score part1: ", score_home_part1)
    print("Home score part2: ", score_home_part2)


if __name__ == "__main__":
    start_time = time.time()
    main()
    print("Runtime:", time.time() - start_time, "seconds")
