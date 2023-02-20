# 传统函数调用 过程 A->B->C
# 我们需要一个可以暂停的函数，并且可以在适当的时候恢复该函数的继续执行
# 出现了协程 -> 有多个入口的函数， 可以暂停的函数， 可以暂停的函数(可以向暂停的地方传入值)

def gen_func():
    # 1. 可以产出值， 2. 可以接收值(调用方传递进来的值)
    html = yield "http://projectsedu.com"
    print("gen_func", html)
    # receive value from send
    yield 2
    return "bobby"


if __name__ == "__main__":
    gen = gen_func()
    # 启动生成器方式有两种， next(), send
    # 在调用send发送非none值之前，我们必须启动一次生成器, 这回运行完第一个yield
    # 方式有两种， 二者完全等价。 1. gen.send(None), 2. next(gen)
    # print(next(gen))
    print(gen.send(None))
    print('before send')
    print(gen.send("ben"))
