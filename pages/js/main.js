   // Обращение
       var name = 'slavik';
       document.getElementById('wlcmsg').innerHTML = `Welcome, ${name}`;


    // Форма
    let authform = document.forms["authentication"];
    let fd = new FormData(authform);
    let data = {};

    for (let [key, prop] of fd) {
        data[key] = prop;
    }

    data = JSON.stringify(data, null, 2);
    console.log(data)

