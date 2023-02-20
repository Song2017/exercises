import itertools
import re
from abc import abstractmethod, ABCMeta

RE_WORD = re.compile('\w+')
"""
1. 迭代器接口
__next__
  返回下一个可用的元素，如果没有元素了，抛出 StopIteration
异常。
__iter__
  返回 self，以便在应该使用可迭代对象的地方使用迭代器，例如在 for 循环中
__getitem__
    如果没有实现 __iter__ 方法，但是实现了 __getitem__ 方法， Python 会创建一个迭代器

2. Iterable VS Iterator
迭代器
  迭代器是这样的对象:实现了无参数的 __next__ 方法，返回序列 中的下一个元素;
  如果没有元素了，那么抛出 StopIteration 异常。 
  Python 中的迭代器还实现了 __iter__ 方法，因此迭代器也可以迭代
Iterable: __iter__
Iterator: __iter__ and __next__

  
3. 可迭代的对象 vs 迭代器
Sentence vs SentenceIterator
可迭代的对象有个 __iter__ 方法，每次都实例化一个新的迭代 器;
而迭代器要实现 __next__ 方法，返回单个元素，此外还要实现 __iter__ 方法，返回迭代器本身
迭代器模式可用来:
   访问一个聚合对象的内容而无需暴露它的内部表示
   支持对聚合对象的多种遍历
   为遍历不同的聚合结构提供一个统一的接口(即支持多态迭代)
   
4. 生成器函数
迭代器用于从集合中取出元素;而生成器用于“凭空”生成元素
  所有生成器都是迭代器，因为生成器完全实现了迭代器接口
只要 Python 函数的定义体中有 yield 关键字，该函数就是生成器函 数。
调用生成器函数时，会返回一个生成器对象。也就是说，生成器函 数是生成器工厂
5. 惰性实现
re.finditer 函数是 re.findall 函数的惰性版本，返回的不是列 表，而是一个生成器，按需生成 re.MatchObject 实例
调用生成器函数返回生成器;生成器产出或生成值
6. 生成器表达式
生成器表达式可以理解为列表推导的惰性版本:不会迫切地构建列表，
而是返回一个生成器，按需惰性生成元素。也就是说，如果列表推导是
制造列表的工厂，那么生成器表达式就是制造生成器的工厂。

9. 内置生成器工具itertools
10. yield from
可以用来代替循环
yield from 还会创建通道，把内层生成器直接与外层生成器的客户端联系起来
13. 生成器处理大量数据
生成器用于生成供迭代的数据
协程是数据的消费者
https://github.com/fluentpython/isis2json/blob/master/isis2json.py
yield 关键字只能把最近的外层函数变成生成器函数。 虽然生成器函数看起来像函数，可是我们不能通过简单的函数调用 把职责委托给另一个生成器函数

"""  # noqa:E501

s = "ss"
iter_s = iter(s)
for _ in range(3):
    try:
        print(next(iter_s))
    except StopIteration as e:
        print("except", e)


class Iterable(metaclass=ABCMeta):
    __slots__ = ()

    @abstractmethod
    def __iter__(self):
        while False:
            yield None

    @classmethod
    def __subclasshook__(cls, C):
        if cls is Iterable:
            if any("__iter__" in B.__dict__ for B in C.__mro__):
                return True
        return NotImplemented


class Iterator(Iterable):
    __slots__ = ()

    @abstractmethod
    def __next__(self):
        'Return the next item from the iterator. When exhausted, raise StopIteration'
        raise StopIteration

    def __iter__(self):
        return self

    @classmethod
    def __subclasshook__(cls, C):
        if cls is Iterator:
            if (any("__next__" in B.__dict__ for B in C.__mro__) and
                any("__iter__" in B.__dict__ for B in C.__mro__)):
                return True
        return NotImplemented


class Sentence:
    def __init__(self, text):
        self.text = text
        self.words = RE_WORD.findall(text)

    def __repr__(self):
        return f"text: {self.text}, words {self.words}"

    def __iter__(self):
        return SentenceIterator(self.words)


class SentenceIterator:
    def __init__(self, words):
        self.words = words
        self.index = 0

    def __next__(self):
        try:
            next_word = self.words[self.index]
        except IndexError:
            raise StopIteration()
        self.index += 1
        return next_word

    def __iter__(self):
        return self


class SentenceGen:
    def __init__(self, text):
        self.text = text
        self.words = RE_WORD.findall(text)

    def __repr__(self):
        return f"text: {self.text}, words {self.words}"

    def __iter__(self):
        for _w in self.words:
            yield _w
        return


class SentenceGenLazy:
    def __init__(self, text):
        self.text = text

    def __repr__(self):
        return f"SentenceGenLazy text: {self.text}"

    def __iter__(self):
        # for match in RE_WORD.finditer(self.text):
        #     yield match.group()
        return (match.group() for match in RE_WORD.finditer(self.text))


ss = SentenceGenLazy("test, 23 44 ads")
print(ss)
for w in ss:
    print(w)

for ig in itertools.groupby("LLLLAaAGGGGEEE"):
    print(ig)


def chain(*_iterators):
    for item in _iterators:
        for _i in item:
            yield _i


def chain2(*_iterators):
    for item in _iterators:
        yield from item


ss = "asdf"
ll = list(range(3))
print(list(chain(ss, ll)))
print(list(chain2(ss, ll)))


def f():
    def do_yield(_n):
        yield _n

    x = 0
    while x < 10:
        x += 1
        do_yield(x)


def f2():
    def do_yield(_n):
        yield _n

    x = 0
    while x < 10:
        x += 1
        yield from do_yield(x)


f1 = f()
f2 = f2()
print(type(f1), type(f2))
