function IsEmail(email) {
    var regex = /^([a-zA-Z0-9_\.\-\+])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$/;
    if(!regex.test(email)) {
        return false;
    }else{
        return true;
    }
}

function IsPhonenumber(phonenumber) {
    var regex = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im;
    if(!regex.test(phonenumber)) {
        return false;
    }else{
        return true;
    }
}

$(document).ready(function() {
    $('#form').submit(function(event) {
        event.preventDefault();
        var username = $.trim($('#username').val())
        var phonenumber = $.trim($('#phonenumber').val())
        var email = $.trim($('#email').val())
        var password = $.trim($('#password').val())
        var confirm_password = $.trim($('#confirm-password').val())

        flag = true

        if(username.length <= 4) {
            $('#username-error').css('visibility', 'visible')
            flag=false
        }

        if(!IsPhonenumber(phonenumber)) {
            $('#phonenumber-error').css('visibility', 'visible')
            flag=false
        }

        if(!IsEmail(email)) {
            $('#email-error').css('visibility', 'visible')
            flag=false
        }

        if(password.length < 5) {
            $('#password-error').css('visibility', 'visible')
            flag=false
        }

        if(password != confirm_password) {
            $('#confirm-password-error').css('visibility', 'visible')
            flag=false
        }

        if(flag) {
            $('.form-items-text').css('visibility', 'hidden')
        }
    })
})