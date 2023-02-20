"""
!! yield 都是一种流程控制工具，使用它可以实现协作式多任务:
    协程可以把控制器让步给中心调度程序，从而激活其他的协程
    
    
- yield item 这行代码会产出一 个值，提供给 next(...) 的调用方;
    此外，还会作出让步，暂停执行 生成器，让调用方继续工作，直到需要使用另一个值时再调用 next()
- 协程可能会从调用方 接收数据，不过调用方把数据提供给协程使用的是 .send(datum) 方 法，而不是 next(...) 函数。
    通常，调用方会把值推送给协程    
1. 生成器如何进化成协程
yield 关键字可以在表达式 中使用，而且生成器 API 中增加了 .send(value) 方法。
    生成器的调用 方可以使用 .send(...) 方法发送数据，发送的数据会成为生成器函数 中 yield 表达式的值
除了 .send(...) 方法，PEP 342 还添加了 .throw(...) 和 .close() 方法:前者的作用是让调用方抛出异常，在生成器中处理;后者的作用 是终止生成器
协程是指一个过程，这个过程与调用方协作，产出由调用方提供的值
2. 用作协程的生成器的基本行为
b = yield a: 执行协程到 yield 表达式，然后产出 a 的值，并且暂停，等待为 b 赋
"""  # noqa:E501
import inspect


def simple_coroutine():
    print("coroutine start")
    x = yield
    print("coroutine received", x)


my_co = simple_coroutine()
print("my_co", my_co, inspect.getgeneratorstate(my_co))
print("激活 coroutine", next(my_co), inspect.getgeneratorstate(my_co))
try:
    print("send", my_co.send(44), inspect.getgeneratorstate(my_co))
except StopIteration as e:
    print(str(e), inspect.getgeneratorstate(my_co))


def simple_coro2(a):
    print("coroutine start a", a)
    b = yield a
    print("coroutine received b", b)
    c = yield a + b
    print('-> Received: c =', c)


my_co = simple_coro2(10)
print("my_co", my_co, inspect.getgeneratorstate(my_co))
print("激活 coroutine", next(my_co), inspect.getgeneratorstate(my_co))
try:
    print("send", my_co.send(20), inspect.getgeneratorstate(my_co))
    print("send", my_co.send(90), inspect.getgeneratorstate(my_co))
except StopIteration as e:
    print(str(e), inspect.getgeneratorstate(my_co))


def average():
    total = 0
    count = 0
    _avg = None
    while count < 10:
        term = yield _avg
        print("term, total", term, total)
        total += term
        count += 1
        _avg = total // count
    return 100


my_average = average()
next(my_average)
for i in range(15):
    try:
        print(my_average.send(i), inspect.getgeneratorstate(my_average))
    except Exception as e:
        print("Exception", str(e), inspect.getgeneratorstate(my_average))
