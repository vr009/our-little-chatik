// Скрип отправки данных из формы к славику


const getLoginFormData = () => {
    const loginData = {};
    let formName = "login_form";
    console.log("побежали отправлять");
    const authLoginForm = document.forms[formName];
    const fd_login = new FormData(authLoginForm);
    for (let [key, prop] of fd_login) {
        loginData[key] = prop;
    };
    dataToSend = loginData;
    console.log(`data: ${dataToSend}`);
    console.log(`name: ${loginData.Username}`);
    console.log(`password: ${loginData.Password}`);

   // userNameMain = fd.data[]
}

const getSignupnFormData = () => {
    const signupData = {};
    let formName = "signup_form";
    console.log("побежали отправлять");
    const authSignUpForm = document.forms[formName];
    const fd_signup = new FormData(authSignUpForm);
    for (let [key, prop] of fd_signup) {
        signupData[key] = prop;
    };
    dataToSend = signupData;
    console.log(`data: ${dataToSend}`);
    console.log(`FirstName: ${signupData.Firstname}`);
    console.log(`Lastname: ${signupData.Lastname}`);
    console.log(`Login: ${signupData.Username}`);
    console.log(`Password: ${signupData.Password}`);

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
};

function validateLoginForm ( ) {
        valid = true;
            if (( document.login_form.Username.value == "" ) || ( document.login_form.Password.value == "" ))
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

function validateSignupForm ( ) {
    valid = true;
    if (( document.signup_form.Firstname.value == "" ) || ( document.signup_form.Lastname.value == "" ))
    {   $('.user_firstName').addClass('form_alert');
        $('.user_lastName').addClass('form_alert');
        $('.user_login').addClass('form_alert');
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

//
// const getData = async() => {
//     try{
//         const response = await fetch("192.168.1.205",{
//             method: "POST",
//             body: JSON.stringify({id: 200})
//         })
//         if (response.ok) {
//             const jsonResponse = await response.json()
//             return jsonResponse
//         }
//         throw new Error('Request failed!');
//     }
//     catch(error) {
//         console.log(error)
//     }
// }
//

const formHandler = (formName) => {
    let callback

    if (formName === "login_form") {
        if (validateLoginForm()) {
            let sending = new Promise(function (resolve,reject) {
                getLoginFormData();
                const req = new XMLHttpRequest();
                // замени тут ниже url на нужный http://192.168.1.205/auth/signup !!!
                // Поставил async true (надо разобраться зачем :) )

                req.open("POST","http://192.168.1.205/auth/signup",true);
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
    } else if (formName === "signup_form") {
        if (validateSignupForm()) {
            let sending = new Promise(function (resolve,reject) {
                const req = new XMLHttpRequest();
                getSignupnFormData();
                // СОЗДАЕМ ТЕСТОВЫЙ ОБЪЕКТ

                const testMessage = {
                    Firstname: "testName",
                    Lastname: "testLastname",
                    Username: "testUsername",
                    Password: "testPassword"
                }

                // замени тут ниже url на нужный http://192.168.1.205/auth/signup !!!
                // Поставил async true (надо разобраться зачем :) )
                req.open("POST","http://192.168.1.205/auth/signup",true);
                console.log("start data sending")
                console.log(dataToSend);
                req.send(JSON.stringify(dataToSend));
                console.log("data sent");
                console.log(JSON.stringify(dataToSend));
                console.log(req.status, req.statusText);
                console.log(`req: ${req.responseText}`);

                if (req.status === 303) {
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
                console.log('Всё пошло по пизде, наверн код 200 не вернулося либо ошибка раньше(')
            })
        }
    }


};

// TODO

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

