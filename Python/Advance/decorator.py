import functools


def log(level):
    print("log level: ", level)

    def log_decorator(func):
        @functools.wraps(func)
        def decorator(request):
            print("decorator parameter", str(request))
            return func(request)

        return decorator

    return log_decorator


@log('debug')
def test(i_req):
    print("test name", test.__name__, str(i_req))
    return "test result"


# print("end", test({"q": 123}))

log_level = log('debug')
test_decorator = log_level(test)
print(test_decorator({"q": 123}))

import pdb
pdb.set_trace()