import { AES, enc } from 'crypto-js'

const secretPassphrase = "testLogin-be3f4a5a-afe1-4c5f-9bd7-db13cc19c831";

export const storageService = {
    setStorage,
    getStorage
}

function setStorage(key, value) {
    return localStorage.setItem(`${key}`, encrypt(value));
}

function getStorage(key) {
    let result = ""
    try {
        let data = localStorage.getItem(`${key}`);
        if (data != null) {
            result = decrypt(data)
        }
    } catch (e) {}
    return result;
}

function encrypt(value) {
    const plaintext = JSON.stringify(value);
    const encrypt = AES.encrypt(plaintext, secretPassphrase).toString();
    return encrypt;
}

function decrypt(encrypt) {
    const bytes = AES.decrypt(encrypt, secretPassphrase);
    const plaintext = bytes.toString(enc.Utf8);
    return plaintext;
}