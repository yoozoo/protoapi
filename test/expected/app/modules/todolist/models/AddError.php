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
            $this->req->validate();
        }
        if (isset($response["error"])) {
            $this->error = $response["error"];
        }
    }

    public function validate()
    {
        if (!isset($this->req)) {
            throw new ProtoApi\GeneralException("'req' is not exist");
        }
        if (!isset($this->error)) {
            throw new ProtoApi\GeneralException("'error' is not exist");
        }
    }
    
    public function set_req(Req $req)
    {
        $this->req = $req;
    }

    public function get_req()
    {
        return $this->req;
    }
    
    public function set_error($error)
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
            "req" => $this->req->to_array(),
            "error" => $this->error,
        );
    }
}
