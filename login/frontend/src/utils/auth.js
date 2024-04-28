import decode from 'jwt-decode'
import axios from 'axios'
import useAccount from "../hooks/useAccount";
import { useAccountStore } from "../stores/account.store";
import { router } from "../router/index";
import { storageService } from '../services/storage.service'
import moment from "moment";

const AUTH_TOKEN_KEY = 'authToken'
const AUTH_Password_KEY = 'Password'

export function loginUser(username, password) {
    return new Promise(async(resolve, reject) => {
        try {
            useAccount
                .ExternalAndInternalLogin(username, password)
                .then((res) => {
                    if (res.status == 200) {
                        setAuthToken(res.data.Token);

                        storageService.setStorage('Password', password)
                        localStorage.setItem("UserName", username);
                        storageService.setStorage("Token", res.data.Token);

                        const accountStore = useAccountStore();
                        accountStore.setloginData(res.data);
                        isLoggedIn();
                        resolve();
                    } else {
                        logoutUser();
                        router.push("/login");
                    }
                });
        } catch (err) {
            console.error('Caught an error during login:', err);
            logoutUser();
            router.push("/login");
        }
    })
}

export function logoutUser() {
    localStorage.removeItem("UserName");
    localStorage.removeItem("Password");
    localStorage.removeItem("Token");
    localStorage.removeItem(AUTH_TOKEN_KEY);
}

export function setAuthToken(token) {
    axios.defaults.headers.Authorization = `Bearer ${token}`;

    return storageService.setStorage(AUTH_TOKEN_KEY, token)
}

export function getAuthToken() {
    return storageService.getStorage(AUTH_TOKEN_KEY)
}

export function clearAuthToken() {
    axios.defaults.headers.Authorization = ``;
    localStorage.removeItem(AUTH_TOKEN_KEY);
}

export function isLoggedIn() {
    let authToken = getAuthToken();
    return !!authToken
}

export function getUserInfo() {
    if (isLoggedIn()) {
        return decode(getAuthToken());
    }
}

function getTokenExpirationDate(encodedToken) {
    let token = decode(encodedToken);
    if (!token.exp) {
        return null;
    }

    let date = new Date(0);
    date.setUTCSeconds(token.exp);

    return date;
}

export function isTokenExpired(token) {
    let expirationDate = getTokenExpirationDate(token);
    let isExpiration = true;
    var day = new Date()
    var dayWrapper = moment(day);
    var dateString = dayWrapper.format("YYYYMMDDHHmmss");

    var expirationWrapper = moment(expirationDate);
    var expirationString = expirationWrapper.format("YYYYMMDDHHmmss");

    console.log("expirationDate : ", parseInt(expirationString, 10), expirationDate)
    console.log("date Now : ", parseInt(dateString, 10), day)

    if (parseInt(expirationString, 10) < parseInt(dateString, 10)) {
        console.log("token Date : ", expirationString, expirationDate)
        console.log("date Now : ", dateString, day)
        isExpiration = false;
    }

    return isExpiration;
}

export function getPassword() {
    return storageService.getStorage(AUTH_Password_KEY)
}