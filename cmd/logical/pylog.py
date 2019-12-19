from src.Solver import Solver
import pandas as pd

if __name__ == "__main__":
    with open('examples/logical_rules', 'r') as content_file:
        content = content_file.read()
    solver = Solver(content)
    print(solver.database)
    while True:
        inp = input("запрос> ")
        solution = solver.find_solutions(inp)
        if isinstance(solution, dict):
            print(pd.DataFrame.from_dict(solution))
        else:
            print(solution)