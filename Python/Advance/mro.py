#  MRO method resolution order (方法解释顺序)：C3算法,
# 		        没有共同祖先时, 类似深度优先
# 		        有共同祖先时,先进行深度优先查找,查到共同祖先就返回进行广度优先


class D:
    def test_message(self):
        print("Dddddd")


class E:
    def test_message(self):
        print("EEEE test_message(self)")


class F(D, E):
    def test_message(self):
        print("F")
        # E.test_message(self)
        super(E, F).test_message()
        # self.test_message()


# 深度遍历优先
f = F()
# print(F.__mro__)

f.test_message()
# ee = E()
# ee.test_message()
# A
# B C
# D E
# F
# enter F
# enter D
# enter B
# enter E
# enter C
# enter A
# leave A
# leave C
# leave E
# leave B
# leave D
# leave F

# class F(E, D):
#     def __init__(self):
#         print('enter F')
#         super().__init__()
#         print('leave F')
#
#
# print(F.__mro__)
