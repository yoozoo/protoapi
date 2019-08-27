<?php

namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class ValidateError extends ProtoApi\BizErrorException implements ProtoApi\Message
{
    protected $errors;

    public function init(array $response)
    {
        if (isset($response["errors"])) {
            $val = $response["errors"];
            $this->set_errors( array_map( function($v) { $tmp = new FieldError(); $tmp->init($v); return $tmp; }, $val) );
        }
    }

    public function validate()
    {
        if (!isset($this->errors)) {
            throw new ProtoApi\GeneralException("'errors' is not exist");
        }
        array_filter($this->errors, function($v) { $v->validate(); return false; });
    }
    
    public function set_errors(array $errors)
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
            "errors" => array_map( function ($v) {  return $v->to_array(); }, $this->errors),
        );
    }
}
