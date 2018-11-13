<?php
namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class Blank implements ProtoApi\Message
{

    public function init(array $response)
    {
    }

    public function validate()
    {
    }
    
    public function to_array()
    {
        return array(
        );
    }
}
