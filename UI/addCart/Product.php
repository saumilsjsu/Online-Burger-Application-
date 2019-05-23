<?php

class Product
{

    public $productArray = array(
        "3DcAM01" => array(
            'id' => '1',
            'name' => 'Burger',
            'code' => '3DcAM01',
            'image' => 'product-images/burger.jpeg',
            'price' => '10.00'
        ),
        "USB02" => array(
            'id' => '2',
            'name' => 'Coke',
            'code' => 'USB02',
            'image' => 'product-images/coke.jpg',
            'price' => '2.00'
        ),
        "wristWear03" => array(
            'id' => '3',
            'name' => 'Fries',
            'code' => 'wristWear03',
            'image' => 'product-images/fries.jpeg',
            'price' => '5.00'
        )
    );

    public function getAllProduct()
    {
        return $this->productArray;
    }
}
