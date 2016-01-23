// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-17-exceptions
// Day 17: Exceptions!

import java.util.*;
import java.io.*;

//Write your code here
class Calculator {
    
    int power(int n, int p) {
        if (n < 0 || p < 0) {
            throw new ArithmeticException("n and p should be non-negative");
        }
         
        int power = (int) Math.pow((double) n, (double) p);
        
        return power;
    }
   
}

class Solution{

    public static void main(String []argh)
    {
        Scanner in = new Scanner(System.in);
        int T=in.nextInt();
        while(T-->0)
        {
            int n = in.nextInt();
            int p = in.nextInt();
            Calculator myCalculator = new Calculator();
            try
            {
                int ans=myCalculator.power(n,p);
                System.out.println(ans);
                
            }
            catch(Exception e)
            {
                System.out.println(e.getMessage());
            }
        }

    }
}
