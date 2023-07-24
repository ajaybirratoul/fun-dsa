class Node(private val value: Int) {
    var both: Int = 0 // XOR of the previous and next node addresses

    fun getValue(): Int {
        return value
    }

    fun getPointer(): Int {
        return both
    }
}

class XORLinkedList {
    private var head: Node? = null
    private var tail: Node? = null

    fun add(value: Int) {
        val newNode = Node(value)
        if (head == null) {
            head = newNode
            tail = newNode
        } else {
            newNode.both = addressOf(tail) // XOR of 0 and previous tail's address
            tail?.both = addressOf(tail?.both ?: 0) xor addressOf(newNode)
            tail = newNode
        }
    }

    fun get(index: Int): Int? {
        var current = head
        var prevAddress = 0
        for (i in 0 until index) {
            val nextAddress = current?.both ?: 0 xor prevAddress
            if (nextAddress == 0) {
                return null
            }
            prevAddress = addressOf(current)
            current = getNodeFromAddress(nextAddress)
        }
        return current?.getValue()
    }

    private fun addressOf(node: Node?): Int {
        return node?.hashCode() ?: 0
    }

    private fun getNodeFromAddress(address: Int): Node? {
        // Simulate a XOR operation with addressOf to find the original node address
        return addressOf(address xor addressOf(null)) as? Node
    }
}

fun main() {
    val xorLinkedList = XORLinkedList()
    xorLinkedList.add(10)
    xorLinkedList.add(20)
    xorLinkedList.add(30)

    println(xorLinkedList.get(0)) // Output: 10
    println(xorLinkedList.get(1)) // Output: 20
    println(xorLinkedList.get(2)) // Output: 30
    println(xorLinkedList.get(3)) // Output: null
}
