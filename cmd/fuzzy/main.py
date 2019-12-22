from fuzz_common import visualize_lex_var, zadeh_and, zadeh_or, mass_center
from matplotlib import pyplot as plt
from vars import priority, preparedness, expediency, rules
from matplotlib import pyplot as plt

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
    plt.show(block = False)

    priorityValue = int(input('Показатель приоритета профиля для студента: '))
    preparednessValue = int(input('Уровень профильной подготовки студента: '))

    expediencyValue = mamdani_inference(
        {priority.name: priorityValue, preparedness.name: preparednessValue},
        rules, expediency)

    print("Показатель целесообразности обучения по профилю: {}".format(expediencyValue))
