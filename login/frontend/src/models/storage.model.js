// import UserInfo  from '../models/user.model'
class StorageModel {
    constructor(token = null, tokenExpire, userInfo
    ,menuPermission,lang,firebaseToken,tokenThaibev,tokenData) {
        this.token = token;
        this.tokenExpire = tokenExpire;
        this.userInfo = userInfo;
        this.menuPermission = menuPermission
        this.lang = lang
        this.firebaseToken = firebaseToken
        this.tokenThaibev= tokenThaibev
        this.tokenData = tokenData
    }
  }

export default StorageModel;