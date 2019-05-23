<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<!-- Add icon library -->
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body {font-family: Mclawsuit  ;}
* {box-sizing: border-box;}

#order_list_header{color: white  ;
     background-image: url(burger_texture.jpg);
     background-position: left top;
     background-size:300px 300px ; 
     background-repeat: repeat;
     padding: 25px; 
    }

.input-container {
  display: -ms-flexbox; /* IE10 */
  display: flex;
  width: 100%;
  margin-bottom: 15px;
}

.icon {
  padding: 10px;
  background: dodgerblue;
  color: white;
  min-width: 50px;
  text-align: center;
}

.input-field {
  width: 100%;
  padding: 10px;
  outline: none;
}

.input-field:focus {
  border: 2px solid dodgerblue;
}

.btn {
  background-color: dodgerblue;
  color: white;
  padding: 4px 4px 8px 10px;
  border: none;
  cursor: pointer;
  width: 100%;
  opacity: 0.9;
}

.btn:hover {
  opacity: 1;
}

.btn1 {
background-color: dodgerblue;
  color: white;
  padding: 1px 1px 2px 2px;
  text-align: center;
  border: none;
  cursor: pointer;
  width: 100%;
  height: 25%;
  opacity: 0.9;
}

.btn1:hover {
 opacity: 1;
}

.b1
{
background-color: dodgerblue;
  color: white;
  padding: 1px 1px 2px 2px;
  text-align: center;
  border: none;
  cursor: pointer;
  width: 100%;
  height: 25%;
  opacity: 0.9;
}

.b1:hover {
 opacity: 1;
}

</style>

<?php

function callAPI($method, $url, $data) {
  //echo $data;
  $curl = curl_init();
  //echo $data;
   switch ($method){
      case "POST":
         curl_setopt($curl, CURLOPT_POST, 1);
         if ($data)
            curl_setopt($curl, CURLOPT_POSTFIELDS, $data);
         break;
      case "PUT":
         curl_setopt($curl, CURLOPT_CUSTOMREQUEST, "PUT");
         if ($data)
            curl_setopt($curl, CURLOPT_POSTFIELDS, $data);                
         break;
      default:
         if ($data)
            $url = sprintf("%s?%s", $url, http_build_query($data));
  }

   // OPTIONS:
   curl_setopt($curl, CURLOPT_URL, $url);
   curl_setopt($curl, CURLOPT_HTTPHEADER, array(
      'APIKEY: 111111111111111111111',
      'Content-Type: application/json',
   ));
   curl_setopt($curl, CURLOPT_RETURNTRANSFER, 1);
   curl_setopt($curl, CURLOPT_HTTPAUTH, CURLAUTH_BASIC);

   // EXECUTE:
   $result = curl_exec($curl);
   if(!$result){die("Connection Failure");}
    curl_close($curl);
    return $result;
 }

?>
</head>

<body>
<center>
 <div id="order_list_header" class="header">
  <center><h1></h1><center>
  <p></p>
  <p></p>
 </div>
<br>

<br>

<div>
  <center><h2><font color="green">Your Order is successful</font></h2></center>
</div>

<table id="Order">
  <tbody>

<?php

if($_SERVER["REQUEST_METHOD"] == "POST"){
    $order_id = $_POST['order_id'];
    $data = callAPI("GET",' http://cmpe281-327234648.us-east-1.elb.amazonaws.com/order/39829087-6bba-45d5-971e-9819fe9191bd \',json_encode(array()));
    // echo $data;
    $data2 = json_decode($data, true);
}
  
  echo '<form method="POST" action=Lastpage.php>';
  echo "<table>";
  echo "<tbody>";   
  foreach ($data2['items'] as $key => $value) {
        echo "<tr>";
        echo "<td>".$value["orderid"]."</td>";
        echo "<td>".'<button class="btn"><i class="fa fa-shopping-cart" aria-hidden="true"></i>'."</td>";
        echo "</tr>";
  }
  echo "</tbody>";
  echo "</table>"; 
  echo "</form>";
?>

</tbody>
</table>

</tbody>
</table>


</center>
</body>
</html>