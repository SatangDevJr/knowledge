import { userService } from "../services/user.service";
import useNotification from "./useNotification";
import { useI18n } from "../i18n";
const { t } = useI18n();

export default {
    onSaveUser
};

function onSaveUser(formState) {
    return userService.postAddUser(formState).then((res) => {
        if (res.status == 200) {
            useNotification.infoNotification(
                `${t("message.success.title")}`,
                `${t("message.success.content")}`
            );
            return res;
        } else {
            useNotification.errorNotification(
                `${t("message.error.title")}`,
                `${res.data.message}`
            );
            return res;
        }
    });
}