


export const utils = {
    _orderby ,
    _uuid,
    // numberInput,
}
// _.orderBy(users, ['user', 'age'], ['asc', 'desc']);
function _orderby(data,key,value){
    const _ = require("lodash");
    return _.orderBy(data,[key],[value])
}

function _uuid(){
    const id = uuidv4();
    return id;
}
// function numberInput(inputValue) {
//     if (
//       inputValue.charAt(inputValue.length - 1) === '.' ||
//       inputValue === '-'
//     ) {
//       format(inputValue.slice(0, -1), inputValue);
//     }
//   };