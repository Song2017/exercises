from queue import Queue

if __name__ == "__main__":
    # 线程通信
    # 1. 线程通信方式- 共享变量
    test = 1
    # 2. queue import Queue
    # deque 是线程安全的
    detail_url_queue = Queue(maxsize=1000)
    # 线程同步: 保证调用同一变量的先后执行
    from threading import Lock, RLock, Condition

    lock = Lock()  #
    lock.acquire()
    lock.release()
    lock = RLock()  # 可重入的锁, 可连续调用acquire
    lock.acquire()
    lock.release()
    cond = Condition()
    cond.acquire()
    cond.wait()
    cond.notify()
    cond.release()

    # 进程间通信
    from multiprocessing import Process, Queue, Pool, Manager, Pipe

    # Queue
    # 要使用multiprocessing， 或是manager的
    # multiprocessing中的queue不能用于pool进程池
    # pool中的进程间通信需要使用manager中的queue
    queue = Queue()
    manager = Manager()
    m_queue = manager.Queue(10)
    # pipe pipe的性能高于queue pipe只能适用于两个进程
    receive_pipe, send_pipe = Pipe()
    # 进程间通信的数据结构
    # Manager中提供的数据结构 progress_dict = Manager().dict()
