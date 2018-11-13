<?php
namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class FieldError implements ProtoApi\Message
{
    protected $fieldName;
    protected $errorType;

    public function init(array $response)
    {
        if (isset($response["fieldName"])) {
            $this->fieldName = $response["fieldName"];
        }
        if (isset($response["errorType"])) {
            $this->errorType = $response["errorType"];
        }
    }

    public function validate()
    {
        if (!isset($this->fieldName)) {
            throw new ProtoApi\GeneralException("'fieldName' is not exist");
        }
        if (!isset($this->errorType)) {
            throw new ProtoApi\GeneralException("'errorType' is not exist");
        }
    }
    
    public function set_fieldName($fieldName)
    {
        $this->fieldName = $fieldName;
    }

    public function get_fieldName()
    {
        return $this->fieldName;
    }
    
    public function set_errorType($errorType)
    {
        $this->errorType = $errorType;
    }

    public function get_errorType()
    {
        return $this->errorType;
    }
    
    public function to_array()
    {
        return array(
            "fieldName" => $this->fieldName,
            "errorType" => $this->errorType,
        );
    }
}
