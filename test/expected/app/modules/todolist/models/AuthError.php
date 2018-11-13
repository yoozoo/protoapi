<?php

namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class AuthError extends ProtoApi\BizErrorException implements ProtoApi\Message
{
    protected $message;

    public function init(array $response)
    {
        if (isset($response["message"])) {
            $this->message = $response["message"];
        }
    }

    public function validate()
    {
        if (!isset($this->message)) {
            throw new ProtoApi\GeneralException("'message' is not exist");
        }
    }
    
    public function set_message($message)
    {
        $this->message = $message;
    }

    public function get_message()
    {
        return $this->message;
    }
    
    public function to_array()
    {
        return array(
            "message" => $this->message,
        );
    }
}
