import base64
import json
from urllib.parse import unquote

import requests
import rsa
from asn1crypto.keys import PublicKeyInfo, RSAPublicKey

# reference https://blog.xiazhiri.com/Mercury-MIPC251C-4-Reverse.html
"""
POST /stok=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx/ds

# 获取信息
{"method":"get","device_info":{"name":["basic_info"]}}

# 获取信息2
{"method":"get","function":{"name":["module_spec"]}}

# 获取预设位置
POST /stok=7e82c0908c8b141c8e9bb8353b54e259/ds
{"method":"get","preset":{"name":["preset"]}}

# 控制云台转到预设位置
{"method":"do","preset":{"goto_preset":{"id":"1"}}}

# 云台水平/垂直移动
{"method":"do","motor":{"move":{"x_coord":"10","y_coord":"0"}}}

# 云台步进 direction 0 / 90
{"method":"do","motor":{"movestep":{"direction":"0"}}}

# 云台停止
{"method":"do","motor":{"stop":"null"}}

# 增加预置点
{"method":"do","preset":{"set_preset":{"name":"name","save_ptz":"1"}}}

# 获取镜头遮蔽信息
{"method":"get","lens_mask":{"name":["lens_mask_info"]}}

# 镜头遮蔽
{"method":"set","lens_mask":{"lens_mask_info":{"enabled":"on"}}}
"""


def tp_encrypt(pwd: str) -> str:
    base = "RDpbLfCPsJZ7fiv"
    char_dict = (
        "yLwVl0zKqws7LgKPRQ84Mdt708T1qQ3Ha7xv3H7NyU84p21BriUWBU43odz3iP4rBL3cD02KZciX"
        "TysVXiV8ngg6vL48rPJyAUw0HurW20xqxv9aYb4M9wK1Ae0wlro510qXeU07kV57fQMc8L6aLgML"
        "wygtc0F10a0Dg70TOoouyFhdysuRMO51yY5ZlOZZLEal1h0t9YQW0Ko7oBwmCAHoic4HYbUyVeU3"
        "sfQ1xtXcPcf1aT303wAQhv66qzW"
    )
    tp_encrypted_pwd = ""

    base_l = len(base)
    pwd_l = len(pwd)
    char_dict_l = len(char_dict)

    for index in range(max(base_l, pwd_l)):
        cl = cr = 0xBB
        if index >= base_l:
            cr = ord(pwd[index])
        elif index >= pwd_l:
            cl = ord(base[index])
        else:
            cl = ord(base[index])
            cr = ord(pwd[index])

        tp_encrypted_pwd += char_dict[(cl ^ cr) % char_dict_l]

    return tp_encrypted_pwd


def convert_rsa_key(key: str) -> rsa.PublicKey:
    pub_keyinfo: RSAPublicKey = PublicKeyInfo.load(base64.b64decode(key))[
        "public_key"
    ].parsed
    return rsa.PublicKey(
        int.from_bytes(pub_keyinfo["modulus"].contents, "big"),
        int.from_bytes(pub_keyinfo["public_exponent"].contents, "big"),
    )


def rsa_encrypt(message: str, key: str) -> str:
    rsa_pubkey = convert_rsa_key(key)
    crypto = rsa.encrypt(message.encode(), rsa_pubkey)
    return base64.b64encode(crypto).decode()


def get_stok(url, username, password):
    # get key nonce
    print("-get rsa and nonce")
    j = post_data(url, json.dumps({"method": "do", "login": {}}))
    key = unquote(j["data"]["key"])
    nonce = str(j["data"]["nonce"])
    print("rsa: ", key)
    print("nonce: ", nonce)

    # encrypt tp
    print("--encrypt password by tp")
    tp_password = tp_encrypt(password)
    tp_password += ":" + nonce
    print("tp_password: ", tp_password)

    # rsa password
    print("--encrypt password by rsa")
    rsa_password = rsa_encrypt(tp_password, key)
    print("rsa_password: ", rsa_password)

    # login
    d = {
        "method": "do",
        "login": {
            "username": username,
            "encrypt_type": "2",
            "password": rsa_password,
        },
    }
    print("--login")
    j = post_data(url, json.dumps(d))
    stok = j["stok"]
    print(stok)
    return stok


def post_data(base_url, data, stok=""):
    url = base_url + (("/stok=" + stok + "/ds") if stok else "")
    print("post: ", url, " data: ", data)
    r = requests.post(url, data)
    print("response: ", str(r.status_code), " ", str(r.json()))
    return r.json()
