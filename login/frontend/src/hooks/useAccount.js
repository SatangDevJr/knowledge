import { accountService } from '../services/account.service'
import useNotification from "./useNotification";
import { useI18n } from "../i18n";
const { t } = useI18n();
export default {
    authenticationLogin,
    ExternalAndInternalLogin,
}

function authenticationLogin(username, password) {

    const params = new URLSearchParams()
    params.append('grant_type', 'password')
    params.append('username', `${username}`)
    params.append('password', `${password}`)
    params.append('client_id', 'cd1371f7-6757-4950-9628-fcf7b5fed41c')
    params.append('scope', 'ptp.api.write')
    params.append('client_secret', 'TUVSWt3n6SUd6sywcA8yBNVykWAjYCPD')

    return accountService.authenLogin(params).then((res) => {
        if (res.status != 200) {
            useNotification.errorNotification(
                `${t("login.message.error.title")}`,
                `${res.data.message}`
            );
            return res
        } else {
            return res
        }

    })
}

function ExternalAndInternalLogin(username, password, stafflogin) {
    const params = {
        grant_type: 'password',
        username: username,
        password: password,
    }

    return accountService.ExternalAndInternalLogin(params).then((res) => {
        if (res.status != 200) {
            useNotification.errorNotification(
                `${t("login.message.error.title")}`,
                `${res.data.message}`
            );
            return res
        } else {
            return res
        }
    })
}