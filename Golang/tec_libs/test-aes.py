import base64

from Crypto import Random
from Crypto.Cipher import AES


def decrypt(key, value, block_segments=True):
    # The base64 library fails if value is Unicode. Luckily, base64 is ASCII-safe.
    value = str(value)
    # We add back the padding ("=") here so that the decode won't fail.
    value = base64.b64decode(value + '=' * (4 - len(value) % 4), '-_')
    iv, value = value[:AES.block_size], value[AES.block_size:]
    if block_segments:
        # Python uses 8-bit segments by default for legacy reasons. In order to support
        # languages that encrypt using 128-bit segments, without having to use data with
        # a length divisible by 16, we need to pad and truncate the values.
        remainder = len(value) % 16
        padded_value = value + b'\0' * (16 - remainder) if remainder else value
        cipher = AES.new(key, AES.MODE_CFB, iv, segment_size=128)
        # Return the decrypted string with the padding removed.
        return cipher.decrypt(padded_value)[:len(value)]
    return AES.new(key, AES.MODE_CFB, iv).decrypt(value)


def encrypt(key, value, block_segments=False):
    iv = Random.new().read(AES.block_size)
    if block_segments:
        # See comment in decrypt for information.
        remainder = len(value) % 16
        padded_value = value + b'\0' * (16 - remainder) if remainder else value
        cipher = AES.new(key, AES.MODE_CFB, iv, segment_size=128)
        value = cipher.encrypt(padded_value)[:len(value)]
    else:
        value = AES.new(key, AES.MODE_CFB, iv).encrypt(value)
    # The returned value has its padding stripped to avoid query string issues.
    return base64.b64encode(iv + value, b'-_').rstrip(b'=')

key = b'KEY61736466617364666173646661712'
print(decrypt(key, "5mntujUyhGu6Udu3RY3q4voiX9A"))
print(encrypt(key, "data阿斯顿发".encode(), block_segments=True))
