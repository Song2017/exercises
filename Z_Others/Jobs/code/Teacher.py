from cached_property import cached_property, threaded_cached_property


class Teacher:
    _id = '123'

    def worker_id(self):
        return self._id
    @property
    def worker_id_number(self):
        return self._id
    @cached_property
    def worker_id_cached(self):
        return self._id
    @threaded_cached_property
    def worker_id_cached(self):
        return self._id        


if __name__ == "__main__":
    # 请修改Teacher类
    s = Teacher()
    s._id = "345"
    assert s.worker_id() == "345"
    assert s.worker_id_number == "345"
    assert s.worker_id_cached == "345"

    s._id = "678"
    assert s.worker_id_number == "678"
    assert s.worker_id_cached == "345"
