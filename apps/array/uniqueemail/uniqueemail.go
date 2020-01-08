package main

import (
	"fmt"
	"strings"
)

// UniqueMail represents a unique email adress user@domain,
// where user is the unique user name (without . and + characters)
type UniqueMail struct {
    user string
    // including @ prefix
    domain string
}

func uniqueUser(email string) UniqueMail {
    var sb strings.Builder
    var uniqueMail UniqueMail
    
    if len(email) == 0 {
        return uniqueMail
    }
    
    skipToAt := false
     // Assume ASCII 
    pos := 0
    c := email[pos]   
    sb.WriteByte(c)
    for c != '@' {
        pos++
        c = email[pos]
        if c == '+' {
            // ignore the following part, until @ found
            skipToAt = true
        } else if c != '.' && c != '@' && !skipToAt {
            sb.WriteByte(c)    
        }
    }
    uniqueMail.user = sb.String()
    // pos is now at the @
    uniqueMail.domain = email[pos:len(email)]
    return uniqueMail
}

func numUniqueEmails(emails []string) int { 
    uniqueMails := make(map[UniqueMail]bool, 0)
    for _, email := range emails {
        uniqueMail := uniqueUser(email)
        uniqueMails[uniqueMail] = true
    }
    return len(uniqueMails)
}

func main() {
	emails := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	num := numUniqueEmails(emails)
	fmt.Println("Unique emails:", num)
}