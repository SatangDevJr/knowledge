import {
    Form,
    Table,
    Breadcrumb,
    PageHeader,
    Input,
    Button,
    InputSearch,
    Badge,
    Modal,
    Select,
    MonthPicker,
    Divider,
    Typography,
    Tag,
    List,
    Upload,
    Spin
} from "ant-design-vue";
const antdConfig = function(app) {
    app
        .use(Upload)
        .use(List)
        .use(Tag)
        .use(Typography)
        .use(Table)
        .use(Breadcrumb)
        .use(PageHeader)
        .use(Input)
        .use(InputSearch)
        .use(Button)
        .use(Badge)
        .use(Modal)
        .use(Form)
        .use(Select)
        .use(MonthPicker)
        .use(Divider)
        .use(Spin);
};
export default antdConfig;