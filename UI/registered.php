<html>
<body>
<?php
if ($_SERVER["REQUEST_METHOD"] == "POST") {
	// collect value of input field
	$name = $_POST['name'];
	$address= $_POST['address'];
	$pincode= $_POST['pincode'];
        $email = $_POST['email'];
	$password=$_POST['password'];
	  if (empty($name)) {
        	echo "Name is empty";
   	 }
    	else {
       		 echo $name;
	}
	echo "<br>";
	  if (empty($address)) {
        	echo "Address is empty";
    	 }
    	else {
       		 echo $address;
	}
	echo "<br>";
    	   if (empty($pinocode)) {
        	echo "Pincode is empty";
   	 }
    	else {
        	echo $pincode;
    	}
	echo "<br>";
   	 if (empty($email)) {
        	echo "Email is empty";
   	 }
    	else {
        	echo $email;
    	}
   	echo "<br>";
    	if (empty($password)) {
        	echo "Password is empty";
   	 }
    	else {
       		 echo $password;
    }
    $url = '';
    $ch = curl_init($url);
         $data=array(
		 'name'=>$name,
		 'address'=>$address,
		 'pincode'=>$pincode,
		 'email'=>$email,
		  'password'=$password
);
         $payload=json_encode(array("user" => $data));

        curl_setopt($ch, CURLOPT_POSTFIELDS,$payload);
        curl_setopt($ch,CURLOPT_HTTPHEADER, array('Content-Type:application/json'));
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        $result = curl_exec($ch);
        curl_close($ch);
}

$data = json_decode(file_get_contents('php://input'), true);

