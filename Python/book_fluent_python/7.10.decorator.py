# 创建一个装饰器 工厂函数，把参数传给它，返回一个装饰器，然后再把它应用到要装饰 的函数上
registry = set()


def register(active=True):
    def decorate(func):
        print(f"begin register {func}")
        if active:
            registry.add(func)
        else:
            registry.discard(func)
        return func

    return decorate


@register()
def f1():
    print("f1 function")


@register(active=False)
def f2():
    print("f2 function")


def f3():
    print("f3 function")


print(registry)

# 闭包
"""
闭包指延伸了作用域的函数，其中包含函数定义体中引用、但是
不在定义体中定义的非全局变量。函数是不是匿名的没有关系，关键是
它能访问定义体之外定义的非全局变量
"""


def make_avg():
    series = []

    def average(val):
        series.append(val)
        total = sum(series)
        return total / len(series)

    return average


avg = make_avg()
print(avg(10))
print("闭包", avg(12), avg.__code__.co_freevars, avg.__code__.co_varnames)
print(avg.__code__.co_code.decode())
# 局部变量
b = 9


def fl(a):
    print(a)
    # print(b) # Python 不要求声明变量，但是假定在函数 定义体中赋值的变量是局部变量
    # b = 3


fl(2)
