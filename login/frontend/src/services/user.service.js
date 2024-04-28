import HttpRequest from "./api.service";
let http = new HttpRequest();

export const userService = {
    postAddUser,
};

function postAddUser(param) {
    const url = `${import.meta.env.VITE_API_URL}/user`;
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