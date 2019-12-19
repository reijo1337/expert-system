from src.Solver import Solver

if __name__ == "__main__":
    with open('examples/logical_rules', 'r') as content_file:
        content = content_file.read()
    solver = Solver(content)
    print(solver.database)
    while True:
        inp = input("запрос> ")
        print(solver.find_solutions(inp))