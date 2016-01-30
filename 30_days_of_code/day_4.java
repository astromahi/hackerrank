
import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

public class Person {

    private int age;

    public Person(int initial_Age) {
        // Add some more code to run some checks on initial_Age
        if (initial_Age < 0) {
            this.age = 0;
            System.out.println("This person is not valid, setting age to 0.");
        }

        if (initial_Age >= 0) {
            this.age = initial_Age;
        }

    }

    public void amIOld() {
        // Do some computations in here and print out the correct statement to the console 
        // using System.out.println()...
        if (this.age < 13) {
            System.out.println("You are young.");
        }

        if (this.age >= 13 && this.age < 18) {
            System.out.println("You are a teenager.");
        }

        if (this.age >= 18) {
            System.out.println("You are old.");
        }

    }

    public void yearPasses() {
        // Increment the age of the person in here
        this.age++;
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int T = sc.nextInt();
        for (int i = 0; i < T; i++) {
            int age = sc.nextInt();
            Person p = new Person(age);
            p.amIOld();
            for (int j = 0; j < 3; j++) {
                p.yearPasses();
            }
            p.amIOld();
            System.out.println();
        }
    }
}
