<?php
## Commiting Final
session_start();
?>
<!DOCTYPE html>
<html lang="en"><head>
<TITLE>PHP Shopping Cart without Database</TITLE>
<link href="style.css" type="text/css" rel="stylesheet" />
</head>
<body>
<h1>Demo Shopping</h1>


<?php 
require_once "product-gallery.php";
?>
<div class="clear-float"></div>
<div id="shopping-cart">
<div class="txt-heading">Shopping Cart <a id="btnEmpty" class="cart-action" onClick="cartAction('empty','');"><img src="images/icon-empty.png" /> Empty Cart</a></div>
<div id="cart-item">
<?php 
require_once "ajax-action.php";
?>



</div>
</div>



<form method="post" action="payment.php" name="loginform" style="max-width:500px; margin:auto">

<div>
  <center><h2>Payment</h2></center>
</div>

  <button type="submit" class="btn"><h4>Payment</h4></button><br>
</form>




<script src="jquery-3.2.1.min.js" type="text/javascript"></script>
<script>
function cartAction(action, product_code) {
    var queryString = "";
    if (action != "") {
        switch (action) {
        case "add":
            queryString = 'action=' + action + '&code=' + product_code
                    + '&quantity=' + $("#qty_" + product_code).val();
            break;
        case "remove":
            queryString = 'action=' + action + '&code=' + product_code;
            break;
        case "empty":
            queryString = 'action=' + action;
            break;
        }
    }
    jQuery.ajax({
        url : "ajax-action.php",
        data : queryString,
        type : "POST",
        success : function(data) {
            $("#cart-item").html(data);
            if (action == "add") {
                $("#add_" + product_code + " img").attr("src",
                        "images/icon-check.png");
                $("#add_" + product_code).attr("onclick", "");
            }
        },
        error : function() {
        }
    });
}
</script>
</body>
</html>
