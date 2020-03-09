from Crypto.PublicKey import RSA
from Crypto.Util.number import long_to_bytes, bytes_to_long, inverse
from Crypto import Random
from operator import mul
from functools import reduce
from decimal import Decimal, localcontext

MESSAGE = b'tryin to create a message long enough to give different values when being encrypted'
EXP = 17

def generate_key(exp=EXP):
    random_generator = Random.new().read
    key = RSA.generate(1024, random_generator, e = exp)
    return key.publickey()

def encrypt(m=MESSAGE):
    key = generate_key()
    key.encrypt
    return key.encrypt(bytes_to_long(MESSAGE), 0)[0], key.n

def CRT(modules, values):
    sum_ = 0
    prod = reduce(mul, modules, 1)

    for modules_i, values_i in zip(modules, values):
        p = prod // modules_i
        sum_ += values_i * inverse(p, modules_i) * p
    return sum_ % prod

def separate(data):
    ciphertexts = tuple(data[i][0] for i in range(len(data)))
    modules = tuple(data[i][1] for i in range(len(data)))
    return ciphertexts, modules

def root(num, e):
    with localcontext() as context:
        context.prec = 1025
        exp = Decimal(1.) / Decimal(e)
        return int(Decimal(num) ** exp)

def crack(data, exp):
    ciphertexts, modules = separate(data)
    crt_value = CRT(modules, ciphertexts)
    return long_to_bytes(root(crt_value, exp))

if __name__ == "__main__":
    data = tuple(encrypt() for i in range(EXP))
    print(crack(data, EXP))
