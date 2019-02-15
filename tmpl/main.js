$(document).ready(function () {
    $("#status-error").hide();
    $("#status-success").hide();
    $("#status-processing").hide();

    // Handler for form submission
    $("form").submit(function(e) {
        e.preventDefault();
        $("#status").text("Processing");
        $("#status-error").hide();
        $("#status-success").hide();
        $("#status-processing").show();
        var formData = new FormData(this);    

        $.ajax({
            url : "/api",
            type: "POST",
            data : formData,
            processData: false,
            contentType: false,
            success: function(data, _, _) {
                $("#status-processing").hide();
                $("#status-success").show();
                $("#result").attr("src", "data:img/png;base64," + data)
            },
            error: function(data, _, _) {
                $("#status-processing").hide();
                $("#status-error").show();  
                $("#status-error").html(data.responseText);
            }
        });
    });

    // Show file name in input dialog
    $('.custom-file-input').on('change', function() { 
        var fileName = $(this).val().split('\\').pop(); 
        $(this).next('.custom-file-label').addClass("selected").html(fileName); 
    });
});