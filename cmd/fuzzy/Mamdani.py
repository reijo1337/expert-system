from fuzz_common import *
from matplotlib import pyplot as plt
from vars import *

class Rule(object):
    def __init__(self, lhs_names, lhs, rhs, rhs_name, op='AND'):
        self.lhs_names = lhs_names
        self.lhs = list(lhs)
        self.rhs = rhs
        self.rhs_name = rhs_name
        self.op = op

    def infer(self, lex_var_values):
        values = []

        for i in range(len(self.lhs)):
            ev = self.lhs[i].fuzzify(lex_var_values[self.lhs[i].name])
            values.append(ev[self.lhs_names[i]])

        if self.op is 'AND':
            return self.rhs_name, constant(self.rhs.left, self.rhs.right, min(values))

        return self.rhs_name, constant(self.rhs.left, self.rhs.right, max(values))


def aggregate(inferred_values, right_lexvar):
    result_fns = []
    for name, fn in inferred_values:
        memb_fn = right_lexvar.memberships[name]
        inferred_fn = zadeh_and(memb_fn, fn)
        result_fns.append(inferred_fn)
    result_fn = result_fns[0]

    for fn in result_fns[1:]:
        result_fn = zadeh_or(result_fn, fn)

    return result_fn


def defuzzify(lexvar, fn):
    return mass_center(fn, lexvar.left, lexvar.right)


def mamdani_inference(vals, rules, rhs):
    inferences = []
    for rule in rules:
        inferences.append(rule.infer(vals))
    resulting_fn = aggregate(inferences, rhs)
    return defuzzify(rhs, resulting_fn)


if __name__ == "__main__":
    visualize_lex_var(priority)
    visualize_lex_var(preparedness)
    visualize_lex_var(expediency)

    priorityValue = int(input('Показатель приоритета профиля для студента: '))
    preparednessValue = int(input('Уровень профильной подготовки студента: '))

    expediencyValue = mamdani_inference({priority.name: priorityValue, preparedness.name: preparednessValue},
                                [R1, R2, R3, R4, R5, R6, R7, R8, R9, R10, R11], expediency)

    print(expediencyValue)
