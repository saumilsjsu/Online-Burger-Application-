<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1">
<!-- Add icon library -->
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css">
<style>
body {font-family: Mclawsuit  ;}
* {box-sizing: border-box;}

#pin_header{color: white  ;
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




</head>

<body>
<center>
 <div id="pin_header" class="header">
	<center><h1>Find your Restaurant!</h1><center>
	<p></p>
	<p></p>
 </div>
<br>

<div id= "pin_tags" class="page_tags">
	<table align="right">
		<tbody>
			
		</tbody>
	</table>
</div>
<br>

<form method="post" action="rest_list.php" name="pinform" style="max-width:500px;margin:auto">

<div>
	<center><h2>Payment Successful</h2></center>
</div>



  <button type="submit" class="btn"><h4>Thanks for Shopping With Us</h4></button><br>
</form>
</center>
</body>
</html>