import scipy.integrate as spi
import numpy as np
from matplotlib import pyplot as plt


class Variable(object):
    def __init__(self, name, memberships, left, right):
        self.name = name
        self.memberships = memberships
        self.left = left
        self.right = right

    def fuzzify(self, x):
        if x < self.left or x > self.right:
            return [0 for _ in self.memberships]
        result = {}
        for (name, memb) in self.memberships.items():
            result[name] = memb(x)
        return result


class Rule(object):
    def __init__(self, lhs_names, lhs, rhs, rhs_name, op='AND'):
        self.lhs_names = lhs_names
        self.lhs = list(lhs)
        self.rhs = rhs
        self.rhs_name = rhs_name
        self.op = op

    def infer(self, lex_var_values):
        values = []
        # fuzzification
        for i in range(len(self.lhs)):
            ev = self.lhs[i].fuzzify(lex_var_values[self.lhs[i].name])
            values.append(ev[self.lhs_names[i]])
        # aggregation
        if self.op is 'AND':
            return self.rhs_name, constant(self.rhs.left, self.rhs.right, min(values))

        return self.rhs_name, constant(self.rhs.left, self.rhs.right, max(values))


def trapezioid(left_start, left_max, right_max, right_end, value=1):
    def inner(x):
        if x <= left_start:
            return 0
        elif x <= left_max:
            return (x - left_start) / (left_max - left_start) * value
        elif x <= right_max:
            return value
        elif x <= right_end:
            return (right_end - x) / (right_end - right_max) * value
        else:
            return 0

    return inner


def triangle(left, right, value):
    middle = (left + right) / 2
    return trapezioid(left, middle, middle, right, value)


def constant(left, right, value):
    def inner(x):
        if (x < left):
            return 0
        elif (x > right):
            return 0
        return value

    return inner


def zadeh_and(left, right):
    return lambda x: min(left(x), right(x))


def prod_and(left, right):
    return lambda x: left(x) * right(x)


def zadeh_or(left, right):
    return lambda x: max(left(x), right(x))


def prod_or(left, right):
    return lambda x: left(x) + right(x) - left(x) * right(x)


def mass_center(distr, left, right):
    def denominator(x):
        return x * distr(x)

    return spi.quad(denominator, left, right, limit=100, limlst=100)[0] / spi.quad(distr, left, right,
                                                                                   limit=100, limlst=100)[0]


def visualize_lex_var(lex_var):
    X = np.linspace(lex_var.left, lex_var.right)
    plt.figure()
    for (name, membership) in lex_var.memberships.items():
        values = []
        for x in X:
            values.append(membership(x))
        plt.plot(X, values, label=name)
    plt.title(lex_var.name)