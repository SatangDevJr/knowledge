import HttpRequest from "./api.service";
let http = new HttpRequest();
export const accountService = {
    authenLogin,
    ExternalAndInternalLogin,
}

function authenLogin(param) {
    const url = `${import.meta.env.VITE_OAUTH}`;
    const config = {
        headers: {
            'Content-Type': 'application/x-www-form-urlencoded'
        }
    }
    const response = http
        .postConfig(url, param, config)
        .then((res) => {
            return res;
        })
        .catch(function(error) {
            return error;
        });
    return response;
}

function ExternalAndInternalLogin(param) {
    const url = `${import.meta.env.VITE_API_URL}/auth/login`;
    const response = http
        .post(url, param, {})
        .then((res) => {
            return res;
        })
        .catch(function(error) {
            return error;
        });
    return response;
}