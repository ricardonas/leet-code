
class LinkedList {
  constructor(value) {
    this.head = {
      value: value,
      next: null
    };
    this.tail = this.head;
    this.length = 1;
  }

  append(value) {
    const newNode = {
      value: value,
      next: null
    }
    this.tail.next = newNode;
    this.tail = newNode;
    this.length++;
    return this;
  }

  prepend(value) {
    const newNode = { value: value, next: this.head };
    this.head = newNode;
    this.length++;
    return this;
  }

  lookup(index) {
    let counter = 0;
    let node = this.head;

    while (counter !== index) {
      node = node.next;
      counter++;
    }

    return node;
  }

  insert(index, value) {
    const newNode = { value: value, next: null };
  
    const leader = this.lookup(index - 1);
    const pointerToNextLeaderNodes = leader.next;
    
    leader.next = newNode;
    newNode.next = pointerToNextLeaderNodes;

    this.length++;
  }

  // [ 5, 3, 7, 10 ]
  remove(index) {
    const leader = this.lookup(index - 1);
    const pointer = leader.next.next;
    leader.next = pointer;
    this.length--;    
  }

  toString() {
    const array = [];
    let node = this.head;
    while (node !== null) {
      array.push(node.value);
      node = node.next;
    }
    return array;
  }
  
}
