
    function submitForm() {
        document.getElementById("post_id").submit();
    }
    document.getElementById("comments-text").addEventListener("keyup", function(event) {
        if (event.key === "Enter") {
            submitForm();
        }
    });
