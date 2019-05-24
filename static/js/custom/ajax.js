$('document').ready(function () {
    var huduma_state = false;
    var service_state = false;

    $('#huduma').on('blur', function () {
        var huduma = $('#huduma').val();
        if (huduma == '') {
            huduma_state = false;
            return;
        }
        $.ajax({
            url: '/referpatient',
            type: 'post',
            data: {
             
                'huduma': huduma,
            },
            success: function (response) {
                if (response.data){
                    huduma_state = false;
                    $('#huduma').parent().removeClass();
                    $('#huduma').parent().addClass("form_error");
                    $('#huduma').siblings("span").text('Patient does not exist');
                } else if (response.data) {
                    username_state = true;
                    $('#huduma').parent().removeClass();
                    $('#huduma').parent().addClass("form_success");
                    $('#huduma').siblings("span").text('Patient available');
                }
            }
        });
    });

    $('#service').on('blur', function () {
        var service = $('#service').val();
        if (service == '') {
            service_state = false;
            return;
        }
        $.ajax({
            url: '/referpatient',
            type: 'post',
            data: {
                                'service': email,
            },
            success: function (response) {
                if (response.data) {
                    service_state = false;
                    $('#service').parent().removeClass();
                    $('#service').parent().addClass("form_error");
                    $('#service').siblings("span").text('Service does not exist');
                } else {
                    service_state = true;
                    $('#service').parent().removeClass();
                    $('#service').parent().addClass("form_success");
                    $('#service').siblings("span").text('Service available');
                }
            }
        });
    });

    $('#reg_btn').on('click', function () {
        var huduma = $('#huduma').val();
        var service = $('#eservice').val();
        if (huduma_state == false || service_state == false) {
            $('#error_msg').text('Fix the errors in the form first');
        } else {
            // proceed with form submission
            $.ajax({
                url: '/referpatient',
                type: 'post',
                data: {
                   
                    'service': service,
                    'huduma': huduma,

                },
                success: function (response) {
                    alert('Details submitted');
                    $('#huduma').val('');
                    $('#service').val('');
                }
            });
        }
    });
});