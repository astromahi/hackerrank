// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-14-scope-and-imports
// Day 14: All about Scope!

import java.io.*;
import java.util.*;
import java.text.*;
import java.math.*;
import java.util.regex.*;

class Difference {

    private int[] elements;
    public int maximumDifference;

    // Add your code here

    Difference(int[] a) {
        this.elements = a;
    }

    int computeDifference() {
        for (int i = 0; i < elements.length; i++) {
            for (int j = i + 1; j < elements.length; j++) {
                int d = Math.abs(elements[i] - elements[j]);
                if (d >= maximumDifference) {
                    this.maximumDifference = d;
                }
            }
        }
        return this.maximumDifference;
    }
} // End of Difference class

public class Solution {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int N = sc.nextInt();
        int[] a = new int[N];
        for (int i = 0; i < N; i++) {
            a[i] = sc.nextInt();
        }

        Difference D = new Difference(a);

        D.computeDifference();

        System.out.print(D.maximumDifference);
    }
}
