package prime

import math

def isPrime(x Int) Bool = {
    isPrime(x, math.Abs(math.Sqrt(x)))
}

def isPrime(x, y Int) Bool = {
    if (y <= 1) {
        True
    } else if (x % y == 0) {
        False
    } else {
        isPrime(x, y - 1)
    }
}

def main = {
    val a = 10
    val d = "a string too"

    println(a
    println(d)

    println(isPrime(10))
    println(isPrime(7))
    println(isPrime(0))
    println(isPrime(900))
}