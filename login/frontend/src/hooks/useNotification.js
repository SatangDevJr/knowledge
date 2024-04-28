import { notification } from 'ant-design-vue';

export default {
    okNotification,
    infoNotification,
    errorNotification,
}

function okNotification(title, descriotion) {
    notification.success({
        message: `${title}`,
        description: `${descriotion}`,
    })
}

function infoNotification(title, descriotion) {
    notification.info({
        message: `${title}`,
        description: `${descriotion}`,
    })
}

function errorNotification(title, descriotion) {
    notification.error({
        message: `${title}`,
        description: `${descriotion}`,
    })
}