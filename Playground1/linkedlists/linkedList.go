package linkedlists
import (
    "fmt"
)

type Node struct {
    value int;
    next *Node;
}

func newNode(value int) *Node {
    var newNode *Node = new(Node);
    newNode.value = value;
    newNode.next = nil;
    return newNode;
}

var head *Node = new(Node);

func insert_node(position, value int) {
    
    if position == 1 {
        newNode := newNode(value);
        head = newNode;
    } else { 
    temp := head;
   for i:=1;i<position; i++{
            if temp.next !=nil {
                temp = temp.next;
            }

   }
    if temp.next == nil {
        temp.next = newNode(value);     
    } else {
        newNode := newNode(value);
        nextNode := temp.next;
        temp.next = newNode;
        newNode.next = nextNode;
    }
}
}

func delete_node(position int) {
    if position == 1 {
        head = head.next;
    }
    var temp = new(Node);
    temp = head;
    for i:=1; i<position-1;i++{
        if temp.next != nil {
            temp = temp.next;
        }
    }
    if temp.next!=nil && temp.next.next != nil{
    nextNode := temp.next.next;
    temp.next = nextNode;
    }else {
        temp.next = new(Node);
    }
}   


func print_ll() {
    temp := head;
    for &temp != nil  {
       fmt.Print(temp.value);
       if temp.next !=nil {
           fmt.Print(" ");
           temp = temp.next;
       }else {
           break;
       } 

    }
}



