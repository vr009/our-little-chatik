const getLoginFormData = () => {
    const loginData = {};
    let formName = "login_form";
    console.log("побежали отправлять");
    const authLoginForm = document.forms[formName];
    const fd_login = new FormData(authLoginForm);
    for (let [key, prop] of fd_login) {
        loginData[key] = prop;
    };
    const dataToSend = JSON.stringify(loginData);
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
    const dataToSend = JSON.stringify(signupData);
    console.log(`data: ${dataToSend}`);
    console.log(`FirstName: ${signupData.Firstname}`);
    console.log(`Lastname: ${signupData.Lastname}`);
    console.log(`Login: ${signupData.Username}`);
    console.log(`Password: ${signupData.Password}`);

    // userNameMain = fd.data[]
}

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