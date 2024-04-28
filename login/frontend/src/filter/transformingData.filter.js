import moment from 'moment';
import numeral from 'numeral';

const transformingData = {
    NumberFormat(value) {
        return numeral(value).format('0,0.00')
    },
    NumberFormat0digit(value) {
        return numeral(value).format('0,0')
    },
    NumberFormatMathfloor(value) {
        return numeral(value).format('0,0.00.[00]',Math.floor)
        // return numeral(value).format('0,0.00.[00]')
    },
    NumberFormat3digit(value) {
        return numeral(value).format('0,0.000')
    },
    DateFromat(value) {
        if(value != ""){
            let val = value;
            let Y = val.substring(0, 4);
            let M = val.substring(5, 7);
            let D = val.substring(8, 10);
            value = D + "/" + M + "/" + Y;
        }
        return value;
    },
    Dateformat(value,formatIn,formatOut){
        return moment(value,formatIn).format(formatOut)
    },
    Timeformat(val){
        return moment(val,'YYYY-MM-DD HH:mm:ss').format('HH:mm')
    },
    ReplaceMATERIAL(value){
        return value.replace('000000000',"")
    },
    bytesToSizeformat(bytes) {
        const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
        if (bytes === 0) return 'n/a'
        const i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)), 10)
        if (i === 0) return `${bytes} ${sizes[i]})`
        return `${(bytes / (1024 ** i)).toFixed(1)} ${sizes[i]}`
    },
    bytesToSize(bytes) {
        const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB']
        if (bytes === 0) return 'n/a'
        const i = parseInt(Math.floor(Math.log(bytes) / Math.log(1024)), 10)
        if (i === 0) return `${bytes}`
        return `${(bytes / (1024 ** i)).toFixed(1)}`
    }
    
}
export default transformingData

