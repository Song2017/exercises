class Test:
    def __init__(self, te=False):
        super().__init__()
        self.te = te

    def __enter__(self):
        print("__enter__")
        if self.te:
            self.exception()
        return self

    def __exit__(self, exc_type, exc_val, exc_tb):
        print("exit" * 10)
        print("exit", exc_type, exc_val, exc_tb)

    def __str__(self):
        """
        __enter__ replace __str__ when __enter__ 返回字符串
        """
        return "test with __str__"

    def __repr__(self):
        return "test with __repr__"

    def exception(self):
        raise Exception("exception")

    def process(self):
        print("process")


if __name__ == '__main__':
    t1 = Test()
    print("t1", t1, "...\n")
    ''' 执行顺序
    __init__
    __enter__: 需要返回self, Exception会快速返回
    t.process(): 
    __exit__: 退出
    '''
    with Test(te=False) as t:
        print("t with")
        t.process()
        print("t print", t, "t print\n")
