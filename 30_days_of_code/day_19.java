// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-19-interfaces
// Day 19: Interfaces!

import java.io.*;
import java.util.*;
interface AdvancedArithmetic{
   int divisorSum(int n);
}

//Write your code here
class Calculator implements AdvancedArithmetic {
    int sum = 0;
    public int divisorSum(int n) {
        int j = 1;
        while (j <= n) {
            if (n % j == 0) {
                //System.out.println("n : " + n + " j : " + j );
                sum = sum + j;
            }
            j++;
        }
        return sum;
    }
}

public class Solution {

    public static void main(String[] args) {
        Scanner sc=new Scanner(System.in);
        int n=sc.nextInt();
      	AdvancedArithmetic myCalculator=new Calculator(); 
        int sum=myCalculator.divisorSum(n);
        System.out.println("I implemented: AdvancedArithmetic\n"+sum);
    }
}