// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-22-binary-search-trees
// Day 22: Binary Search Trees

import java.util.*;
import java.io.*;
class Node{
    Node left,right;
    int data;
    Node(int data){
        this.data=data;
        left=right=null;
    }
}
class Solution {
		public static int getHeight(Node root){
      //Write your code here
        if (root == null) {
            return 0;
        }
        
        int left_half = getHeight(root.left);
        int right_half = getHeight(root.right);
        
        int height;
        if (left_half > right_half) {
            height = left_half + 1;
        } else {
            height = right_half + 1;
        }
        
        return height;
    }

        public static Node insert(Node root,int data){
        if(root==null){
            return new Node(data);
        }
        else{
            Node cur;
            if(data<=root.data){
                cur=insert(root.left,data);
                root.left=cur;
            }
            else{
                cur=insert(root.right,data);
                root.right=cur;
            }
            return root;
        }
    }
	 public static void main(String args[]){
        Scanner sc=new Scanner(System.in);
        int T=sc.nextInt();
        Node root=null;
        while(T-->0){
            int data=sc.nextInt();
            root=insert(root,data);
        }
        int height=getHeight(root);
        System.out.println(height);
    }
}
