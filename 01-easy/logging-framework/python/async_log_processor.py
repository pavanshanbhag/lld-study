import logging
from concurrent.futures import ThreadPoolExecutor

from log_appender import LogAppender
from log_message import LogMessage

logger = logging.getLogger(__name__)


class AsyncLogProcessor:
    def __init__(self):
        self.executor = ThreadPoolExecutor(max_workers=1, thread_name_prefix="AsyncLogProcessor")
        self.shutdown_flag = False

    def process(self, log_message: LogMessage, appenders: list[LogAppender]):
        if self.shutdown_flag:
            logger.warning("Logger is shut down. Cannot process log message.")
            return

        def process_task():
            for appender in appenders:
                appender.append(log_message)

        self.executor.submit(process_task)

    def stop(self):
        self.shutdown_flag = True
        self.executor.shutdown(wait=True)
