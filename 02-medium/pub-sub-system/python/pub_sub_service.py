from typing import Dict
from topic import Topic
from subscriber import Subscriber
from message import Message
from concurrent.futures import ThreadPoolExecutor


class PubSubService:
    def __init__(self):
        self.topic_registry: Dict[str, Topic] = {}
        self.delivery_executor = ThreadPoolExecutor()

    def create_topic(self, topic_name: str):
        if topic_name not in self.topic_registry:
            self.topic_registry[topic_name] = Topic(topic_name, self.delivery_executor)
        print(f"Topic {topic_name} created")

    def subscribe(self, topic_name: str, subscriber: Subscriber):
        topic = self.topic_registry.get(topic_name)
        if topic is None:
            raise ValueError(f"Topic not found: {topic_name}")
        topic.add_subscriber(subscriber)
        print(f"Subscriber '{subscriber.get_id()}' subscribed to topic: {topic_name}")

    def unsubscribe(self, topic_name: str, subscriber: Subscriber):
        topic = self.topic_registry.get(topic_name)
        if topic is not None:
            topic.remove_subscriber(subscriber)
        print(f"Subscriber '{subscriber.get_id()}' unsubscribed from topic: {topic_name}")

    def publish(self, topic_name: str, message: Message):
        print(f"Publishing message to topic: {topic_name}")
        topic = self.topic_registry.get(topic_name)
        if topic is None:
            raise ValueError(f"Topic not found: {topic_name}")
        topic.broadcast(message)

    def shutdown(self):
        print("PubSubService shutting down...")
        self.delivery_executor.shutdown(wait=True)
        print("PubSubService shutdown complete.")
