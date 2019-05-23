<html>
<body>
<?php
if ($_SERVER["REQUEST_METHOD"] == "POST") {
    // collect value of input field
	$email = $_POST['email'];
	$password=$_POST['password'];
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

?>
</body>
</html>
