// Скрип отправки данных из формы к славику


const getFormData = (formName) => {
    const data = {};
    const userNameMain = '';
    console.log("побежали отправлять");
    const authform = document.forms[formName];
    const fd = new FormData(authform);
    for (let [key, prop] of fd) {
        data[key] = prop;
    };
    dataToSend = JSON.stringify(data);
    console.log(`data: ${dataToSend}`);
    console.log(`name: ${data.user_login}`);
    console.log(`password: ${data.user_password}`);
   // userNameMain = fd.data[]
}

// function sendJSON(formName) {
//     getFormData(formName);
//     const req = new XMLHttpRequest();
//     req.open("POST","/example/text.txt",false);
//     console.log("start data sending")
//     req.send(dataToSend);
//     console.log("data sent")
//     console.log(req.status, req.statusText);
//     console.log(`req: ${req.responseText}`);
// };

const sayError = () => {
    alert ("Please, check fields of form, at least one of these are empty");
}

function validateLoginForm ( ) {
        valid = true;
            if (( document.authentication.user_login.value == "" ) || ( document.authentication.user_password.value == "" ))
            {       $('.user_login').addClass('form_alert');
                    $('.user_password').addClass('form_alert');
                    // $('.user_login').css('border-color','red');
                    // $('.user_password').css('border-color','red')
                    setTimeout(sayError,500)
                    // alert ("Please, check fields of form, at least one of these are empty");
                    valid = false;
            } else {
                $('.user_login').removeClass('form_alert');
                $('.user_password').removeClass('form_alert');
            }
            return valid;
};

const formHandler = (formName) => {
    let callback
    if (validateLoginForm()) {
            let sending = new Promise(function (resolve,reject) {
                getFormData(formName);
                const req = new XMLHttpRequest();
                req.open("POST","/example/text.txt",false);
                console.log("start data sending")
                req.send(dataToSend);
                console.log("data sent")
                console.log(req.status, req.statusText);
                console.log(`req: ${req.responseText}`);

                if (req.status === 200) {
                    resolve(callback = true)
                } else {
                    reject(callback = false)
                }
            })

            sending.then(() => {
                setTimeout(document.authentication.reset(),1000);
                console.log('Всё ок, код 200 инфа сотка')
                redirectTo("messages.html")
            })
            sending.catch(() => {
                console.log('Всё пошло по пизде, наверн код 200 не вернулося(')
            })
        }
};



// Обращение к пользователю
const welcomeUser = (name,blockID) => {
    if (userNameMain = '') {
        document.getElementById(`${blockID}`).innerHTML = `Welcome, ${name}`;
    } else {
        document.getElementById(`${blockID}`).innerHTML = `Welcome, ${user_login}`;
    }
};

const logOut = () => {
  alert("Вы не можете выйти, тк разрабы дураки и ПОКА не знают как это сделать.")
};

const redirectTo = (link) => {
    document.location.href = link;
}