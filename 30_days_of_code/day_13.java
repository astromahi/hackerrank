// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-13-abstract-classes
// Day 13: Abstract Classes!
import java.util.*;
abstract class Book
{
    String title;
    String author;
    Book(String t,String a){
        title=t;
        author=a;
    }
    abstract void display();
}

//Write MyBook Class
class MyBook extends Book
{
    int price;
    MyBook(String _title, String _author, int _price) {
        super(_title, _author);
        this.price = _price;
    }
    
    void display() {
        System.out.println("Title: "+title+"\nAuthor: "+author+"\nPrice: "+this.price);
    }
}

public class Solution
{
   
   public static void main(String []args)
   {
      Scanner sc=new Scanner(System.in);
      String title=sc.nextLine();
      String author=sc.nextLine();
      int price=sc.nextInt();
      Book new_novel=new MyBook(title,author,price);
      new_novel.display();
      
   }
}

