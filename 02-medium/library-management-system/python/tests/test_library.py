from item_type import ItemType
from library_management_system import LibraryManagementSystem


def test_checkout_and_return() -> None:
    library = LibraryManagementSystem()
    copies = library.add_item(ItemType.BOOK, "B001", "The Hobbit", "Tolkien", 1)
    member = library.add_member("MEM01", "Alice")

    library.checkout(member.get_id(), copies[0].get_id())
    assert not copies[0].is_available()

    library.return_item(copies[0].get_id())
    assert copies[0].is_available()
