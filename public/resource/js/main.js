var daterangepicker_locale = {
    "format": "YYYY-MM-DD",
    "separator": " - ",
    "applyLabel": "确定",
    "cancelLabel": "取消",
    "fromLabel": "从",
    "toLabel": "到",
    "customRangeLabel": "自定义",
    "daysOfWeek": [
        "日",
        "一",
        "二",
        "三",
        "四",
        "五",
        "六"
    ],
    "monthNames": [
        "一月",
        "二月",
        "三月",
        "四月",
        "五月",
        "六月",
        "七月",
        "八月",
        "九月",
        "十月",
        "十一月",
        "十二月"
    ],
    "firstDay": 1
};

function alertMessageSuccess(message) {
    alertMessageCommon(message, 'success', 'messenger-on-bottom');
}

function alertMessageError(message) {
    alertMessageCommon(message, 'error', 'messenger-on-bottom');
}

function alertMessageCommon(message, type, position) {
    Messenger({
        extraClasses: "messenger-fixed " + position
    }).post({
        message: message,
        type: type
    });
}

function js_inarray(value, array) {
    var count = 0;
    for (var i = 0; i < array.length; i++) {
        if (value == array[i]) {
            count++;
        }
    }

    return count;
}

function confirmSwal(title, text, func) {
    swal({
        title: title,
        text: text,
        type: "warning",
        showCancelButton: true,
        closeOnConfirm: false,
        showLoaderOnConfirm: true,
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        animation: false
    }, function () {
        func();
    });
}

function alertSwal(title, text, iconstr) {
    var icontype = '';
    switch (iconstr) {
        case 'success':
        case 'error':
        case 'warning':
        case 'info':
            icontype = iconstr;
            break;
        default:
            icontype = 'info';
    }
    swal({
        title: title,
        text: text,
        type: iconstr,
        confirmButtonText: '好的',
        animation: false
    });
}

function promptSwalConfig(title, text) {
    return {
        title: title,
        text: text,
        type: "input",
        showCancelButton: true,
        closeOnConfirm: false,
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        animation: false
    };
}