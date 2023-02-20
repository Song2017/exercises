from collections import defaultdict, Mapping, deque
from collections import OrderedDict
from collections import ChainMap
from collections import Counter

my_dict = {}
print(isinstance(my_dict, Mapping))
dd = defaultdict(tuple)
print(dd['default'])

od = OrderedDict()
od[2] = 2
od[1] = 1
print(od.keys())

cd1 = {1: 11, "2": 22, "test": 3}
cd2 = {"2": "test2", 4: "4"}
cd12 = ChainMap(cd1, cd2)
cd21 = ChainMap(cd2, cd1)
print(cd12["2"], cd21["2"], cd21.keys())

c = Counter('gallahad')
c2 = Counter({'red': 4, 'blue': 2})
c3 = Counter(cats=4, dogs=8)
print(c, sorted(c2.elements()), sorted(c3.elements()))

d = deque("test")
d.append("end")
d.appendleft("begin")
d.extend('second')
print(d)
d.rotate(-1)
print(d)

def round_robin(*iterables):
    iterators = deque(map(iter, iterables))
    while iterators:
        try:
            while True:
                yield next(iterators[0])
                iterators.rotate(-1)
        except StopIteration:
            iterators.popleft()


print(list(round_robin('ABC', 'D', 'EF')))
