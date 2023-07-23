import arrow
from celery import Celery
from celery.backends.redis import RedisBackend
from celery.schedules import crontab

broker_url = 'redis://:123@139.196.213.1:6379/1'
backend_url = 'redis://:123@139.196.1.108:6379/2'


class CustomRedisBackend(RedisBackend):
    task_keyprefix = 'cronjob'

    def __init__(self, app=None, **kwargs):
        kwargs['host'] = kwargs.get("url") or backend_url
        kwargs['app'] = app
        super().__init__(**kwargs)

    def get_key_for_task(self, task_id, key=''):
        """
        overwrite cache key
        """
        key_t = self.key_t
        return key_t('_').join([
            self.task_keyprefix, key_t(task_id)[:9], key_t(key),
        ])


_app = Celery(
    'hello',
    broker='redis://:Pass123456@139.196.213.108:6379/1',
)

_app.conf.update(
    result_backend='test2.CustomRedisBackend',
)


@_app.task
def my_cronjob(arg1):
    # Code for your periodic task goes here
    print("Executing periodic task")
    return f"my_periodic_task: {str(arrow.now())}, arg {str(arg1)}"


@_app.on_after_configure.connect
def setup_periodic_tasks(sender, **kwargs):
    sender.add_periodic_task(crontab(minute='*'), my_cronjob.s(1))


if __name__ == '__main__':
    # celery -A app beat --loglevel=info
    _app.worker_main(['worker', '-B'])
