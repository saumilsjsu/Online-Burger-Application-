<?php

if($_SERVER["REQUEST_METHOD"] == "POST") 
{
        // collect value of input field
 
        $pincode= $_POST['pincode'];

	if (empty($pincode)) 
	{
		echo "Pincode is empty, please enter your pincode";
	}
	
	$url = '';
	$ch =curl_init($url);
	$pin_data=array(
			'pincode'=>$pincode
		);
		$payload=json_encode(array("" => $pin_data));

		curl_setopt($ch, CURLOPT_POSTFIELDS,$payload);
	        curl_setopt($ch,CURLOPT_HTTPHEADER, array(''));
        	curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        	$result = curl_exec($ch);
        	curl_close($ch);
}
$pin_data=json_decode (file_get_contents('php://input),true);
?>