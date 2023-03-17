jQuery(function ($) {
    "use strict";

    toastr.options = {
        "positionClass": "toast-top-center",
    };

    const wow = new WOW({});
    const form = $('#contact-form');
    const submitButton = form.find('input[type=submit]');

    wow.init();
    form.validator();

    form.on('submit', function (e) {
        if (!e.isDefaultPrevented()) {
            submitButton.attr('disabled', true)
            $.ajax({
                type: form.attr('method'),
                url: form.attr('action'),
                data: form.serialize(),
                success: function (response) {
                    toastr.success('Your message has been successfully sent!');
                    submitButton.attr('disabled', false)
                },
                error: function (response) {
                    toastr.error('Sorry, an error occurred while sending your message. Please try again in a few minutes.')
                    submitButton.attr('disabled', false)
                },
            });

            return false;
        }
    })
});
