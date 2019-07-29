package common

import "regexp"

func ValidateEmail(email string) bool {
    Re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return Re.MatchString(email)
}

func ValidateMobile(phone string) bool {
    Re := regexp.MustCompile(`^[+][0-9]{11,}`)
    return Re.MatchString(phone)
}


