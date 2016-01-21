// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-15-linked-list
// Day 15: Linked List!

import java.io.*;
import java.util.*;
class Node{
	int data;
	Node next;
	Node(int d){
        data=d;
        next=null;
    }
	
}
class Solution
{
	public static  Node insert(Node head,int data)
	{
  	//Complete this method
        Node n = new Node(data);
        
        if (head == null)  {
            head = n;
            return head;
        }
     
        Node temp = head;
        while (temp.next != null) {
            temp = temp.next;
        }
        
        temp.next = n;
        
        //System.out.println("Head data: " + head.data + "Head next: " + head.next);
        return head;
	}

	public static void display(Node head)
    {
          Node start=head;
          while(start!=null)
          {
              System.out.print(start.data+" ");
              start=start.next;
          }
    }
    public static void main(String args[])
    {
          Scanner sc=new Scanner(System.in);
          Node head=null;
          int N=sc.nextInt();
          while(N-->0){
              int ele=sc.nextInt();
              head=insert(head,ele);
          }
          display(head);

   }
}