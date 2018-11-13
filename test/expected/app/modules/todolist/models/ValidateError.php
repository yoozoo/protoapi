<?php

namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class ValidateError extends ProtoApi\BizErrorException implements ProtoApi\Message
{
    protected $errors;

    public function init(array $response)
    {
        if (isset($response["errors"])) {
            $this->errors = array();
            foreach ($response["errors"] as $errors) {
                $tmp = new FieldError();
                $tmp->init($errors);
                $tmp->validate();
                $this->errors[] = $tmp;
            }
        }
    }

    public function validate()
    {
        if (!isset($this->errors)) {
            throw new ProtoApi\GeneralException("'errors' is not exist");
        }
    }
    
    public function set_errors(Errors $errors)
    {
        $this->errors = $errors;
    }

    public function get_errors()
    {
        return $this->errors;
    }
    
    public function to_array()
    {
        return array(
            "errors" => $this->errors->to_array(),
        );
    }
}
