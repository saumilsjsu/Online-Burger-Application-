<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<!-- Add icon library -->
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body {font-family: Mclawsuit  ;}
* {box-sizing: border-box;}

#rest_list_header{color: white  ;
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

/* Set a style for the submit button */
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
 <div id="rest_list_header" class="header">
  <center><h1>Here's your restaurants!</h1><center>
  <p></p>
  <p></p>
 </div>
<br>

<div id= "pin_tags" class="page_tags">
  <table align="right">
    <tbody>
      <tr>
        <formaction="rest_pin.php">
          <td>
            <?php
            echo "Hello There";
            ?>
          </td>
        <td>
          <button class="b1" type="submit">Home</button>
        </td>
        </form>
        <form action="yourcart.php">
          <td>
          <button class="b1" type="submit">My Cart</button>
        </td>
        </form>
        <form action="logout.php">  
          <td>
          <button class="b1" type="submit">Log Out</button>
        </td>
        </form>
        </td>
      </tr>
    </tbody>
  </table>
</div>
<br>

<div>
  <center><h2>Choose One</h2></center>
</div>

<table id="Restaurants">
  <tbody>

<?php

if($_SERVER["REQUEST_METHOD"] == "POST"){
  //$restaurant = $_POST['restaurant'];
  $pin_rest=$_POST['pin_rest'];
  $data = callAPI("GET",'http://cmpe281-327234648.us-east-1.elb.amazonaws.com/restaurant/zipcode/95113',json_encode(array()));
 // echo $data;
  $data1 = json_decode($data, true);
 // echo $data[$result];
 }

elseif($pin_rest != "pin_rest"){
      echo "No restaurants were found for this Pincode!";
}
  
  echo "<table>";
    foreach ($data1 as $key => $value){   
        echo "<tbody>";     
        echo "<tr>";
        echo "<form  method='POST' action=rest_menu.php>";
        echo "<td>".$value["restaurantName"]."</td>";
        //echo "<td>".$value["distance"]."</td>";
        echo '<td><button type="submit" class="btn" >Select</button></td>';
        echo "<input type='hidden' name='rest_id' value=".$value["restaurantName"]."/>";

        echo "</tr>";
        echo "</tbody>";

    }
   echo "</table>"; 

    ?>
</center>
</body>
</html>