from threading import Lock

class PaymentProcessor:
    _lock = Lock()

    def process_payment(self, payment):
        # Process payment using the selected payment method
        payment.process_payment()