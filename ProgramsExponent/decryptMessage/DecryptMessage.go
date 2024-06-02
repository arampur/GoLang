/*An infamous gang of cybercriminals named “The Gray Cyber Mob”, which is behind many hacking attacks and drug trafficking scandals,
has recently been targeted by the FBI. After intercepting a few messages which looked to be nonsense at first,
the agency realized that the group indeed encrypts their messages, and studied their method of encryption.

The essages consist of lowercase Latin letters only, and every word is encrypted separately as follows:

Convert every letter to its ASCII value.
Add 1 to the first letter, and then for every letter from the second one to the last one, add the value of the previous letter.
Subtract 26 from every letter until it is in the range of lowercase letters a-z in ASCII. Convert the values back to letters.
For instance, to encrypt the word "crime"

Step 1:	99
Step 2:	100
Step 3:	100
Result:	d
*

input:  word = "dnotq"
output: "crime"

input:  word = "flgxswdliefy"
output: "encyclopedia"*/

package main

func main() {

}

func Decrypt(word string) string {
	// iterate and convert every letter to ASCII value

	// Iterate through the letters. Add 1 to first, get the ASCII value of prev letter and add it

	// sub 26 from every letter and see ascii range letters a-z (97 - 122)
	return ""
}
