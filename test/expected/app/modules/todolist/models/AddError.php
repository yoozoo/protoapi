<?php

namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class AddError extends ProtoApi\BizErrorException implements ProtoApi\Message
{
    protected $req;
    protected $error;

    public function init(array $response)
    {
        if (isset($response["req"])) {
            $this->req = new AddReq();
            $this->req->init($response["req"]);
        }
        if (isset($response["error"])) {
            $this->set_error ( $response["error"] );
        }
    }

    public function validate()
    {
        if (!isset($this->req)) {
            throw new ProtoApi\GeneralException("'req' is not exist");
        }
        $this->req->validate();
        if (!isset($this->error)) {
            throw new ProtoApi\GeneralException("'error' is not exist");
        }
    }
    
    public function set_req(AddReq $req)
    {
        $this->req = $req;
    }

    public function get_req()
    {
        return $this->req;
    }
    
    public function set_error(string $error)
    {
        $this->error = $error;
    }

    public function get_error()
    {
        return $this->error;
    }
    
    public function to_array()
    {
        return array(
            "req" =>  $this->req->to_array(),
            "error" => $this->error,
        );
    }
}
