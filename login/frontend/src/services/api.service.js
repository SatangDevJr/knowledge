// import Vue from "vue";
// import { environment } from '../config'
import axios from "axios";
import { logoutUser } from '../utils/auth'
import { router } from "../router/index";

import { storageService } from '../services/storage.service'
//import { StorageService } from "./storage.service";
// import { alertmsgService } from "./alertmessage.service";

class HttpRequest {
    constructor() {
        // const Token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySW5mbyI6W3siVXNlcklkIjo0LCJFbWFpbCI6ImljZV9zdXBhbnV0QGhvdG1haWwuY29tIiwiQ29udGFjdE51bWJlciI6bnVsbCwiRnVsbE5hbWUiOiLguJnguLLguKIg4Lio4Li44Lig4LiT4Lix4LiQ4LiR4LmMIOC4p-C4h-C4qOC5jOC4q-C4tOC4o-C4seC4jeC4geC4oyIsIkxhbmd1YWdlQ29kZSI6IlRIIn0seyJVc2VySWQiOjQsIkVtYWlsIjoiaWNlX3N1cGFudXRAaG90bWFpbC5jb20iLCJDb250YWN0TnVtYmVyIjpudWxsLCJGdWxsTmFtZSI6Ik1yLiBTdXBhbnV0IFdvbmdoaXJhbmtvbiIsIkxhbmd1YWdlQ29kZSI6IkVOIn1dLCJFeHBpcmVzIjo4NjQwMH0.kJ1MqKC2uC2Dt5oFoCX1vwFdHR7qOJvrKvfN12AesYY';

        this.axiosInstance = axios.create({
            baseURL: "",
            timeout: 300000,
        });

        this.axiosInstance.interceptors.response.use(
            function(response) {
                // Do something with response data
                return response;
            },
            function(error) {
                return Promise.reject(error.response);
            }
        );
    }
    get(url) {
        return this.axiosInstance.get(url);
    }
    getConfig(url, config) {
        return this.axiosInstance.get(url, config);
    }
    postConfig(url, data, config) {
        return this.axiosInstance.post(url, data, config);
    }
    post(url, data) {
        return this.axiosInstance.post(url, data);
    }
    put(url, data) {
        return this.axiosInstance.put(url, data);
    }
    delete(url, data) {
        return this.axiosInstance.delete(url, data);
    }
    patch(url, data) {
        return this.axiosInstance.patch(url, data);
    }
}
export default HttpRequest;