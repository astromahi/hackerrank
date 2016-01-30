// https://www.hackerrank.com/contests/30-days-of-code/challenges/day-28-regex-patterns-intro-to-databases
// Day 28: RegEx, Patterns, and Intro to Databases!

import java.util.Scanner;
import java.util.regex.*;

public class Solution {

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int testCases = Integer.parseInt(in.nextLine());
        while (testCases > 0) {
            String pattern = in.nextLine();
            //Write your code
            try {
                Pattern p = Pattern.compile(pattern);
                if (p != null) {
                    System.out.println("Valid");
                }
            } catch (PatternSyntaxException ex) {
                System.out.println("Invalid");
            }

        }
    }
}
