
 export  class FCMModel{
    constructor(Id,TokenClient,UserId,AppName,IsLogin){
        this.Id = Id;
        this.TokenClient = TokenClient;
        this.UserId = UserId;
        this.AppName = AppName;
        this.IsLogin = IsLogin;
    }

}

export class MessageModel{
    AppName = '';
    CreateDate = '';
    Id = '';
    isread = false;
    notification = Notification;
    userId = '';
}

export class Notification {
    title = '';
    body = ''
}



export class MessageAllModel {
//    constructor(Result,TotalRows){
//        super(MessageModel)
//        this.Result = MessageModel;
//        this.TotalRows = TotalRows;
//    }
    // super(MessageModel)
    Result = MessageModel;
    TotalRows = 0;
}

export class MessageRequest{
    constructor(UserId,AppName,Limit){
        this.UserId = UserId;
        this.AppName = AppName;
        this.Limit = Limit;
    }
}