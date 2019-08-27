<?php
namespace app\modules\todolist\models;

use Yoozoo\ProtoApi;

class AddResp implements ProtoApi\Message
{
    protected $count;

    public function init(array $response)
    {
        if (isset($response["count"])) {
            $this->set_count ( $response["count"] );
        }
    }

    public function validate()
    {
        if (!isset($this->count)) {
            throw new ProtoApi\GeneralException("'count' is not exist");
        }
    }
    
    public function set_count(int $count)
    {
        $this->count = $count;
    }

    public function get_count()
    {
        return $this->count;
    }
    
    public function to_array()
    {
        return array(
            "count" => $this->count,
        );
    }
}
