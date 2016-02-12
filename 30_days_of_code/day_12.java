// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-12-inheritance
// Day 12: Inheritance!

import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

class Student {

    protected String firstName;
    protected String lastName;
    int phone;

    Student(String fname, String lname, int p) {
        firstName = fname;
        lastName = lname;
        phone = p;

    }

    public void display() {// display the details of the student
        System.out.println("First Name: " + firstName + "\nLast Name: " + lastName + "\nPhone: " + phone);
    }

}

class Grade extends Student {

    private int score;

    Grade(String fname, String lname, int p, int _score) {
        super(fname, lname, p);
        this.score = _score;
    }

    public char calculate() {
        char grade;
        if (score < 40) {
            grade = 'D';
        } else if (score < 60) {
            grade = 'B';
        } else if (score < 75) {
            grade = 'A';
        } else if (score < 90) {
            grade = 'E';
        } else {
            grade = 'O';
        }

        return grade;
    }

}

public class Solution {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        String firstName = sc.nextLine();
        String lastName = sc.nextLine();
        int phone = sc.nextInt();
        int score = sc.nextInt();
        Student stu = new Grade(firstName, lastName, phone, score);
        stu.display();
        Grade g = (Grade) stu;
        System.out.println("Grade: " + g.calculate());

    }
}
