/*
58. Length of Last Word

Given a string s consisting of words and spaces, return the length of the last word in the string.

A word is a maximal consisting of non-space characters only.



Example 1:

Input: s = "Hello World"
Output: 5
Explanation: The last word is "World" with length 5.

Example 2:

Input: s = "   fly me   to   the moon  "
Output: 4
Explanation: The last word is "moon" with length 4.

Example 3:

Input: s = "luffy is still joyboy"
Output: 6
Explanation: The last word is "joyboy" with length 6.

*/

package main

import "strings"

func lengthOfLastWord(s string) int {
	lastWord := strings.Split(strings.TrimSpace(s), " ")
	return len(lastWord[len(lastWord)-1])
}

/*
without using builtin functions

func lengthOfLastWord(s string) int {
    ans:=0
    prev:=0
    for i:=0;i<len(s);i++{
        if s[i]==' '{
            if ans>0{
                prev=ans
            }
            ans=0
        }else{
            ans+=1
        }
    }
    if ans==0{
        return prev
    }
    return ans
}
*/
