class Node {
    constructor(value){
        this.value = value;
        this.next  = null;
    }
}

class Linklist {
    constructor(){
        this.head = new Node('head')
    }

    insert(value, item = null) {
        let node = new Node(value);
        let currNode;
        if(item){
            currNode = this.find(item);
            if(!currNode) return false;
            node.next = currNode.next;
            currNode.next = node;
        }else{
            currNode = this.head
            while(currNode && currNode.next) {
                currNode = currNode.next;
            }
            currNode.next = node;
        }
        return true;
    }

    remove(value){
        let prevNode = this.findPrev(value);
        let currNode = this.find(value);
        if(currNode == false) return false;
        if(currNode.next != null && prevNode.next){
            prevNode.next = currNode.next;
        }else{
            prevNode.next = null;
        }
        return true;
    }

    findPrev(value) {
        let currNode = this.head;
        while(currNode.next != null && currNode.next.value != value){
            currNode = currNode.next;
        }
        return currNode
    }

    find(value) {
        let currNode = this.head;
        while(currNode && currNode.value != value){
            currNode = currNode.next;
        }
        return currNode || false;
    }

    display(){
        let currNode = this.head;
        while(currNode){
            console.log(`${currNode.value}`);
            currNode = currNode.next;
        }
    }

    reverseList() {
        let tempNode = null;
        let nextNode = null;
        let currNode = this.head;
        while(currNode) {
            tempNode = currNode.next;
            currNode.next = nextNode;
            nextNode = currNode;
            currNode = tempNode;
        }
        this.head = nextNode;
    }
}

/* TEST CODE */
const linkList = new Linklist();
linkList.insert(1)
linkList.insert(2)
linkList.insert(3)
linkList.insert(4)
linkList.display()
// linkList.remove(3)
console.log('reverseList ----------')
linkList.reverseList()
linkList.display()
/* TEST END */

module.exports = Linklist;